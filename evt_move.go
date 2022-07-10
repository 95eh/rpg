package rpg

import (
	"github.com/95eh/eg"
)

func NewEvtMoveStart(actor eg.IActor) eg.IActorEvent {
	com, _ := actor.GetComponent(Ac_Transform)
	tnf := com.(IActorComTransform)
	pos := tnf.Position()
	forward := tnf.Forward()
	return &EvtMoveStart{
		ActorId:   actor.Id(),
		ActorType: actor.Type(),
		PositionX: pos.X,
		PositionY: pos.Y,
		PositionZ: pos.Z,
		ForwardX:  forward.X,
		ForwardY:  forward.Y,
		ForwardZ:  forward.Z,
	}
}

type EvtMoveStart struct {
	ActorId   int64
	ActorType eg.TActor
	PositionX float32
	PositionY float32
	PositionZ float32
	ForwardX  float32
	ForwardY  float32
	ForwardZ  float32
}

func (e *EvtMoveStart) Type() eg.TActorEvent {
	return Evt_MoveStart
}

func NewEvtMoveStop(actor eg.IActor) eg.IActorEvent {
	com, _ := actor.GetComponent(Ac_Transform)
	tnf := com.(IActorComTransform)
	pos := tnf.Position()
	return &EvtMoveStop{
		ActorId:   actor.Id(),
		ActorType: actor.Type(),
		PositionX: pos.X,
		PositionY: pos.Y,
		PositionZ: pos.Z,
	}
}

type EvtMoveStop struct {
	ActorId   int64
	ActorType eg.TActor
	PositionX float32
	PositionY float32
	PositionZ float32
}

func (e *EvtMoveStop) Type() eg.TActorEvent {
	return Evt_MoveStop
}

func NewEvtPosChange(actor eg.IActor) eg.IActorEvent {
	com, _ := actor.GetComponent(Ac_Transform)
	tnf := com.(IActorComTransform)
	pos := tnf.Position()
	forward := tnf.Forward()
	return &EvtPosChange{
		ActorId:   actor.Id(),
		ActorType: actor.Type(),
		PositionX: pos.X,
		PositionY: pos.Y,
		PositionZ: pos.Z,
		ForwardX:  forward.X,
		ForwardY:  forward.Y,
		ForwardZ:  forward.Z,
	}
}

type EvtPosChange struct {
	ActorId   int64
	ActorType eg.TActor
	PositionX float32
	PositionY float32
	PositionZ float32
	ForwardX  float32
	ForwardY  float32
	ForwardZ  float32
}

func (e *EvtPosChange) Type() eg.TActorEvent {
	return Evt_ChangePos
}
