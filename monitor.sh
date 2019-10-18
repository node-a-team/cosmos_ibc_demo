#!/bin/bash

while true; do clear; printf "\n######################################################################\n" && printf "\t\tSender\tRouter_Receive\tRouter_Send\tReceiver\n" && printf "alice -> bob:\t%d\t%d\t\t%d\t\t%d\n" $(gaiacli --home alice/a0/gaiacli q ibcmocksend sequence chan_a-r_for-b --trust-node)  $(gaiacli --home route/r0/gaiacli q ibcmockrecv sequence chan_r-a_for-b --trust-node) $(gaiacli --home route/r0/gaiacli q ibcmocksend sequence chan_r-b_for-a --trust-node) $(gaiacli --home bob/b0/gaiacli q ibcmockrecv sequence chan_b-r_for-a --trust-node) && printf "alice -> carol:\t%d\t%d\t\t%d\t\t%d\n" $(gaiacli --home alice/a0/gaiacli q ibcmocksend sequence chan_a-r_for-c --trust-node)  $(gaiacli --home route/r0/gaiacli q ibcmockrecv sequence chan_r-a_for-c --trust-node) $(gaiacli --home route/r0/gaiacli q ibcmocksend sequence chan_r-c_for-a --trust-node) $(gaiacli --home carol/c0/gaiacli q ibcmockrecv sequence chan_c-r_for-a --trust-node) && printf "######################################################################\n"; sleep 3; done