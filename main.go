package main

import (
	"mini-chain/core"
)

func main() {
	bc := core.NewBlockChain()

	cli := CLI{bc}
	cli.Run()
}
