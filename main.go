package main

import (
	"github.com/pdstuber/gameboy-emulator/cmd"
	"github.com/pdstuber/gameboy-emulator/pkg/util"
)

func main() {
	_ = util.X()
	//fmt.Println(tile)
	cmd.Execute()
}
