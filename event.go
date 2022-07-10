package rpg

import "github.com/95eh/eg"

const (
	Evt_Visible eg.TActorEvent = iota
	Evt_Invisible
	Evt_MoveStart
	Evt_MoveStop
	Evt_ChangePos
	Evt_ForwardChange
	Evt_Max
)
