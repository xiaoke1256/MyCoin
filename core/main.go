package main

import (
	"fmt"

	"xiaoke1256.com/mycoin/miner"
)

func init() {
	fmt.Println("init has excuted")
}

func main() {
	miner.Mine()
}
