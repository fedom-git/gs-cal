package core

import (
	"github.com/fedom-git/gs-cal/util/log"
	"gopkg.in/yaml.v3"
	"os"
)

type Character struct {
	Panel         `yaml:"panel"`
	Skills        map[string]Skill  `yaml:"skills"`
	Teammates     []string          `yaml:"teammates"`
	GiveBuffs     map[string][]Buff `yaml:"giveBuffs"`
	RecievedBuffs map[string][]Buff `yaml:"takeBuffs"`

	//set after calc

	Final      map[string]float64
	Augment    map[string]float64 `yaml:"augment"`
	Resistance map[string]float64 `yaml:"resistance"`
	Defence    map[string]float64 `yaml:"defence"`
}

const (
	Physical      = 1
	Pyro          = 1 << 1
	Hydro         = 1 << 2
	Anemo         = 1 << 3
	Electro       = 1 << 4
	Dendro        = 1 << 5
	Cryo          = 1 << 6
	Geo           = 1 << 7
	Allo          = 1 << 8
	Amplification = 1 << 9
	Upheaval      = 1 << 10
	NormalAttack  = 1 << 11
	ChargedAttack = 1 << 12
	E             = 1 << 13
	Q             = 1 << 14

	UpheavalBaseDamage = 723.0
)

var (
	DamageTypeMap = map[string]int{
		"physical":      Physical,
		"pyro":          Pyro,
		"hydro":         Hydro,
		"anemo":         Anemo,
		"electro":       Electro,
		"dendro":        Dendro,
		"Cryo":          Cryo,
		"geo":           Geo,
		"allo":          Allo,
		"amplification": Amplification,
		"upheaval":      Upheaval,
		"normalAttack":  NormalAttack,
		"chargedAttack": ChargedAttack,
		"e":             E,
		"q":             Q,
	}

	UpheavalDamageMap = map[string]float64{
		"superConduction": UpheavalBaseDamage * 1,
		"diffusion":       UpheavalBaseDamage * 1.2,
		"electricShock":   UpheavalBaseDamage * 2.4,
		"iceBreak":        UpheavalBaseDamage * 3,
		"overload":        UpheavalBaseDamage * 4,
	}
)

type Skill struct {
	SkillType        string    `yaml:"skillType"` //skill type: damage, buff, heal,
	ElementType      string    `yaml:"elementType"`
	DamageSrc        string    `yaml:"damageSrc"` // a,e,q, diffusion
	BuffType         string    `yaml:"buffType"`
	RawRate          []float64 `yaml:"rawRate"`
	RawRateSum       float64
	Reactions        map[string][]float64 `yaml:"reactions"`
	AmplificationSum map[string]float64
	ReactionFator    map[string]float64
	//ReactionType    string  `yaml:"reactionType"`
	//ReactionFactor  float64 `yaml:"reactionFactor"` //1.5 or 2.0
	//ReactionCount   float64 `yaml:"reactionCount"`
}

func (c *Character) postMarshal() {
	ParseSkill(&c.Skills)
}

func ParseCharacter(path string) *Character {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Log.Println("parse error: ", err.Error())
		return nil
	}
	character := &Character{}
	err = yaml.Unmarshal(file, character)
	if err != nil {
		log.Log.Println("unmarshal error: ", err.Error())
		return nil
	}
	character.postMarshal()
	log.Log.Printf("parsed %s:\n%+v", path, character)
	return character
}

func (c *Character) RecieveBuffs(buffMap *BuffMap) {
	for k, vec := range *buffMap {
		c.RecievedBuffs[k] = append(c.RecievedBuffs[k], vec...)
	}
}

func (c *Character) EquipWeapon(entity GsEntity) {
	weaponBase := entity.GetBase()
	p := weaponBase.Panel
	c.Base["ba"] += p["ba"]
	c.RecieveBuffs(entity.EmitBuffs())
}

func (c *Character) CalcPanel() {
	c.Final = c.Base
	for _, v := range c.RecievedBuffs["panel"] {
		c.Final[v.Type] += v.V
	}
	//get final
	c.Final["a"] = c.Base["ba"]*(1+c.Base["ra"]) + c.Base["fa"]
	c.Final["h"] = c.Base["bh"]*(1+c.Base["rh"]) + c.Base["fh"]
	c.Final["d"] = c.Base["bd"]*(1+c.Base["rd"]) + c.Base["fd"]
}

func (c *Character) CalcSkill(skillType string) *map[string]float64 {
	skill, ok := c.Skills[skillType]
	if !ok {
		return nil
	}
	resMap := make(map[string]float64)
	//rate zone
	rateVec, ok := c.RecievedBuffs["rateAugment"]
	if ok {
		for _, ra := range rateVec {
			if ra.Type == skillType {
				skill.RawRateSum += ra.V * float64(len(skill.RawRate))
				for k, _ := range skill.Reactions {
					skill.AmplificationSum[k] += ra.V * float64(len(skill.Reactions[k]))
				}
			}
		}
	}

	//(A*(rate+reactionRate*factor)*(1+cr*cd) + upheavalRate*factor) * resistanceFactor
	//ampFactor := AmplificationFactor(c.Final["em"])
	reactionVec, ok := c.RecievedBuffs["reactionAugment"]
	if ok {
		for _, r := range reactionVec {
			if fac, found := skill.ReactionFator[r.Type]; found {
				skill.ReactionFator[r.Type] += fac
			}
		}
	}
	czone := 1 + c.Final["cr"]*c.Final["cd"]/10000
	rzone := skill.RawRateSum
	for k, v := range skill.AmplificationSum {
		rzone += v * BaseReactionFactor(k, skill.ElementType) * skill.ReactionFator[k]
	}
	azone := c.Final["a"] * rzone * czone

	resMap["a"] = azone
	return &resMap
}

func AmplificationFactor(em float64) float64 {
	return 25/9*em/(em+1400) + 1
}

func UpheavalFactor(em float64) float64 {
	return 16*em/(em+2000) + 1
}

func DefenceFactor(dd float64) float64 {
	lv := 90.0
	elv := 100.0
	return (100.0 + lv) / (100.0 + lv + (1.0-dd)*(100.0+elv))
}
