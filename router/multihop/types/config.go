package types

import (
	"sync"
)

var (
	Wait sync.WaitGroup

	RpcServer string = "localhost:26657"
	RestServer string = "localhost:2317"

	GaiacliDirectory string = "../r0/gaiacli"
//	Node1 string = "localhost:26657"
	Node2 string = "localhost:16657"
	From1 string = "r0"
	From2 string = "a0"


)

