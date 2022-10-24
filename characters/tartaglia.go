package characters

import (
	"fmt"
	"github.com/fedom-git/gs-cal/core"
)

func main() {
	tt := core.ParseCharacter("DB/characters/tartaglia.yml")
	fmt.Println(tt)
}
