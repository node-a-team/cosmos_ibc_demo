package txMonitor

import (
	"fmt"
	"time"

	block "multihop/function/block"
	txs "multihop/function/txs"

	t "multihop/types"
)

var ()

func Start(senderRpcServer string, senderRestServer string) {

	defer t.Wait.Done()

	fmt.Printf("------ %s Monitor Start!\n", senderRpcServer)
	currentBlockHeight, previousBlockHeight := "0", "0"

	for {
		blockStatus := block.BlockStatus(senderRpcServer)
		currentBlockHeight = blockStatus.Result.Block.Header.Height

		if currentBlockHeight != previousBlockHeight {
			previousBlockHeight = currentBlockHeight

			senderChainId := blockStatus.Result.Block.Header.Chain_id

			fmt.Printf("[Sender Chain: %s]\tBlockHeight: %s\n", senderChainId, currentBlockHeight)

			txsNum := len(blockStatus.Result.Block.Data.Txs)
			if txsNum != 0 {
				for _, tx := range blockStatus.Result.Block.Data.Txs {
					go txs.TxSearch(tx, senderChainId, senderRpcServer, senderRestServer)
				}
			}

		}

		time.Sleep(1 * time.Second)
	}

}
