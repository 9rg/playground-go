package main

import (
	"github.com/9rg/playground-go/properness/internal/app/unluckMain"
)

func main() {
	u := unluckMain.ClosureUnluck("taro")
	u(500)
	u("PC")
	u(1500.5)
}
