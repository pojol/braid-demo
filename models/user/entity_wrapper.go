package user

import (
	"braid-demo/constant"
	"context"
	"fmt"
	"reflect"

	"github.com/pojol/braid/core"
	"github.com/pojol/braid/core/actor"
)

type EntityWrapper struct {
	ID       string                `bson:"_id"`
	cs       core.ICacheStrategy   `bson:"-"`
	Bag      *EntityBagModule      `bson:"bag"`
	Airship  *EntityAirshipModule  `bson:"airship"`
	User     *EntityUserModule     `bson:"user"`
	TimeInfo *EntityTimeInfoModule `bson:"time_info"`

	// Used to determine if it was read from cache
	isCache bool `bson:"-"`
}

func (e *EntityWrapper) GetID() string {
	return e.ID
}

func (e *EntityWrapper) SetModule(moduleType reflect.Type, module interface{}) {
	switch moduleType {
	case reflect.TypeOf(&EntityBagModule{}):
		e.Bag = module.(*EntityBagModule)
	case reflect.TypeOf(&EntityAirshipModule{}):
		e.Airship = module.(*EntityAirshipModule)
	case reflect.TypeOf(&EntityUserModule{}):
		e.User = module.(*EntityUserModule)
	case reflect.TypeOf(&EntityTimeInfoModule{}):
		e.TimeInfo = module.(*EntityTimeInfoModule)
	}
}

func (e *EntityWrapper) GetModule(moduleType reflect.Type) interface{} {
	switch moduleType {
	case reflect.TypeOf(&EntityBagModule{}):
		return e.Bag
	case reflect.TypeOf(&EntityAirshipModule{}):
		return e.Airship
	case reflect.TypeOf(&EntityUserModule{}):
		return e.User
	case reflect.TypeOf(&EntityTimeInfoModule{}):
		return e.TimeInfo
	}
	return nil
}

func NewEntityWapper(id string) *EntityWrapper {
	e := &EntityWrapper{
		ID: id,
	}
	e.cs = actor.BuildEntityLoader(constant.MongoDatabase, constant.MongoCollection, e)
	return e
}

func (e *EntityWrapper) Load(ctx context.Context) error {
	err := e.cs.Load(ctx)
	if err != nil {
		return fmt.Errorf("load entity %v err %v", e.ID, err.Error())
	}

	e.setModulesAndIDs()
	e.isCache = true

	return nil
}

func (e *EntityWrapper) setModulesAndIDs() {
	v := reflect.ValueOf(e).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.Ptr && field.Type().Elem().Kind() == reflect.Struct {
			moduleType := field.Type()
			if module := e.cs.GetModule(moduleType); module != nil {
				field.Set(reflect.ValueOf(module))
				if idField := reflect.ValueOf(module).Elem().FieldByName("ID"); idField.IsValid() && idField.CanSet() {
					idField.SetString(e.ID)
				}
			}
		}
	}
}

func (e *EntityWrapper) Sync(ctx context.Context) error {
	return e.cs.Sync(ctx)
}

func (e *EntityWrapper) Store(ctx context.Context) error {
	return e.cs.Store(ctx)
}

func (e *EntityWrapper) IsDirty() bool {
	return e.cs.IsDirty()
}
