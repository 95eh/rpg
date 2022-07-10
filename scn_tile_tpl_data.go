package rpg

import (
	"github.com/95eh/eg"
	"math"
)

func NewTileSceneTplData(conf TileSceneConf) *tileSceneTplData {
	return &tileSceneTplData{
		conf:             conf,
		visionAroundTags: make(map[string][]string),
		eventAroundTags:  make(map[string][]string),
	}
}

type tileSceneTplData struct {
	conf             TileSceneConf
	visionAroundTags map[string][]string
	eventAroundTags  map[string][]string
}

func (t *tileSceneTplData) Conf() TileSceneConf {
	return t.conf
}

func (t *tileSceneTplData) TestTilePosition(x, y float32) (tx, ty int32, nx, ny float32, err eg.IErr) {
	nx = eg.ClampFloat(1, float32(t.conf.Width)-1, x)
	ny = eg.ClampFloat(1, float32(t.conf.Length)-1, y)
	tx = int32(math.Floor(float64(nx) / float64(t.conf.TileSize)))
	ty = int32(math.Floor(float64(ny) / float64(t.conf.TileSize)))
	return
}

func (t *tileSceneTplData) GetTileTagByPosition(x, y float32) (string, eg.IErr) {
	tx, ty, _, _, err := t.TestTilePosition(x, y)
	if err != nil {
		return "", err
	}
	return MergeSceneTileKey(tx, ty), nil
}

func (t *tileSceneTplData) GetVisionTileAroundTags(tx, ty int32) (aroundTags []string, centerTag string, err eg.IErr) {
	return getTileAroundTags(t.visionAroundTags, tx, ty)
}

func (t *tileSceneTplData) GetEventTileAroundTags(tx, ty int32) (aroundTags []string, centerTag string, err eg.IErr) {
	return getTileAroundTags(t.eventAroundTags, tx, ty)
}

func (t *tileSceneTplData) GetVisionTileChanged(otx, oty, itx, ity int32) (out, in []string) {
	o, _, _ := t.GetVisionTileAroundTags(otx, oty)
	i, _, _ := t.GetVisionTileAroundTags(itx, ity)
	om := make(map[string]struct{}, len(out))
	for _, tag := range o {
		om[tag] = struct{}{}
	}
	in = make([]string, 0, len(i))
	for _, tag := range i {
		if _, ok := om[tag]; ok {
			delete(om, tag)
			continue
		}
		in = append(in, tag)
	}
	out = make([]string, 0, len(om))
	for tag := range om {
		out = append(out, tag)
	}
	return
}

func getTileAroundTags(m map[string][]string, tx, ty int32) (aroundTags []string, centerTag string, err eg.IErr) {
	var ok bool
	aroundTags, ok = m[MergeSceneTileKey(tx, ty)]
	if !ok {
		err = eg.NewErr(eg.EcNotExistKey, eg.M{
			"tile x": tx,
			"tile y": ty,
		})
		return
	}
	centerTag = MergeSceneTileKey(tx, ty)
	return
}
