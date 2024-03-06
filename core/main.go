package main

import (
	"fmt"

	"log"

	"os"

	"io"

	"xiaoke1256.com/mycoin/miner"
)

func init() {
	fmt.Println("init has excuted")
}

func init() {
	f, err := os.OpenFile("D:/logs/MyCoin.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 07777)
	//defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	multiWriter := io.MultiWriter(os.Stdout, f)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//logger = *log.New(f, "[miner]t", log.Ltime)
	log.Println("日志初始化成功")
}

func main() {
	miner.Mine()
}
