package script

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pojol/gobot/driver/utils"
	lua "github.com/yuin/gopher-lua"
)

// from https://github.com/cjoudrey/gluahttp

type HttpModule struct {
	repolst []Report
	client  *http.Client
}

func NewHttpModule() *HttpModule {
	transport := &http.Transport{
		ForceAttemptHTTP2: false,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Second * 120,
	}

	return NewHttpModuleWithDo(client)
}

func NewHttpModuleWithDo(client *http.Client) *HttpModule {
	return &HttpModule{
		client: client,
	}
}

func (h *HttpModule) Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"get":     h.get,
		"post":    h.post,
		"put":     h.put,
		"request": h.request,
	})
	registerHttpResponseType(mod, L)
	L.Push(mod)
	return 1
}

func (h *HttpModule) get(L *lua.LState) int {
	return h.doRequestAndPush(L, "GET", L.ToString(1), L.ToTable(2))
}

func (h *HttpModule) post(L *lua.LState) int {
	return h.doRequestAndPush(L, "POST", L.ToString(1), L.ToTable(2))
}

func (h *HttpModule) put(L *lua.LState) int {
	return h.doRequestAndPush(L, "PUT", L.ToString(1), L.ToTable(2))
}

func (h *HttpModule) request(L *lua.LState) int {
	return h.doRequestAndPush(L, L.ToString(1), L.ToString(2), L.ToTable(3))
}

func (h *HttpModule) doRequest(L *lua.LState, method string, url string, options *lua.LTable) (*lua.LUserData, error) {

	req, err := http.NewRequest(method, url, nil)
	var reslen int
	if err != nil {
		fmt.Printf("new request %v err : %v\n", method, err.Error())
		return nil, err
	}

	if ctx := L.Context(); ctx != nil {
		req = req.WithContext(ctx)
	}

	if options != nil {
		if reqCookies, ok := options.RawGet(lua.LString("cookies")).(*lua.LTable); ok {
			reqCookies.ForEach(func(key lua.LValue, value lua.LValue) {
				req.AddCookie(&http.Cookie{Name: key.String(), Value: value.String()})
			})
		}

		switch reqQuery := options.RawGet(lua.LString("query")).(type) {
		case lua.LString:
			req.URL.RawQuery = reqQuery.String()
		}

		body := options.RawGet(lua.LString("body"))

		switch reqBody := body.(type) {
		case *lua.LTable:
			m, err := utils.Table2Map(reqBody)
			if err != nil {
				fmt.Println("table 2 map err", err.Error())
				return nil, err
			}
			byt, err := json.Marshal(m)
			if err != nil {
				fmt.Println("ltable marshal err", err.Error())
				return nil, err
			}
			req.Body = ioutil.NopCloser(bytes.NewReader(byt))
			req.Header.Set("Content-Type", "application/json")
		case lua.LString:
			req.Body = ioutil.NopCloser(bytes.NewBufferString(reqBody.String()))
			req.Header.Set("Content-Type", "application/json")
		}

		reqTimeout := options.RawGet(lua.LString("timeout"))
		if reqTimeout != lua.LNil {
			duration := time.Duration(0)
			switch reqTimeout := reqTimeout.(type) {
			case lua.LNumber:
				duration = time.Second * time.Duration(int(reqTimeout))
			case lua.LString:
				duration, err = time.ParseDuration(string(reqTimeout))
				if err != nil {
					fmt.Printf("parse timeout err %v\n", err.Error())
					return nil, err
				}
			}
			ctx, cancel := context.WithTimeout(req.Context(), duration)
			req = req.WithContext(ctx)
			defer cancel()
		}

		// Basic auth
		if reqAuth, ok := options.RawGet(lua.LString("auth")).(*lua.LTable); ok {
			user := reqAuth.RawGetString("user")
			pass := reqAuth.RawGetString("pass")
			if !lua.LVIsFalse(user) && !lua.LVIsFalse(pass) {
				req.SetBasicAuth(user.String(), pass.String())
			} else {
				return nil, fmt.Errorf("auth table must contain no nil user and pass fields")
			}
		}

		// Set these last. That way the code above doesn't overwrite them.
		if reqHeaders, ok := options.RawGet(lua.LString("headers")).(*lua.LTable); ok {
			reqHeaders.ForEach(func(key lua.LValue, value lua.LValue) {
				req.Header.Set(key.String(), value.String())
			})
		}
	}

	rep := Report{
		MsgID: url,
	}

	res, err := h.client.Do(req)
	if err != nil {
		err = fmt.Errorf("client do err : %v", err.Error())
		rep.Err = err.Error()
		h.repolst = append(h.repolst, rep)
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("read body err : %v", err.Error())
		rep.Err = err.Error()
		h.repolst = append(h.repolst, rep)
		return nil, err
	}

	h.repolst = append(h.repolst, rep)
	return newHttpResponse(res, &body, reslen, L), nil
}

func (h *HttpModule) doRequestAndPush(L *lua.LState, method string, url string, options *lua.LTable) int {
	response, err := h.doRequest(L, method, url, options)

	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(fmt.Sprintf("%s", err)))
		return 2
	}

	L.Push(response)
	return 1
}

func (h *HttpModule) GetReport() []Report {

	rep := []Report{}
	rep = append(rep, h.repolst...)

	h.repolst = h.repolst[:0]

	return rep
}
