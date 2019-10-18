package ibc

import (
//	"fmt"
	"os/exec"

//	block "multihop/function/block"

)

var (
	home string = "/data_kusama/node/ibc-testnets/ibc_b/b0/gaiacli"
)

func QueryIbcmockrecv(channel string) string {

	cmd := "gaiacli --home=" +home +" q ibcmockrecv sequence "+channel +" --trust-node"
        out, _ := exec.Command("/bin/bash", "-c", cmd).Output()

	return string(out)
}
