package model

type Body struct {
	HeadHash           [32]byte
	TransactionCounter uint32            //交易计数器（本包中有几个交易）
	Transactions       []CoreTransaction //交易
}
