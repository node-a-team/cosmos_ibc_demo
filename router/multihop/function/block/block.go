package block

import (
	t "multihop/types"

	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
)

var (
	preBlockStatus t.BlockStatus
)

func BlockStatus(senderRpc string) t.BlockStatus {

	var blockStatus t.BlockStatus

	cmd := "curl -s " + senderRpc + "/block"
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	json.Unmarshal(out, &blockStatus)

	currentBlockHeight, _ := strconv.Atoi(blockStatus.Result.Block.Header.Height)

	preBlockStatus = previousBlockStatus(senderRpc, currentBlockHeight-1)

	return preBlockStatus

}

func previousBlockStatus(senderRpc string, blockHeight int) t.BlockStatus {

	var blockStatus t.BlockStatus

	cmd := "curl -s " + senderRpc + "/block?height=" + fmt.Sprint(blockHeight)
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	json.Unmarshal(out, &blockStatus)

	return blockStatus

}

func ChainId() string {

	return preBlockStatus.Result.Block.Header.Chain_id
}
