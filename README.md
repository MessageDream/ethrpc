# Ethrpc
[![Build Status](https://travis-ci.org/MessageDream/ethrpc.svg?branch=master)](https://travis-ci.org/MessageDream/ethrpc)
[![Coverage Status](https://coveralls.io/repos/github/MessageDream/ethrpc/badge.svg?branch=master)](https://coveralls.io/github/MessageDream/ethrpc?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/MessageDream/ethrpc)](https://goreportcard.com/report/github.com/MessageDream/ethrpc)
[![GoDoc](https://godoc.org/github.com/MessageDream/ethrpc?status.svg)](https://godoc.org/github.com/MessageDream/ethrpc)
[![Donate with Ethereum](https://en.cryptobadges.io/badge/micro/0xf4144308d6D67A1F00a61A596c0eB7B08411344a)](https://en.cryptobadges.io/donate/0xf4144308d6D67A1F00a61A596c0eB7B08411344a)

this library is forked from [onrik/ethrpc](https://github.com/onrik/ethrpc)

## Methods supported (so far...)

### Web3

- [x] [web3_clientVersion](https://wiki.parity.io/JSONRPC-web3-module#web3_clientVersion)
- [x] [web3_sha3](https://wiki.parity.io/JSONRPC-web3-module#web3_sha3)

### Net

- [x] [net_version](https://wiki.parity.io/JSONRPC-net-module#net_version)
- [x] [net_peerCount](https://wiki.parity.io/JSONRPC-net-module#net_peerCount)
- [x] [net_listening](https://wiki.parity.io/JSONRPC-net-module#net_listening)

### Eth

- [x] [eth_protocolVersion](https://wiki.parity.io/JSONRPC-eth-module#eth_protocolversion)
- [x] [eth_coinbase](https://wiki.parity.io/JSONRPC-eth-module#eth_coinbase)
- [x] [eth_mining](https://wiki.parity.io/JSONRPC-eth-module#eth_mining)
- [x] [eth_hashrate](https://wiki.parity.io/JSONRPC-eth-module#eth_hashrate)
- [x] [eth_gasPrice](https://wiki.parity.io/JSONRPC-eth-module#eth_gasprice)
- [x] [eth_accounts](https://wiki.parity.io/JSONRPC-eth-module#eth_accounts)
- [x] [eth_blockNumber](https://wiki.parity.io/JSONRPC-eth-module#eth_blocknumber)
- [x] [eth_getBalance](https://wiki.parity.io/JSONRPC-eth-module#eth_getbalance)
- [x] [eth_getStorageAt](https://wiki.parity.io/JSONRPC-eth-module#eth_getstorageAt)
- [x] [eth_getTransactionCount](https://wiki.parity.io/JSONRPC-eth-module#eth_gettransactioncount)
- [x] [eth_getBlockTransactionCountByHash](https://wiki.parity.io/JSONRPC-eth-module#eth_getblocktransactioncountbyhash)
- [x] [eth_getBlockTransactionCountByNumber](https://wiki.parity.io/JSONRPC-eth-module#eth_getblocktransactioncountbynumber)
- [x] [eth_getUncleCountByBlockHash](https://wiki.parity.io/JSONRPC-eth-module#eth_getunclecountbyblockhash)
- [x] [eth_getUncleCountByBlockNumber](https://wiki.parity.io/JSONRPC-eth-module#eth_getunclecountbyblocknumber)
- [x] [eth_getCode](https://wiki.parity.io/JSONRPC-eth-module#eth_getcode)
- [x] [eth_sign](https://wiki.parity.io/JSONRPC-eth-module#eth_sign)
- [x] [eth_sendTransaction](https://wiki.parity.io/JSONRPC-eth-module#eth_sendtransaction)
- [x] [eth_sendRawTransaction](https://wiki.parity.io/JSONRPC-eth-module#eth_sendrawtransaction)
- [x] [eth_call](https://wiki.parity.io/JSONRPC-eth-module#eth_call)
- [x] [eth_estimateGas](https://wiki.parity.io/JSONRPC-eth-module#eth_estimategas)
- [x] [eth_getBlockByHash](https://wiki.parity.io/JSONRPC-eth-module#eth_getblockbyhash)
- [x] [eth_getBlockByNumber](https://wiki.parity.io/JSONRPC-eth-module#eth_getblockbynumber)
- [x] [eth_getTransactionByHash](https://wiki.parity.io/JSONRPC-eth-module#eth_gettransactionbyhash)
- [x] [eth_getTransactionByBlockHashAndIndex](https://wiki.parity.io/JSONRPC-eth-module#eth_gettransactionbyblockhashandindex)
- [x] [eth_getTransactionByBlockNumberAndIndex](https://wiki.parity.io/JSONRPC-eth-module#eth_gettransactionbyblocknumberandindex)
- [x] [eth_getTransactionReceipt](https://wiki.parity.io/JSONRPC-eth-module#eth_gettransactionreceipt)
- [ ] [eth_getUncleByBlockHashAndIndex](https://wiki.parity.io/JSONRPC-eth-module#eth_getunclebyblockhashandindex)
- [ ] [eth_getUncleByBlockNumberAndIndex](https://wiki.parity.io/JSONRPC-eth-module#eth_getunclebyblocknumberandindex)
- [x] [eth_getCompilers](https://wiki.parity.io/JSONRPC-eth-module#eth_getcompilers)
- [ ] [eth_compileLLL](https://wiki.parity.io/JSONRPC-eth-module#eth_compilelll)
- [ ] [eth_compileSolidity](https://wiki.parity.io/JSONRPC-eth-module#eth_compilesolidity)
- [ ] [eth_compileSerpent](https://wiki.parity.io/JSONRPC-eth-module#eth_compileserpent)
- [x] [eth_newFilter](https://wiki.parity.io/JSONRPC-eth-module#eth_newfilter)
- [x] [eth_newBlockFilter](https://wiki.parity.io/JSONRPC-eth-module#eth_newblockfilter)
- [x] [eth_newPendingTransactionFilter](https://wiki.parity.io/JSONRPC-eth-module#eth_newpendingtransactionfilter)
- [x] [eth_uninstallFilter](https://wiki.parity.io/JSONRPC-eth-module#eth_uninstallfilter)
- [x] [eth_getFilterChanges](https://wiki.parity.io/JSONRPC-eth-module#eth_getfilterchanges)
- [x] [eth_getFilterLogs](https://wiki.parity.io/JSONRPC-eth-module#eth_getfilterlogs)
- [x] [eth_getLogs](https://wiki.parity.io/JSONRPC-eth-module#eth_getlogs)
- [ ] [eth_getWork](https://wiki.parity.io/JSONRPC-eth-module#eth_getwork)
- [ ] [eth_submitWork](https://wiki.parity.io/JSONRPC-eth-module#eth_submitwork)
- [ ] [eth_submitHashrate](https://wiki.parity.io/JSONRPC-eth-module#eth_submitashrate)

### shh

- [ ] [shh_post](https://wiki.parity.io/JSONRPC-shh-module#shh_post)
- [ ] [shh_version](https://wiki.parity.io/JSONRPC-shh-module#shh_version)
- [ ] [shh_newIdentity](https://wiki.parity.io/JSONRPC-shh-module#shh_newIdentity)
- [ ] [shh_hasIdentity](https://wiki.parity.io/JSONRPC-shh-module#shh_hasIdentity)
- [ ] [shh_newGroup](https://wiki.parity.io/JSONRPC-shh-module#shh_newGroup)
- [ ] [shh_addToGroup](https://wiki.parity.io/JSONRPC-shh-module#shh_addToGroup)
- [ ] [shh_newFilter](https://wiki.parity.io/JSONRPC-shh-module#shh_newFilter)
- [ ] [shh_uninstallFilter](https://wiki.parity.io/JSONRPC-shh-module#shh_uninstallFilter)
- [ ] [shh_getFilterChanges](https://wiki.parity.io/JSONRPC-shh-module#shh_getFilterChanges)
- [ ] [shh_getMessages](https://wiki.parity.io/JSONRPC-shh-module#shh_getMessages)


### Personal

- [ ] [personal_listAccounts](https://wiki.parity.io/JSONRPC-personal-module#personal_listaccounts)
- [ ] [personal_newAccount](https://wiki.parity.io/JSONRPC-personal-module#personal_newaccount)
- [ ] [personal_sendTransaction](https://wiki.parity.io/JSONRPC-personal-module#personal_sendtransaction)
- [ ] [personal_signTransaction](https://wiki.parity.io/JSONRPC-personal-module#personal_signtransaction)
- [ ] [personal_unlockAccount](https://wiki.parity.io/JSONRPC-personal-module#personal_unlockaccount)
- [ ] [personal_sign](https://wiki.parity.io/JSONRPC-personal-module#personal_sign)
- [ ] [personal_ecRecover](https://wiki.parity.io/JSONRPC-personal-module#personal_ecrecover)


##### Usage:
```go
package main

import (
    "fmt"
    "log"

    "github.com/MessageDream/ethrpc"
)

func main() {
    client := ethrpc.NewClient("http://127.0.0.1:8545")

    // client.Module.Function(...params...) // e.g. client.Eth.GetBalance("0x00000000000000000001")

    version, err := client.Web3.ClientVersion()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(version)

    // Send 1 eth
    txid, err := client.Eth.SendTransaction(ethrpc.T{
        From:  "0x6247cf0412c6462da2a51d05139e2a3c6c630f0a",
        To:    "0xcfa202c4268749fbb5136f2b68f7402984ed444b",
        Value: ethrpc.Eth1(),
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(txid)
}
```

[![Donate with Ethereum](https://en.cryptobadges.io/badge/big/0x7396c97d72807c692f78240e7c82b00738fa0517?showBalance=true)](https://en.cryptobadges.io/donate/0x7396c97d72807c692f78240e7c82b00738fa0517)

