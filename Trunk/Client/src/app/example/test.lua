package.path = package.path .. ';app/protobuf/?.lua;app/Protocol/?.lua'
package.cpath = package.cpath .. ';app/protobuf/?.so'

require 'person_pb'

local person= person_pb.Person()
person.id = 1001
person.name = "今心我在"
person.email = "lyn0328@qq.com"

local home = person.Extensions[person_pb.Phone.phones]:add()
home.num = "18028750328"
home.type = person_pb.Phone.HOME

local data = person:SerializeToString()
return data
--[[
local msg = person_pb.Person()
msg:ParseFromString(data)
print(msg.id)
print("test.lua")--]]
