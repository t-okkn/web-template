package main

import (
	"flag"
	"fmt"
)

const LISTEN_PORT string = ":8501"

var (
	Version string
	Revision string
)

// <summary>: main関数（サーバを開始します）
func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		fmt.Println(Version, Revision)
		return
	}

	SetupRouter().Run(LISTEN_PORT)
}
