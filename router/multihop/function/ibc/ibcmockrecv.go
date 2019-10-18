package ibc

import (
	"fmt"
	"os/exec"

	t "multihop/types"
)

func Receive(senderRpcServer string, senderChainId string, receiverChannel string) {

	fmt.Println("\n\n\n ## Receive")

	cmd := "echo 12345678 | gaiacli --home=" + t.GaiacliDirectory + " tx ibc channel pull ibcmockrecv " + receiverChannel + " --node1=tcp://" + t.RpcServer + " --node2=tcp://" + senderRpcServer + " --chain-id2=" + senderChainId + " --from=" + t.From1
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()

	fmt.Println("receivCmd: ", cmd)
	fmt.Println(string(out))
}
