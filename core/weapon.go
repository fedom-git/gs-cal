package core

type Weapon struct {
	at int     `yaml:"at"`
	ar float32 `yaml:"ar"` //unit 1%
	hr float32 `yaml:"hr"`
	cr float32 `yaml:"cr"`
	cd float32 `yaml:"cd"`
	em int     `yaml:"em"`
	ec int     `yaml:"ec"`
}

func (w *Weapon) Activate() {
	w.cr = 20 //sky bow
}
