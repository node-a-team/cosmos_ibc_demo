package types

import (
	"time"
)

type BlockStatus struct {
	Result	BlockResult	`json:"result"`
//	Block Block
}

type BlockResult struct {
	Block	Block
}

type Block struct {
	Header		BlockHeader
	Data		BlockData
	Evidence	BlockEvidence
	Last_commit	BlockLastCommit
}

type BlockHeader struct {
	Chain_id		string
	Height			string
	Time			time.Time
	Num_txs			string
	Total_txs		string
	Proposer_address	string
}

type BlockData struct {
	Txs	[]string
}

type BlockEvidence struct {
	Evidence	string
}

type BlockLastCommit struct {
	Precommits	[]PrecommitValidator
}


type PrecommitValidator struct {
	Type			int
	Height			string
	Round			string
	Timestamp		time.Time
	Validator_address	string
	Validator_index		string
	Signature		string
}

