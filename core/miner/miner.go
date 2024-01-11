package miner

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/shopspring/decimal"

	"xiaoke1256.com/mycoin/model"

	"math/rand"

	"xiaoke1256.com/mycoin/crypt"

	"xiaoke1256.com/mycoin/db"

	"errors"

	"fmt"
)

type CoinConfig struct {
	Coinbase      string          /*Coinbase 账号*/
	InitAddresses []string        /* 5个初始账号*/
	InitFund      decimal.Decimal /*初始区块*/
}

var Config CoinConfig

func init() {
	jsonFile, err := os.Open("coinCongfig.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
		panic(err)
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)

	err = decoder.Decode(&Config)
	if err != nil {
		fmt.Println("Cannot get configuration from file")
		panic(err)
	}
	fmt.Println("Config:")
	fmt.Println(Config)
}

func Mine() {
	//看看数据库里有没有区块
	var parentHead *model.CoreBlockheader = db.SearchLastOne("head", "timestamp", model.CoreBlockheader{})
	//没有则挖创世区块
	if parentHead == nil {
		MineForGenesis()
		return
	}
	//有则以现有区块为父区块挖下一个区块。
	MineFromParent(*parentHead)
}

/*
挖创世区块
*/
func MineForGenesis() {
	//构造block
	newBlock := model.CoreBlock{}

	//构造交易(这是个coinbase交易)
	t1 := model.CoreTransaction{}
	t1.Version = "1.0"
	t1.InputCounter = 0
	OutputCount := len(Config.InitAddresses)
	if OutputCount <= 0 {
		panic(errors.New("请设置初始账号"))
	}
	if OutputCount > 5 {
		OutputCount = 5
	}
	t1.OutputCounter = uint32(OutputCount)

	outputs := make([]model.CoreOutput, OutputCount, OutputCount)
	for i := 0; i < OutputCount; i++ {
		output := model.CoreOutput{}
		output.Amt = Config.InitFund
		output.LockScriptSize = uint32(len(Config.InitAddresses[i]))
		output.LockScript = Config.InitAddresses[i]
		outputs[i] = output
	}
	t1.Outputs = outputs
	t1.LockTime = time.Now()

	newBlock.TransactionCounter = 1
	ts := []model.CoreTransaction{t1}
	newBlock.Transactions = ts

	//构造head
	head := model.CoreBlockheader{}
	head.Version = "1.0"
	head.ParentHeadHash = crypt.DoubleSha256([]byte{'G', 'E', 'N', 'E', 'S', 'I', 'S'}) //创世区块的父区块是个默认值
	head.TransactionsMerkleRoot = newBlock.GetTransactionMerkleRoot()
	head.Timestamp = time.Now()
	head.Target = [4]byte{0x00, 0x00, 0x8F, 0xFF}

	newBlock.Blockheader = head

	digging(newBlock)

}

/*
* 挖呀挖呀挖
 */
func digging(newBlock model.CoreBlock) {
	head := newBlock.Blockheader
	lastSeek := time.Now().UnixNano()
	rand.Seed(lastSeek)
	//挖吖挖吖挖
	for true {
		// if time.Now().UnixNano() > lastSeek+10 {
		// 	lastSeek = time.Now().UnixNano()
		// 	rand.Seed(lastSeek)
		// }
		randNum1 := rand.Intn(2 ^ 16)
		randNum2 := rand.Intn(2 ^ 16)
		randNum3 := rand.Intn(2 ^ 16)
		randNum4 := rand.Intn(2 ^ 16)
		var bytes4 [4]byte = [4]byte{byte(randNum1), byte(randNum2), byte(randNum3), byte(randNum4)}
		head.Nonce = bytes4
		//校验是否满足target
		hashedBytes := crypt.Md5(head.ToBytes())
		fmt.Printf("hashedBytes: %x \n", hashedBytes)
		var hashedBytes4 [4]byte
		copy(hashedBytes4[:4], hashedBytes[0:4])
		fmt.Printf("hashedBytes4: %x \n", hashedBytes4)
		if compareBytes(hashedBytes4, head.Target) < 0 {
			//满足需求
			break
		}
	}
	//挖出来了就保存区块
	db.Connect()
	//db.Insert("block", newBlock)
	db.Insert("head", newBlock.Blockheader)
	db.Insert("body", model.Body{crypt.DoubleSha256(newBlock.Blockheader.ToBytes()), newBlock.TransactionCounter, newBlock.Transactions})
}

/*
 * 从父块开始挖
 */
func MineFromParent(parentHead model.CoreBlockheader) {

	//构造block
	newBlock := model.CoreBlock{}

	//获取待打包的交易
	//检查每笔交易的合法性，即有无双花情况
	//创建Coinbase 交易
	t1 := model.CoreTransaction{}
	t1.Version = "1.0"
	t1.InputCounter = 0
	t1.OutputCounter = 1
	outputs := make([]model.CoreOutput, 1, 1)
	output := model.CoreOutput{}
	output.Amt = Config.InitFund
	output.LockScriptSize = uint32(len(Config.Coinbase))
	output.LockScript = Config.Coinbase
	outputs[0] = output

	newBlock.TransactionCounter = 1
	ts := []model.CoreTransaction{t1}
	newBlock.Transactions = ts

	//构造新的Head
	head := model.CoreBlockheader{}
	head.Version = "1.0"
	head.ParentHeadHash = crypt.DoubleSha256(parentHead.ToBytes()) //父区块的Hash
	head.TransactionsMerkleRoot = newBlock.GetTransactionMerkleRoot()
	head.Timestamp = time.Now()

	//检查前两个区块玩出来的时间差，以决定要不要跳转目标难度
	//parentHead.ParentHeadHash
	head.Target = [4]byte{0x00, 0x00, 0x8F, 0xFF}

	newBlock.Blockheader = head

	//挖呀挖呀挖
	digging(newBlock)
}

/*
两个byte数组比大小
*/
func compareBytes(bytes1 [4]byte, bytes2 [4]byte) int {
	for i := 0; i < 4; i++ {
		if bytes1[i] < bytes2[i] {
			fmt.Println("-1")
			return -1
		} else if bytes1[i] > bytes2[i] {
			fmt.Println("1")
			return 1
		} else {
			continue
		}
	}
	return 0
}
