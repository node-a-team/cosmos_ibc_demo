package txs

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func txsConvert(txs string) string {

	// txs -> 검색 가능한 tx hash로 변환
	// base64로decode-> sha256 encode -> hax encode
	txsDecode, _ := base64.StdEncoding.DecodeString(txs)

	txsSha256 := sha256.New()
	txsSha256.Write(txsDecode)

	r := txsSha256.Sum(nil)

	return fmt.Sprintf("%x", r)

}

func TxSearch(txs string, senderChainId string, senderRpcServer string, senderRestServer string) {

	tx := txsConvert(txs)

	IbcTxCheck(tx, senderChainId, senderRpcServer, senderRestServer)

}
