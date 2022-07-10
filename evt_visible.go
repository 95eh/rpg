package rpg

import (
	"github.com/95eh/eg"
)

func NewEvtVisible(actor eg.IActor) eg.IActorEvent {
	c, _ := actor.GetComponent(Ac_Transform)
	tnf := c.(*AcTransform)
	tx, ty := tnf.Tile()
	return &EvtVisible{
		ActorId:   actor.Id(),
		ActorType: actor.Type(),
		TileX:     tx,
		TileY:     ty,
		Pos:       tnf.Position(),
	}
}

type EvtVisible struct {
	ActorId   int64
	ActorType eg.TActor
	TileX     int32
	TileY     int32
	Pos       eg.Vec3
}

func (e *EvtVisible) Type() eg.TActorEvent {
	return Evt_Visible
}

type EvtInvisible struct {
	ActorId   int64
	ActorType eg.TActor
	TileX     int32
	TileY     int32
	Pos       eg.Vec3
}

func (e *EvtInvisible) Type() eg.TActorEvent {
	return Evt_Invisible
}

func NewEvtInvisible(actor eg.IActor) eg.IActorEvent {
	c, _ := actor.GetComponent(Ac_Transform)
	tnf := c.(*AcTransform)
	tx, ty := tnf.Tile()
	return &EvtInvisible{
		ActorId:   actor.Id(),
		ActorType: actor.Type(),
		TileX:     tx,
		TileY:     ty,
		Pos:       eg.Vec3{},
	}
}
