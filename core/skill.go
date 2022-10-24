package core

import (
	"github.com/fedom-git/gs-cal/util/log"
	"strings"
)

func ParseSkill(m *map[string]Skill) {
	for k, s := range *m {
		switch s.SkillType {

		case "damage":
			s.RawRateSum = 0
			for _, v := range s.RawRate {
				s.RawRateSum += v
			}
			for k, vec := range s.Reactions {
				s.ReactionFator[k] = 1
				for _, v := range vec {
					s.AmplificationSum[k] += v
				}
			}
		case "buff":

		}

		log.Log.Println(log.Indent1, "calculated ", k, s)
	}
}

func BaseReactionFactor(rType string, self string) float64 { //1.5 or 2
	switch rType {
	case "vaporize":
		if strings.Compare(self, "hydro") == 0 {
			return 2
		}
		return 1.5 //must be pyro, if input right
	case "melt":
		if strings.Compare(self, "pyro") == 0 {
			return 2
		}
		return 1.5
	case "superConduction":
		return 1
	case "diffusion":
		return 1.2
	case "electricShock":
		return 2.4
	case "iceBreak":
		return 3
	case "overload":
		return 4
	default:
		return 1
	}
	return 1
}
