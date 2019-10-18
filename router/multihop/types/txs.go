package types

import (
	"time"
)


type Txs struct {
	Height		string
	Txhash		string
	Gas_wanted	string
	Gas_used	string
	Tx		TxStruct

}

type TxStruct struct {
	Type		string
	Value		TxValue
	Timestamp	time.Time
}

type TxValue struct {
	Msg	[]TxMsg
	Memo	string
}

type TxMsg struct {
	Type	string
	Value	MsgValue
}

type MsgValue struct {
	Sequence	string
	ChannelID	string
	Signer		string
}

