package core

type Weapon struct {
	Panel      `yaml:"panel"`
	weaponType string  `default:"bow" yaml:"weaponType"`
	GiveBuffs  BuffMap `yaml:"giveBuffs"`
}

//func (w *Weapon) GetEntity() *FixedPanel {
//	return &w.Base
//}

func (w *Weapon) GenerateBuffs() *BuffMap {
	return &w.GiveBuffs
}
