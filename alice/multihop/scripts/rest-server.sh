#!/bin/bash

#### use: ./rest-server [laddr port]
## ex: ./rest-server 1317

#nohup gaiacli --home=../../a0/gaiacli rest-server --laddr=tcp://localhost:$1 2>&1 >> ../rest-server.log &
nohup gaiacli --home=../../a0/gaiacli rest-server --laddr=tcp://localhost:$1 > /dev/null &

curl -s localhost:$1/node_info | jq
