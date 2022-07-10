package rpg

import (
	"github.com/95eh/eg"
	"math"
)

func NewMTileScene() IMTileScene {
	return &mTileScene{
		IModule:       eg.NewModule(),
		typeToTplData: make(map[eg.TScene]*tileSceneTplData),
	}
}

type mTileScene struct {
	eg.IModule
	typeToTplData map[eg.TScene]*tileSceneTplData
}

func (M *mTileScene) Type() string {
	return MTileScene
}

func (M *mTileScene) AddSceneTplConf(sceneType eg.TScene, conf TileSceneConf) eg.IErr {
	_, ok := M.typeToTplData[sceneType]
	if ok {
		return eg.NewErr(eg.EcNotExistKey, eg.M{
			"scene type": sceneType,
		})
	}
	var tplData = NewTileSceneTplData(conf)
	tw := int32(math.Ceil(float64(conf.Width) / float64(conf.TileSize)))
	tl := int32(math.Ceil(float64(conf.Length) / float64(conf.TileSize)))
	visionSize := int32(math.Ceil(float64(conf.VisionDis) / float64(conf.TileSize)))
	eventSize := int32(math.Ceil(float64(conf.EventDis) / float64(conf.TileSize)))
	for x := int32(0); x < tw; x++ {
		for y := int32(0); y < tl; y++ {
			tileTag := MergeSceneTileKey(x, y)
			minX, maxX, minY, maxY := GetTileAroundSize(tw, tl, x, y, visionSize)
			tplData.visionAroundTags[tileTag] = GetTileAroundTiles(minX, maxX, minY, maxY)
			minX, maxX, minY, maxY = GetTileAroundSize(tw, tl, x, y, eventSize)
			tplData.eventAroundTags[tileTag] = GetTileAroundTiles(minX, maxX, minY, maxY)
		}
	}
	M.typeToTplData[sceneType] = tplData
	return nil
}

func (M *mTileScene) GetSceneTplData(sceneType eg.TScene) (ITileSceneData, eg.IErr) {
	sceneData, ok := M.typeToTplData[sceneType]
	if !ok {
		return nil, eg.NewErr(eg.EcNotExistKey, eg.M{
			"scene type": sceneType,
		})
	}
	return sceneData, nil
}
