package core

import "github.com/fedom-git/gs-cal/util/log"

type Panel struct {
	Name string             `yaml:"name"`
	Base map[string]float64 `yaml:"base"`
}

func (p *Panel) GetBase() *map[string]float64 {
	return &p.Base
}

type FixedPanel struct {
	Name    string  `yaml:"name"`
	EType   string  `yaml:"eType"`
	H       float64 `yaml:"h"`
	Hr      float64 `yaml:"hr"`
	A       float64 `yaml:"a"`
	Ar      float64 `yaml:"at"`
	D       float64 `yaml:"df"`
	Dr      float64 `yaml:"dr"`
	Cr      float64 `yaml:"cr"`
	Cd      float64 `yaml:"cd"`
	Em      float64 `yaml:"em"`
	Ec      float64 `yaml:"ec"`
	Pyro    float64 `yaml:"pyro"`
	Hydro   float64 `yaml:"hydro"`
	Anemo   float64 `yaml:"anemo"`
	Electro float64 `yaml:"electro"`
	Dendro  float64 `yaml:"dendro"`
	Cryo    float64 `yaml:"cryo"`
	Geo     float64 `yaml:"geo"`
	Allo    float64 `yaml:"allo"`
}

type Base struct {
	Name  string             `yaml:"name"`
	Panel map[string]float64 `yaml:"panel"`
}

func (b *Base) GetBase() *Base {
	return b
}

type Buff struct {
	Type string  `yaml:"type"`
	V    float64 `yaml:"v"`
}

type BuffMap map[string][]Buff

func (bm *BuffMap) GetBuffVecByName(name string) *[]Buff {
	vec, ok := (*bm)[name]
	if !ok {
		log.Log.Printf("%s not exist in current buffMap, %+v", name, bm)
		return nil
	}
	return &vec
}

type GsEntity interface {
	//GetEntity() *FixedPanel
	GetBase() *Base
	GetPanel() *map[string]float64
	EmitBuffs() *BuffMap
	TakeBuffs(buffs *BuffMap)
	PostMarshal()
}

const (
	CharacterPath = "./DB/characters/"
	WeaponPath    = "./DB/weapons/"
	ArtifactPath  = "./DB/artifacts/"
)

func GetCharacterPath(name string) string {
	return CharacterPath + name + ".yml"
}

func GetWeaponPath(name string) string {
	return WeaponPath + name + ".yml"
}

func GetArtifactPath(name string) string {
	return ArtifactPath + name + ".yml"
}
