package weapons

import "github.com/fedom-git/gs-cal/core"

type Rust struct {
	weaponType string          `default:"bow" yaml:"weaponType"`
	Base       core.FixedPanel `yaml:"base"`
	GiveBuffs  core.BuffMap    `yaml:"giveBuffs"`
}

func (r *Rust) GenerateBuff() *core.BuffMap {
	return &r.GiveBuffs
}
