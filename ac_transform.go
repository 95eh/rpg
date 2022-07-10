package rpg

import (
	"github.com/95eh/eg"
	"github.com/95eh/eg/scene"
	"time"
)

const (
	TnfMoveTickDur = int64(100)
	TnfMoveTicker  = "tnf_move"
	TnfPos         = "tnf_pos"
	TnfFor         = "tnf_for"
	TnfMoveSpeed   = "tnf_move_speed"
)

var (
	TnfMoveTickSkipDur = TnfMoveTickDur >> 1
)

func NewAcTransform(o *eg.Object) eg.IActorComponent {
	return &AcTransform{
		position:  o.MustVec3(TnfPos, eg.NewVec3(10, 0, 10)),
		forward:   o.MustVec3(TnfFor, eg.NewVec3(0, 0, 0)),
		moveSpeed: o.MustFloat32(TnfMoveSpeed, 4),
	}
}

type AcTransform struct {
	scene.ActorComponent
	tileSceneData      ITileSceneData
	position           eg.Vec3
	forward            eg.Vec3
	moveSpeed          float32
	tileX              int32
	tileY              int32
	moving             bool
	moveTickerId       int64
	moveLastTime       int64
	moveTargetPosition eg.Vec3
}

func (t *AcTransform) Position() eg.Vec3 {
	return t.position
}

func (t *AcTransform) SetPosition(v eg.Vec3) {
	t.position = v
}

func (t *AcTransform) Forward() eg.Vec3 {
	return t.forward
}

func (t *AcTransform) MoveSpeed() float32 {
	return t.moveSpeed
}

func (t *AcTransform) SetMoveSpeed(moveSpeed float32) {
	t.moveSpeed = moveSpeed
}

func (t *AcTransform) Tile() (int32, int32) {
	return t.tileX, t.tileY
}

func (t *AcTransform) SetTile(tx, ty int32) {
	t.tileX, t.tileY = tx, ty
}

func (t *AcTransform) GetTileTag() string {
	return MergeSceneTileKey(t.tileX, t.tileY)
}

func (t *AcTransform) GetVisionTileTags() ([]string, string) {
	tags, tag, _ := t.tileSceneData.GetVisionTileAroundTags(t.tileX, t.tileY)
	return tags, tag
}

func (t *AcTransform) GetEventTileTags() ([]string, string) {
	tags, tag, _ := t.tileSceneData.GetEventTileAroundTags(t.tileX, t.tileY)
	return tags, tag
}

func (t *AcTransform) Moving() bool {
	return t.moving
}

func (t *AcTransform) StartMove(forward eg.Vec3) (tags []string) {
	if t.moving {
		t.tickMove()
	} else {
		t.moving = true
		t.moveLastTime = time.Now().UnixMilli()
		t.moveTickerId = eg.SId().GetRegionId()
		eg.Timer().AddQuickTicker(TnfMoveTicker, t.moveTickerId, t.tickMove)
	}
	t.forward = forward
	tx, ty := t.tileX, t.tileY
	tags, _, _ = t.tileSceneData.GetEventTileAroundTags(tx, ty)
	return
}

func (t *AcTransform) StopMove() (tags []string) {
	tx, ty := t.tileX, t.tileY
	tags, _, _ = t.tileSceneData.GetEventTileAroundTags(tx, ty)
	if !t.moving {
		return
	}
	t.tickMove()
	t.moving = false
	eg.Timer().DelQuickTicker(TnfMoveTicker, t.moveTickerId)
	t.moveTickerId = 0
	return
}

func (t *AcTransform) Type() eg.TActorComponent {
	return Ac_Transform
}

func (t *AcTransform) Move(offset eg.Vec3) {
	actor := t.Actor()
	actorId := actor.Id()
	x, y, z := t.position.X+offset.X, t.position.Y+offset.Y, t.position.Z+offset.Z
	tx, ty, x, z, _ := t.tileSceneData.TestTilePosition(x, z)
	t.position.X, t.position.Y, t.position.Z = x, y, z

	//所在格子没变化
	if tx == t.tileX && ty == t.tileY {
		return
	}
	//eg.Log().Debug("tile change", eg.M{
	//	"tx":  tx,
	//	"ty":  ty,
	//	"pos": t.position,
	//})
	currTag := t.GetTileTag()
	out, in := t.tileSceneData.GetVisionTileChanged(t.tileX, t.tileY, tx, ty)
	t.tileX, t.tileY = tx, ty
	tag := t.GetTileTag()

	evtInvisible := NewEvtInvisible(actor)
	evtVisible := NewEvtVisible(actor)
	var evtMoveStart eg.IActorEvent
	s := eg.Scene().SpawnScheduler(0, nil).
		GetActorAndScene(actorId).
		FnScene(func(scene eg.IScene, s eg.ISceneWorkerScheduler) eg.IErr {
			scene.RemoveActorTags(actorId, currTag)
			scene.AddActorTags(actorId, tag)
			return nil
		}).
		SetTags(out).
		GetTagsActors().
		ResetActorEvents(1).
		FnActors(func(actor eg.IActor, s eg.ISceneWorkerScheduler) eg.IErr {
			actor.ProcessEvent(evtInvisible)
			return s.PushEvent(NewEvtInvisible(actor))
		}).
		ResetActors().
		SetTags(in).
		GetTagsActors().
		ExpansionActorEvents(2)
	if t.moving {
		evtMoveStart = NewEvtMoveStart(actor)
		s.FnActors(func(actor eg.IActor, s eg.ISceneWorkerScheduler) eg.IErr {
			if actor.Id() == actorId {
				return nil
			}
			actor.ProcessEvent(evtVisible)
			actor.ProcessEvent(evtMoveStart)
			err := s.PushEvent(NewEvtVisible(actor))
			if err != nil {
				return err
			}
			com, _ := actor.GetComponent(Ac_Transform)
			tnf := com.(*AcTransform)
			if tnf.moving {
				err := s.PushEvent(NewEvtMoveStart(actor))
				if err != nil {
					return err
				}
			}
			return nil
		})
	} else {
		s.FnActors(func(actor eg.IActor, s eg.ISceneWorkerScheduler) eg.IErr {
			if actor.Id() == actorId {
				return nil
			}
			actor.ProcessEvent(evtVisible)
			s.PushEvent(NewEvtVisible(actor))
			com, _ := actor.GetComponent(Ac_Transform)
			tnf := com.(*AcTransform)
			if tnf.moving {
				s.PushEvent(NewEvtMoveStart(actor))
			}
			return nil
		})
	}
	s.ActorProcessEvents().
		Do(nil, nil)
}

func (t *AcTransform) Start() (err eg.IErr) {
	t.tileSceneData, err = TileScene().GetSceneTplData(t.Actor().Scene().Type())
	if err != nil {
		return
	}
	t.tileX, t.tileY, t.position.X, t.position.Z, err =
		t.tileSceneData.TestTilePosition(t.position.X, t.position.Z)
	if err != nil {
		return
	}
	t.Actor().Scene().AddActorTags(t.Actor().Id(), MergeSceneTileKey(t.tileX, t.tileY))
	return
}

func (t *AcTransform) Dispose() eg.IErr {
	t.Actor().Scene().RemoveActorTags(t.Actor().Id(), MergeSceneTileKey(t.tileX, t.tileY))
	t.tileSceneData = nil
	return nil
}

func (t *AcTransform) tickMove() {
	eg.Share().Push(eg.Int64ToStr(t.Actor().Id()), func() {
		now := time.Now().UnixMilli()
		durMs := now - t.moveLastTime
		if durMs < TnfMoveTickSkipDur {
			return
		}
		t.moveLastTime = now
		t.Move(eg.Vec3Mulf(t.forward, float32(durMs)/1000*t.moveSpeed))
	})
}
