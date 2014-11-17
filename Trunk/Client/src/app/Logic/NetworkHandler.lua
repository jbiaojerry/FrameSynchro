local NetworkHandler = class("NetworkHandler")
local SocketTCP = require("framework.cc.net.SocketTCP")
data = require("app.example.test.lua")

NetworkHandler.EVENT_SERVER_SAY_HELLO = "EVENT_SERVER_SAY_HELLO"

function NetworkHandler:ctor()
    print ("NetworkHandler:ctor")
    cc(self):addComponent("components.behavior.EventProtocol"):exportMethods()
end

function NetworkHandler:GetIntance()
    print ("NetworkHandler:GetIntance")
    if self.instance then
        return self.instance
    end
    print ("NetworkHandler:GetIntance1")
    self.instance = NetworkHandler.new()
    return self.instance
end

function NetworkHandler:Release()
    self.instance = nil
end

function NetworkHandler:Init()
    self.SocketTCP = SocketTCP.new();
    self.SocketTCP:addEventListener(self.SocketTCP.EVENT_DATA, handler(self, self.OnData))
    self.SocketTCP:connect("127.0.0.1", 8080, false)
    print (self.SocketTCP.isConnected)
    --self.SocketTCP:send("hello server.")
    self.SocketTCP:send(data)
end

function NetworkHandler:Shutdown()
end

function NetworkHandler:OnData(__event)
    self:dispatchEvent({name=NetworkHandler.EVENT_SERVER_SAY_HELLO, data = __event.data})
end

return NetworkHandler
