Implementing interaction with SimpleStarage smart contract through golang in private blockchain.  
Creating private network:
1) Create several accounts:
geth account new
2) Create genesis.json file to initialize network:
puppeth
3) Initialize the genesis file:
geth init genesis.json
    -if we have an error:
    delete existing database - geth removedb
    -if id doesn't help - remove the folder /home/taketa/.ethereum/geth/chaindata
4) run node:
geth --networkid 123 --rpc console
5) copy the node address and put into it our ip (exmp: enode://6ccd34b8f01c3ae3240eee1467dbf97217c85b866c2756e8eb1ddd67377f4c6209afff709e908b6863e44cf48b80ba215894da3931aa60f1d5a1bcb9fac239a8@192.168.88.62:30303)

Connecting peers to node:
1) initialize genesis.json:
geth init genesis.json
2) connect to node:
geth --networkid 123 --bootnodes enode://6ccd34b8f01c3ae3240eee1467dbf97217c85b866c2756e8eb1ddd67377f4c6209afff709e908b6863e44cf48b80ba215894da3931aa60f1d5a1bcb9fac239a8@192.168.88.62:30303 console

Check if peers connect:
admin.peers
