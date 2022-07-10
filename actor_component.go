package rpg

import "github.com/95eh/eg"

const (
	Ac_Transform eg.TActorComponent = iota
	Ac_Life
	Ac_Equip
	Ac_Rand_Move
)

type IActorComTransform interface {
	eg.IActorComponent
	Position() eg.Vec3
	SetPosition(v eg.Vec3)
	Forward() eg.Vec3
	MoveSpeed() float32
	SetMoveSpeed(moveSpeed float32)
	Tile() (int32, int32)
	SetTile(tx, ty int32)
	GetTileTag() string
	GetVisionTileTags() ([]string, string)
	GetEventTileTags() ([]string, string)
	Moving() bool
	StartMove(forward eg.Vec3) (tags []string)
	StopMove() (tags []string)
	Move(offset eg.Vec3)
}

type IActorComLife interface {
	eg.IActorComponent
	Name() string
	SetName(v string)
}
