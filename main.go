package main

import (
	"github.com/fedom-git/gs-cal/core"
	"github.com/fedom-git/gs-cal/util"
	"github.com/fedom-git/gs-cal/util/log"
)

func main() {
	log.InitLog()
	//tt := core.ParseCharacter("DB/characters/yoimiya.yml")
	//log.Log.Println(*tt)
	rust := &core.Weapon{}
	util.ParseFromYaml(core.GetWeaponPath("rust"), rust)
}
