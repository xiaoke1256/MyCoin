package main

import (
	"fmt"

	_ "xiaoke1256.com/mycoin/db"
	"xiaoke1256.com/mycoin/miner"
)

func init() {
	fmt.Println("init has excuted")
}

func main() {
	miner.Mine()
}
