/*
Copyright © 2024 Hantsaniala Eléo <hantsaniala@gmail.com>
*/
package main

import (
	"fmt"

	"github.com/fussaa/fussaa-cli/cmd"
)

var Version string

const (
	banner = `
░█▀▀░█░█░█▀▀░█▀▀░█▀█░█▀█░░░░░█▀▀░█░░░▀█▀
░█▀▀░█░█░▀▀█░▀▀█░█▀█░█▀█░▄▄▄░█░░░█░░░░█░
░▀░░░▀▀▀░▀▀▀░▀▀▀░▀░▀░▀░▀░░░░░▀▀▀░▀▀▀░▀▀▀
		%s

`
)

func main() {
	fmt.Printf(banner, Version)
	cmd.Execute()
}
