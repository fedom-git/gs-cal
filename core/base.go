package core

type Base struct {
	name    string  `yaml:"name"`
	eType   string  `yaml:"EType"`
	h       int     `yaml:"h"`
	hr      float32 `yaml:"hr"`
	a       int     `yaml:"a"`
	ar      float32 `yaml:"at"`
	d       int     `yaml:"df"`
	dr      float32 `yaml:"dr"`
	cr      float32 `yaml:"cr"`
	cd      float32 `yaml:"cd"`
	em      int     `yaml:"em"`
	ec      int     `yaml:"ec"`
	pyro    float32 `yaml:"pyro"`
	hydro   float32 `yaml:"hydro"`
	anemo   float32 `yaml:"anemo"`
	electro float32 `yaml:"electro"`
	dendro  float32 `yaml:"dendro"`
	cryo    float32 `yaml:"cryo"`
	geo     float32 `yaml:"geo"`
	allo    float32 `yaml:"allo"`
}
