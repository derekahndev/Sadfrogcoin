package main

import (
	"github.com/derekahndev/sadfrogcoin/cli"
	"github.com/derekahndev/sadfrogcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
