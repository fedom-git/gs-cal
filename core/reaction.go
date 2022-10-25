package core

import "strings"

const (
	UpheavalBaseDamage = 723.0
)

var (
	AmplificationFactorMap = map[string]float64{
		"vaporize":       2,
		"vaporize-hydro": 2,
		"vaporize-pyro":  1.5,
		"melt":           2,
		"melt-pyro":      2,
		"melt-cryo":      1.5,
	}
	UpheavalFactorMap = map[string]float64{
		"superConduct":   1,
		"swirl":          1.2,
		"electroCharged": 2.4,
		"shattered":      3,
		"overload":       4,
	}

	UpheavalDamageMap = map[string]float64{
		"superConduction": UpheavalBaseDamage * 1,
		"swirl":           UpheavalBaseDamage * 1.2,
		"electroCharged":  UpheavalBaseDamage * 2.4,
		"shattered":       UpheavalBaseDamage * 3,
		"overload":        UpheavalBaseDamage * 4,
	}
)

func IsAmplification(k string) bool {
	_, ok := AmplificationFactorMap[k]
	return ok
}
func BaseReactionFactor(rType string, self string) float64 {
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
	case "electroCharged":
		return 2.4
	case "shattered":
		return 3
	case "overload":
		return 4
	default:
		return 1
	}
	return 1
}

func GetUpheavalElementType(u *string) string {
	switch *u {
	case "superConduction":
		return "cryo"
	case "eletroCharged":
		return "dendro"
	case "shattered":
		return "physical"
	case "overload":
		return "pyro"
	default:
		return "allo" //TODO
	}
}
