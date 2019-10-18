package ibc

import (
	"fmt"
	"os/exec"
	"strings"

	t "multihop/types"
)

func Send(senderChannel string) {

	fmt.Println("\n\n\n ## Send")

	nextSequence := getNextSequence(senderChannel)

	cmd := "echo 12345678 | gaiacli --home=" + t.GaiacliDirectory + " tx ibcmocksend sequence " + senderChannel + " " + nextSequence + " --from=" + t.From1 + " -o text -y"
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()

	fmt.Println("sendCmd: ", cmd)
	fmt.Println(string(out))
}

func getNextSequence(senderChannel string) string {

	cmd := "echo 12345678 | gaiacli --home=" + t.GaiacliDirectory + " q ibcmocksend next " + senderChannel + " --trust-node -o text"
	out, _ := exec.Command("/bin/bash", "-c", cmd).Output()

	return strings.Trim(string(out), "\n\t\r")

}
