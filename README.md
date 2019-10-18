# Cosmos IBC_Demo Test
- IBC document: <a>https://github.com/cosmos/gaia/tree/joon/ibc-gaia-interface/cmd/gaiacli</a>
- Test video: https://www.youtube.com/watch?v=S6DKib4jINk



# Gaiad Command
<strong>Gaiad start</strong>
- alice chain
```bash
$ gaiad --home=alice/a0/gaiad start
```
- router chain
```bash
$ gaiad --home=router/r0/gaiad start
```
- bob chain
```bash
$ gaiad --home=bob/b0/gaiad start
```
- carol chain
```bash
$ gaiad --home=carol/c0/gaiad start
```

# IBC Command
<strong>Client </strong>
- Creation
```bash
$ echo 12345678 | gaiacli --home alice/a0/gaiacli tx ibc client create client_a-r_for-b $(gaiacli --home router/r0/gaiacli q ibc client consensus-state) --from=a0 -y -o text && sleep 1 && \
echo 12345678 | gaiacli --home alice/a0/gaiacli tx ibc client create client_a-r_for-c $(gaiacli --home router/r0/gaiacli q ibc client consensus-state) --from=a0 -y -o text && sleep 1 && \
echo 12345678 | gaiacli --home router/r0/gaiacli tx ibc client create client_r-a_for-b $(gaiacli --home alice/a0/gaiacli q ibc client consensus-state) --from=r0 -y -o text && sleep 1 && \
echo 12345678 | gaiacli --home router/r0/gaiacli tx ibc client create client_r-a_for-c $(gaiacli --home alice/a0/gaiacli q ibc client consensus-state) --from=r0 -y -o text && sleep 1 && \
echo 12345678 | gaiacli --home router/r0/gaiacli tx ibc client create client_r-b_for-a $(gaiacli --home bob/b0/gaiacli q ibc client consensus-state) --from=r0 -y -o text && sleep 1 && \
echo 12345678 | gaiacli --home router/r0/gaiacli tx ibc client create client_r-c_for-a $(gaiacli --home carol/c0/gaiacli q ibc client consensus-state) --from=r0 -y -o text && sleep 1 && \
echo 12345678 | gaiacli --home bob/b0/gaiacli tx ibc client create client_b-r_for-a $(gaiacli --home router/r0/gaiacli q ibc client consensus-state) --from=b0 -y -o text && sleep 1 && \
echo 12345678 | gaiacli --home carol/c0/gaiacli tx ibc client create client_c-r_for-a $(gaiacli --home router/r0/gaiacli q ibc client consensus-state) --from=c0 -y -o text
```
- Confirm
``` bash
$ gaiacli --home alice/a0/gaiacli q ibc client client client_a-r_for-b --indent && \
gaiacli --home alice/a0/gaiacli q ibc client client client_a-r_for-c --indent && \
gaiacli --home router/r0/gaiacli q ibc client client client_r-a_for-b --indent && \
gaiacli --home router/r0/gaiacli q ibc client client client_r-a_for-c --indent && \
gaiacli --home router/r0/gaiacli q ibc client client client_r-b_for-a --indent && \
gaiacli --home router/r0/gaiacli q ibc client client client_r-c_for-a --indent && \
gaiacli --home bob/b0/gaiacli q ibc client client client_b-r_for-a --indent && \
gaiacli --home carol/c0/gaiacli q ibc client client client_c-r_for-a --indent
```


<strong>Connection</strong>
- Creation
```bash
$ gaiacli --home=alice/a0/gaiacli tx ibc connection handshake \
conn_a-r_for-b client_a-r_for-b $(gaiacli --home router/r0/gaiacli q ibc client path) \
conn_r-a_for-b client_r-a_for-b $(gaiacli --home alice/a0/gaiacli q ibc client path) \
  --chain-id2=router --from1=a0 --from2=r0 --node1=tcp://localhost:16657 --node2=tcp://localhost:26657 && sleep 1 && \
gaiacli --home=alice/a0/gaiacli tx ibc connection handshake \
conn_a-r_for-c client_a-r_for-c $(gaiacli --home router/r0/gaiacli q ibc client path) \
conn_r-a_for-c client_r-a_for-c $(gaiacli --home alice/a0/gaiacli q ibc client path) \
  --chain-id2=router --from1=a0 --from2=r0 --node1=tcp://localhost:16657 --node2=tcp://localhost:26657 && sleep 1 && \
gaiacli --home=router/r0/gaiacli tx ibc connection handshake \
conn_r-b_for-a client_r-b_for-a $(gaiacli --home bob/b0/gaiacli q ibc client path) \
conn_b-r_for-a client_b-r_for-a $(gaiacli --home router/r0/gaiacli q ibc client path) \
  --chain-id2=bob --from1=r0 --from2=b0 --node1=tcp://localhost:26657 --node2=tcp://localhost:36657 && sleep 1 && \
gaiacli --home=router/r0/gaiacli tx ibc connection handshake \
conn_r-c_for-a client_r-c_for-a $(gaiacli --home carol/c0/gaiacli q ibc client path) \
conn_c-r_for-a client_c-r_for-a $(gaiacli --home router/r0/gaiacli q ibc client path) \
  --chain-id2=carol --from1=r0 --from2=c0 --node1=tcp://localhost:26657 --node2=tcp://localhost:46657 
```

- Confirm
```bash
$ gaiacli --home alice/a0/gaiacli q ibc connection connection  conn_a-r_for-b --indent --trust-node && \
gaiacli --home alice/a0/gaiacli q ibc connection connection  conn_a-r_for-c --indent --trust-node && \
gaiacli --home router/r0/gaiacli q ibc connection connection conn_r-a_for-b --indent --trust-node && \
gaiacli --home router/r0/gaiacli q ibc connection connection conn_r-a_for-c --indent --trust-node && \
gaiacli --home router/r0/gaiacli q ibc connection connection conn_r-b_for-a --indent --trust-node && \
gaiacli --home router/r0/gaiacli q ibc connection connection conn_r-c_for-a --indent --trust-node && \
gaiacli --home bob/b0/gaiacli q ibc connection connection conn_b-r_for-a --indent --trust-node && \
gaiacli --home carol/c0/gaiacli q ibc connection connection conn_c-r_for-a --indent --trust-node
```


<strong>Channel </strong>
- Creation
```bash
$ gaiacli --home alice/a0/gaiacli tx ibc channel handshake \
  ibcmocksend chan_a-r_for-b conn_a-r_for-b \
  ibcmockrecv chan_r-a_for-b conn_r-a_for-b \
  --node1=tcp://localhost:16657 --node2=tcp://localhost:26657 --chain-id2=router --from1=a0 --from2=r0 && sleep 1 && \
gaiacli --home alice/a0/gaiacli tx ibc channel handshake \
  ibcmocksend chan_a-r_for-c conn_a-r_for-c \
  ibcmockrecv chan_r-a_for-c conn_r-a_for-c \
  --node1=tcp://localhost:16657 --node2=tcp://localhost:26657 --chain-id2=router --from1=a0 --from2=r0 && sleep 1 && \
gaiacli --home router/r0/gaiacli \
  tx ibc channel handshake \
  ibcmocksend chan_r-b_for-a conn_r-b_for-a \
  ibcmockrecv chan_b-r_for-a conn_b-r_for-a \
  --node1=tcp://localhost:26657 --node2=tcp://localhost:36657 --chain-id2=bob --from1=r0 --from2=b0 && sleep 1 && \
gaiacli --home router/r0/gaiacli \
  tx ibc channel handshake \
  ibcmocksend chan_r-c_for-a conn_r-c_for-a \
  ibcmockrecv chan_c-r_for-a conn_c-r_for-a \
  --node1=tcp://localhost:26657 --node2=tcp://localhost:46657 --chain-id2=carol --from1=r0 --from2=c0 
```

- Confirm
```bash
$ gaiacli --home alice/a0/gaiacli q ibc channel channel ibcmocksend chan_a-r_for-b --indent --trust-node && \
gaiacli --home alice/a0/gaiacli q ibc channel channel ibcmocksend chan_a-r_for-c --indent --trust-node && \
gaiacli --home router/r0/gaiacli q ibc channel channel ibcmockrecv chan_r-a_for-b --indent --trust-node && \
gaiacli --home router/r0/gaiacli q ibc channel channel ibcmockrecv chan_r-a_for-c --indent --trust-node && \
gaiacli --home router/r0/gaiacli q ibc channel channel ibcmocksend chan_r-b_for-a --indent --trust-node && \
gaiacli --home router/r0/gaiacli q ibc channel channel ibcmocksend chan_r-c_for-a --indent --trust-node && \
gaiacli --home bob/b0/gaiacli q ibc channel channel ibcmockrecv chan_b-r_for-a --indent --trust-node && \
gaiacli --home carol/c0/gaiacli q ibc channel channel ibcmockrecv chan_c-r_for-a --indent --trust-node
```


<strong>Send </strong>
- alice -> router(for bob)
```bash
$ echo 12345678 | gaiacli --home alice/a0/gaiacli tx ibcmocksend sequence chan_a-r_for-b $(gaiacli --home alice/a0/gaiacli q ibcmocksend next chan_a-r_for-b --trust-node) --from a0 -o text -y
```

- alice -> router(for carol)
```bash
$echo 12345678 | gaiacli --home alice/a0/gaiacli tx ibcmocksend sequence chan_a-r_for-c $(gaiacli --home alice/a0/gaiacli q ibcmocksend next chan_a-r_for-c --trust-node) --from a0 -o text -y
```

- router -> bob
```bash
$ echo 12345678 | gaiacli --home router/r0/gaiacli tx ibcmocksend sequence chan_r-b_for-a $(gaiacli --home router/r0/gaiacli q ibcmocksend next chan_r-b_for-a --trust-node) --from r0 -o text -y
```

- router -> carol
```bash
$ echo 12345678 | gaiacli --home router/r0/gaiacli tx ibcmocksend sequence chan_r-b_for-c $(gaiacli --home router/r0/gaiacli q ibcmocksend next chan_r-b_for-c --trust-node) --from r0 -o text -y
```

<strong> Receive </strong>
- router(from alice, for bob)
```bash
$ echo 12345678 | gaiacli --home router/r0/gaiacli tx ibc channel pull ibcmockrecv chan_r-a_for-b \
  --node1=tcp://localhost:26657 --node2=tcp://localhost:16657 --chain-id2=alice --from=r0
```
- router(from alice, for carol)
```bash
$ echo 12345678 | gaiacli --home router/r0/gaiacli tx ibc channel pull ibcmockrecv chan_r-a_for-c \
  --node1=tcp://localhost:26657 --node2=tcp://localhost:16657 --chain-id2=alice --from=r0
```
- bob(from router)
```bash
$ echo 12345678 | gaiacli --home bob/b0/gaiacli tx ibc channel pull ibcmockrecv chan_b-r_for-a \
  --node1=tcp://localhost:36657 --node2=tcp://localhost:26657 --chain-id2=router --from=b0
```
- carol(from router)
```bash
$ echo 12345678 | gaiacli --home carol/c0/gaiacli tx ibc channel pull ibcmockrecv chan_c-r_for-a \
  --node1=tcp://localhost:46657 --node2=tcp://localhost:26657 --chain-id2=router --from=c0
```

# Etc Scripts
- alice's rest-server
```bash
$ nohup gaiacli --home alice/a0/gaiacli rest-server --laddr=tcp://localhost:1317 > /dev/null &
```

- Sequence monitor
```bash
$ printf "\n##########################################################################\n" && \
printf "\t\tSender\trouterr_Receive\trouterr_Send\tReceiver\n" && \
printf "alice -> bob:\t%d\t%d\t\t%d\t\t%d\n" $(gaiacli --home alice/a0/gaiacli q ibcmocksend sequence chan_a-r_for-b --trust-node)  $(gaiacli --home router/r0/gaiacli q ibcmockrecv sequence chan_r-a_for-b --trust-node) $(gaiacli --home router/r0/gaiacli q ibcmocksend sequence chan_r-b_for-a --trust-node) $(gaiacli --home bob/b0/gaiacli q ibcmockrecv sequence chan_b-r_for-a --trust-node) && \
printf "alice -> carol:\t%d\t%d\t\t%d\t\t%d\n" $(gaiacli --home alice/a0/gaiacli q ibcmocksend sequence chan_a-r_for-c --trust-node)  $(gaiacli --home router/r0/gaiacli q ibcmockrecv sequence chan_r-a_for-c --trust-node) $(gaiacli --home router/r0/gaiacli q ibcmocksend sequence chan_r-c_for-a --trust-node) $(gaiacli --home carol/c0/gaiacli q ibcmockrecv sequence chan_c-r_for-a --trust-node) && \
printf "##########################################################################\n"; \
sleep 3; done
```
