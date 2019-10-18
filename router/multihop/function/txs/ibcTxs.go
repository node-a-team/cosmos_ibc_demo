package txs

import (
	"encoding/json"
	"fmt"
	"os/exec"
//	"strings"

	ibc "multihop/function/ibc"
	t "multihop/types"
)

func IbcTxCheck(tx string, senderChainId string, senderRpcServer string, senderRestServer string) {

	var txStatus t.Txs

	fmt.Printf("!!! TX Check(from %s)\n", senderRpcServer)

	cmd := "curl -s " + senderRestServer + "/txs/" + tx
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()
	json.Unmarshal(out, &txStatus)

	msgLen := len(txStatus.Tx.Value.Msg)

	if txStatus.Tx.Value.Msg[msgLen-1].Type == "ibcmock/MsgSequence" {

		senderChannel := txStatus.Tx.Value.Msg[msgLen-1].Value.ChannelID
		fmt.Printf("\n\n##### IBC Packet from %s\n", senderChannel)

//		senderPacketNum := strings.Replace(txStatus.Tx.Value.Msg[msgLen-1].Value.Sequence, "sequence-packet-", "", -1)
//		latestRecievePacketNum := ibc.QueryIbcmockrecv(senderChannel)

//		if senderPacketNum > latestRecievePacketNum {

			receivingChannel := ibc.RoutingForReceiving(senderChannel)
			fmt.Printf("#### Receiving Channel: %s\n", receivingChannel)
			ibc.Receive(senderRpcServer, senderChainId, receivingChannel)

			destinationChannel := ibc.RoutingForDestination(senderChannel)
			fmt.Printf("#### Send Packet to %s\n", destinationChannel)
			ibc.Send(destinationChannel)

//		} else {
//			fmt.Println("check senderPacketNum..")
//		}
	}
}
