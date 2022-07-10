package rpg

import (
	"github.com/95eh/eg"
	"strconv"
)

func GetTileAroundSize(tw, tl, x, y, size int32) (minX, maxX, minY, maxY int32) {
	minX, _ = eg.FloorInt32(0, x-size)
	maxX, _ = eg.CeilInt32(tw, x+size+1)
	minY, _ = eg.FloorInt32(0, y-size)
	maxY, _ = eg.CeilInt32(tl, y+size+1)
	return
}

func MergeSceneTileKey(x, y int32) string {
	return strconv.FormatInt(int64(x), 10) + "_" + strconv.FormatInt(int64(y), 10)
}

func GetTileAroundTiles(minX, maxX, minY, maxY int32) []string {
	tiles := make([]string, 0, (maxX-minX)*(maxY-minY))
	for i := minX; i < maxX; i++ {
		for j := minY; j < maxY; j++ {
			tiles = append(tiles, MergeSceneTileKey(i, j))
		}
	}
	return tiles
}
