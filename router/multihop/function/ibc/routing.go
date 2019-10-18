package ibc

import (
)

func RoutingForReceiving(senderChannel string) (receivingChannel string) {

	switch senderChannel {
	case "chan_a-r_for-b":
		receivingChannel = "chan_r-a_for-b"
	case "chan_a-r_for-c":
                receivingChannel = "chan_r-a_for-c"
	default:
		receivingChannel = ""
	}

	return receivingChannel
}

func RoutingForDestination(senderChannel string) (destinationChannel string) {

        switch senderChannel {
        case "chan_a-r_for-b":
                destinationChannel = "chan_r-b_for-a"
        case "chan_a-r_for-c":
                destinationChannel = "chan_r-c_for-a"
        default:
                destinationChannel = ""
        }

        return destinationChannel
}

