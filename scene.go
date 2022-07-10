package rpg

import "github.com/95eh/eg"

type TileSceneConf struct {
	Width     int32 //地图宽
	Length    int32 //地图长
	TileSize  int32 //瓦片宽
	EventDis  int32 //事件最小距离
	VisionDis int32 //视野最小距离
}

type ITileSceneData interface {
	Conf() TileSceneConf
	GetVisionTileAroundTags(tx, ty int32) (aroundTags []string, centerTag string, err eg.IErr)
	GetEventTileAroundTags(tx, ty int32) (aroundTags []string, centerTag string, err eg.IErr)
	GetVisionTileChanged(otx, oty, itx, ity int32) (out, in []string)
	TestTilePosition(x, y float32) (tx, ty int32, nx, ny float32, err eg.IErr)
	GetTileTagByPosition(x, y float32) (string, eg.IErr)
}

type IMTileScene interface {
	eg.IModule
	AddSceneTplConf(sceneType eg.TScene, conf TileSceneConf) eg.IErr
	GetSceneTplData(sceneType eg.TScene) (sceneData ITileSceneData, err eg.IErr)
}

const (
	MTileScene = "MTileScene"
)

var (
	_MTileScene IMTileScene
)

func TileScene() IMTileScene {
	return _MTileScene
}

func init() {
	eg.RegisterModule(MTileScene, func(module eg.IModule) {
		_MTileScene = module.(IMTileScene)
	})
}
