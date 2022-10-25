package core

import (
	"github.com/fedom-git/gs-cal/util/log"
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
					if IsAmplification(k) {
						s.AmplificationSum[k] += v
					} else {
						s.UpheavalCount[k] += v
					}

				}
			}
		case "buff":

		}

		log.Log.Println(log.Indent1, "calculated ", k, s)
	}
}
