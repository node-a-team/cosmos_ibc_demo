package main

import (
	txMonitor "multihop/function/txMonitor"

	t "multihop/types"
)

var (
)



func main() {


	t.Wait.Add(2)

	senderRpcServerList := []string{"localhost:16657"}
	senderRestServerList := []string{"localhost:1317"}

	for i, sender := range senderRpcServerList {

		go txMonitor.Start(sender, senderRestServerList[i])
	}

	t.Wait.Wait()



}
