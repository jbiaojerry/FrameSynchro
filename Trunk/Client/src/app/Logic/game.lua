
require("config")
require("framework.init")
local NetworkHandler = require("app.Logic.NetworkHandler")

-- define global module
game = {}

function game.startup()
    cc.FileUtils:getInstance():addSearchPath("res/")
    display.addSpriteFrames(GAME_TEXTURE_DATA_FILENAME, GAME_TEXTURE_IMAGE_FILENAME)
    --cc.FrameMgr:getInstance():getFrameCnt();
    NetworkHandler:Init()
    game.enterMainScene()
end

function game.exit()
    os.exit()
end

function game.enterMainScene()
    display.replaceScene(require("app.scenes.MainScene").new(), "fade", 0.6, display.COLOR_WHITE)
end
