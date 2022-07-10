package rpg

import (
	"github.com/95eh/eg"
	"github.com/95eh/eg/scene"
)

const (
	LfName = "lf_name"
)

func NewAcLife(o *eg.Object) eg.IActorComponent {
	return &AcLife{
		name: o.MustString(LfName, "null"),
	}
}

type AcLife struct {
	scene.ActorComponent
	name string
}

func (t *AcLife) Type() eg.TActorComponent {
	return Ac_Life
}

func (t *AcLife) Name() string {
	return t.name
}

func (t *AcLife) SetName(v string) {
	t.name = v
}
