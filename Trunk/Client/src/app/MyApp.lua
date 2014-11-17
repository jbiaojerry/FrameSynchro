
require("config")
require("framework.init")
NetworkHandler = require("app.Logic.NetworkHandler")
local MyApp = class("MyApp", cc.mvc.AppBase)

function MyApp:ctor()
    MyApp.super.ctor(self)
end

function MyApp:run()
    cc.FileUtils:getInstance():addSearchPath("res/")
    self:enterScene("MainScene")
    NetworkHandler:GetIntance():Init()
end

return MyApp
