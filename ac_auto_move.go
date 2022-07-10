package rpg

import (
	"github.com/95eh/eg"
	"github.com/95eh/eg/scene"
	"math/rand"
	"time"
)

func NewAcAutoMove(o *eg.Object) eg.IActorComponent {
	return &AcAutoMove{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

type AcAutoMove struct {
	scene.ActorComponent
	rand *rand.Rand
}

func (t *AcAutoMove) Type() eg.TActorComponent {
	return Ac_Rand_Move
}

func (t *AcAutoMove) Start() eg.IErr {
	t.Move()
	return nil
}

func (t *AcAutoMove) Move() {
	delay := 8000 + t.rand.Int63n(2000)
	actorId := t.Actor().Id()
	eg.Timer().After(delay, func() {
		eg.Scene().SpawnScheduler(0, nil).
			GetActorAndScene(actorId).
			FnActor(func(actor eg.IActor, s eg.ISceneWorkerScheduler) eg.IErr {
				c, _ := t.Actor().GetComponent(Ac_Transform)
				tnf := c.(*AcTransform)
				v := eg.Vec2Normalize(eg.Vec2{
					X: float32(t.rand.Int31n(10)) - 4.5,
					Y: float32(t.rand.Int31n(10)) - 4.5,
				})
				//eg.Log().Debug("move forward", misc.M{
				//	"x": v.X,
				//	"y": v.Y,
				//})
				tnf.StartMove(eg.NewVec3(v.X, 0, v.Y))
				t.Move()
				return nil
			}).Do(nil, nil)
	})
}
