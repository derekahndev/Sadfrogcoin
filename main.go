package main

import (
	"github.com/derekahndev/sadfrogcoin/blockchain"
	"github.com/derekahndev/sadfrogcoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
