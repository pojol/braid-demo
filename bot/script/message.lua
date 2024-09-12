
-- 字节序
-- 大端：BigEndian
-- 小端：LittleEndian
ByteOrder = "LittleEndian"


-- 从报文中解析出消息ID和消息体
-- 示例，用户可以参照实际报文格式进行解析
function WSUnpackMsg(buf, errmsg)

    if errmsg ~= "nil" then
        return 0, ""
    end

    local msg = message.new(buf, ByteOrder, 0)

    local msgId = msg:readi2()
    local msgbody = msg:readBytes(2, -1)

    return msgId, msgbody
    
end

function WSPackMsg(msgid, msgbody)

    local msg = message.new("", ByteOrder, 6+#msgbody)
    msg:writei2(msgid)
    msg:writeBytes(msgbody)

    return msg:pack()

end

------------------------------------------------------------------------
-- msglen : conn will first read the predefined message length field 
function TCPUnpackMsg(msglen, buf, errmsg)
    if errmsg ~= "nil" then
        return 0, ""
    end

    local msg = message.new(buf, ByteOrder, 0)
    local msgId = msg:readi2()
    local msgbody = msg:readBytes(2, -1)

    return msgId, msgbody

end

function TCPPackMsg(msgid, msgbody)
    local msglen = #msgbody+2

    local msg = message.new("", ByteOrder, msglen)
    msg:writei2(msgid)
    msg:writeBytes(msgbody)

    return msg:pack()

end