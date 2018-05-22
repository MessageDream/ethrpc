package ethrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
)

// EthError - ethereum error
type EthError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err EthError) Error() string {
	return fmt.Sprintf("Error %d (%s)", err.Code, err.Message)
}

type ethResponse struct {
	ID      int             `json:"id"`
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   *EthError       `json:"error"`
}

type ethRequest struct {
	ID      int           `json:"id"`
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

// Client - Ethereum rpc client
type Client struct {
	endpoint   string
	httpClient httpClient
	log        logger
	Debug      bool
	Web3       *Web3
	Net        *Net
	Eth        *Eth
	Personal   *Personal
}

// New create new rpc client with given endpoint
func newClient(endpoint string, options ...func(client *Client)) *Client {
	client := &Client{
		endpoint:   endpoint,
		httpClient: http.DefaultClient,
		log:        log.New(os.Stderr, "", log.LstdFlags),
	}

	client.Web3 = &Web3{client}
	client.Net = &Net{client}
	client.Eth = &Eth{client}
	client.Personal = &Personal{client}

	for _, option := range options {
		option(client)
	}

	return client
}

// NewClient create new rpc client with given endpoint
func NewClient(endpoint string, options ...func(rpc *Client)) *Client {
	return newClient(endpoint, options...)
}

func (rpc *Client) call(method string, target interface{}, params ...interface{}) error {
	result, err := rpc.Call(method, params...)
	if err != nil {
		return err
	}

	if target == nil {
		return nil
	}

	return json.Unmarshal(result, target)
}

// Call returns raw response of method call
func (rpc *Client) Call(method string, params ...interface{}) (json.RawMessage, error) {
	request := ethRequest{
		ID:      1,
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	response, err := rpc.httpClient.Post(rpc.endpoint, "application/json", bytes.NewBuffer(body))
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if rpc.Debug {
		rpc.log.Println(fmt.Sprintf("%s\nRequest: %s\nResponse: %s\n", method, body, data))
	}

	resp := new(ethResponse)
	if err := json.Unmarshal(data, resp); err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, *resp.Error
	}

	return resp.Result, nil

}

// RawCall returns raw response of method call (Deprecated)
func (rpc *Client) RawCall(method string, params ...interface{}) (json.RawMessage, error) {
	return rpc.Call(method, params...)
}

// Eth1 returns 1 ethereum value (10^18 wei)
func (rpc *Client) Eth1() *big.Int {
	return Eth1()
}

// Eth1 returns 1 ethereum value (10^18 wei)
func Eth1() *big.Int {
	return big.NewInt(1000000000000000000)
}

type Web3 struct {
	client *Client
}

// Web3ClientVersion returns the current client version.
func (web *Web3) ClientVersion() (string, error) {
	var clientVersion string

	err := web.client.call("web3_clientVersion", &clientVersion)
	return clientVersion, err
}

// Web3Sha3 returns Keccak-256 (not the standardized SHA3-256) of the given data.
func (web *Web3) Sha3(data []byte) (string, error) {
	var hash string

	err := web.client.call("web3_sha3", &hash, fmt.Sprintf("0x%x", data))
	return hash, err
}

type Net struct {
	client *Client
}

// NetVersion returns the current network protocol version.
func (net *Net) Version() (string, error) {
	var version string

	err := net.client.call("net_version", &version)
	return version, err
}

// NetListening returns true if client is actively listening for network connections.
func (net *Net) Listening() (bool, error) {
	var listening bool

	err := net.client.call("net_listening", &listening)
	return listening, err
}

// NetPeerCount returns number of peers currently connected to the client.
func (net *Net) PeerCount() (int, error) {
	var response string
	if err := net.client.call("net_peerCount", &response); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

type Eth struct {
	client *Client
}

// EthProtocolVersion returns the current ethereum protocol version.
func (eth *Eth) ProtocolVersion() (string, error) {
	var protocolVersion string

	err := eth.client.call("eth_protocolVersion", &protocolVersion)
	return protocolVersion, err
}

// EthSyncing returns an object with data about the sync status or false.
func (eth *Eth) Syncing() (*Syncing, error) {
	result, err := eth.client.RawCall("eth_syncing")
	if err != nil {
		return nil, err
	}
	syncing := new(Syncing)
	if bytes.Equal(result, []byte("false")) {
		return syncing, nil
	}
	err = json.Unmarshal(result, syncing)
	return syncing, err
}

// EthCoinbase returns the client coinbase address
func (eth *Eth) Coinbase() (string, error) {
	var address string

	err := eth.client.call("eth_coinbase", &address)
	return address, err
}

// EthMining returns true if client is actively mining new blocks.
func (eth *Eth) Mining() (bool, error) {
	var mining bool

	err := eth.client.call("eth_mining", &mining)
	return mining, err
}

// EthHashrate returns the number of hashes per second that the node is mining with.
func (eth *Eth) Hashrate() (int, error) {
	var response string

	if err := eth.client.call("eth_hashrate", &response); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

// EthGasPrice returns the current price per gas in wei.
func (eth *Eth) GasPrice() (big.Int, error) {
	var response string
	if err := eth.client.call("eth_gasPrice", &response); err != nil {
		return big.Int{}, err
	}

	return ParseBigInt(response)
}

// EthAccounts returns a list of addresses owned by client.
func (eth *Eth) Accounts() ([]string, error) {
	accounts := []string{}

	err := eth.client.call("eth_accounts", &accounts)
	return accounts, err
}

// EthBlockNumber returns the number of most recent block.
func (eth *Eth) BlockNumber() (int, error) {
	var response string
	if err := eth.client.call("eth_blockNumber", &response); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

// EthGetBalance returns the balance of the account of given address in wei.
func (eth *Eth) GetBalance(address, block string) (big.Int, error) {
	var response string
	if err := eth.client.call("eth_getBalance", &response, address, block); err != nil {
		return big.Int{}, err
	}

	return ParseBigInt(response)
}

// EthGetStorageAt returns the value from a storage position at a given address.
func (eth *Eth) GetStorageAt(data string, position int, tag string) (string, error) {
	var result string

	err := eth.client.call("eth_getStorageAt", &result, data, IntToHex(position), tag)
	return result, err
}

// EthGetTransactionCount returns the number of transactions sent from an address.
func (eth *Eth) GetTransactionCount(address, block string) (int, error) {
	var response string

	if err := eth.client.call("eth_getTransactionCount", &response, address, block); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

// EthGetBlockTransactionCountByHash returns the number of transactions in a block from a block matching the given block hash.
func (eth *Eth) GetBlockTransactionCountByHash(hash string) (int, error) {
	var response string

	if err := eth.client.call("eth_getBlockTransactionCountByHash", &response, hash); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

// EthGetBlockTransactionCountByNumber returns the number of transactions in a block from a block matching the given block
func (eth *Eth) GetBlockTransactionCountByNumber(number int) (int, error) {
	var response string

	if err := eth.client.call("eth_getBlockTransactionCountByNumber", &response, IntToHex(number)); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

// EthGetUncleCountByBlockHash returns the number of uncles in a block from a block matching the given block hash.
func (eth *Eth) GetUncleCountByBlockHash(hash string) (int, error) {
	var response string

	if err := eth.client.call("eth_getUncleCountByBlockHash", &response, hash); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

// EthGetUncleCountByBlockNumber returns the number of uncles in a block from a block matching the given block number.
func (eth *Eth) GetUncleCountByBlockNumber(number int) (int, error) {
	var response string

	if err := eth.client.call("eth_getUncleCountByBlockNumber", &response, IntToHex(number)); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

// EthGetCode returns code at a given address.
func (eth *Eth) GetCode(address, block string) (string, error) {
	var code string

	err := eth.client.call("eth_getCode", &code, address, block)
	return code, err
}

// EthSign signs data with a given address.
// Calculates an Ethereum specific signature with: sign(keccak256("\x19Ethereum Signed Message:\n" + len(message) + message)))
func (eth *Eth) Sign(address, data string) (string, error) {
	var signature string

	err := eth.client.call("eth_sign", &signature, address, data)
	return signature, err
}

// EthSendTransaction creates new message call transaction or a contract creation, if the data field contains code.
func (eth *Eth) SendTransaction(transaction T) (string, error) {
	var hash string

	err := eth.client.call("eth_sendTransaction", &hash, transaction)
	return hash, err
}

// EthSendRawTransaction creates new message call transaction or a contract creation for signed transactions.
func (eth *Eth) SendRawTransaction(data string) (string, error) {
	var hash string

	err := eth.client.call("eth_sendRawTransaction", &hash, data)
	return hash, err
}

// EthCall executes a new message call immediately without creating a transaction on the block chain.
func (eth *Eth) Call(transaction T, tag string) (string, error) {
	var data string

	err := eth.client.call("eth_call", &data, transaction, tag)
	return data, err
}

// EthEstimateGas makes a call or transaction, which won't be added to the blockchain and returns the used gas, which can be used for estimating the used gas.
func (eth *Eth) EstimateGas(transaction T) (int, error) {
	var response string

	err := eth.client.call("eth_estimateGas", &response, transaction)
	if err != nil {
		return 0, err
	}

	return ParseInt(response)
}

func (eth *Eth) getBlock(method string, withTransactions bool, params ...interface{}) (*Block, error) {
	var response proxyBlock
	if withTransactions {
		response = new(proxyBlockWithTransactions)
	} else {
		response = new(proxyBlockWithoutTransactions)
	}

	err := eth.client.call(method, response, params...)
	if err != nil {
		return nil, err
	}
	block := response.toBlock()

	return &block, nil
}

// EthGetBlockByHash returns information about a block by hash.
func (eth *Eth) GetBlockByHash(hash string, withTransactions bool) (*Block, error) {
	return eth.getBlock("eth_getBlockByHash", withTransactions, hash, withTransactions)
}

// EthGetBlockByNumber returns information about a block by block number.
func (eth *Eth) GetBlockByNumber(number int, withTransactions bool) (*Block, error) {
	return eth.getBlock("eth_getBlockByNumber", withTransactions, IntToHex(number), withTransactions)
}

func (eth *Eth) getTransaction(method string, params ...interface{}) (*Transaction, error) {
	transaction := new(Transaction)

	err := eth.client.call(method, transaction, params...)
	return transaction, err
}

// EthGetTransactionByHash returns the information about a transaction requested by transaction hash.
func (eth *Eth) GetTransactionByHash(hash string) (*Transaction, error) {
	return eth.getTransaction("eth_getTransactionByHash", hash)
}

// EthGetTransactionByBlockHashAndIndex returns information about a transaction by block hash and transaction index position.
func (eth *Eth) GetTransactionByBlockHashAndIndex(blockHash string, transactionIndex int) (*Transaction, error) {
	return eth.getTransaction("eth_getTransactionByBlockHashAndIndex", blockHash, IntToHex(transactionIndex))
}

// EthGetTransactionByBlockNumberAndIndex returns information about a transaction by block number and transaction index position.
func (eth *Eth) GetTransactionByBlockNumberAndIndex(blockNumber, transactionIndex int) (*Transaction, error) {
	return eth.getTransaction("eth_getTransactionByBlockNumberAndIndex", IntToHex(blockNumber), IntToHex(transactionIndex))
}

// EthGetTransactionReceipt returns the receipt of a transaction by transaction hash.
// Note That the receipt is not available for pending transactions.
func (eth *Eth) GetTransactionReceipt(hash string) (*TransactionReceipt, error) {
	transactionReceipt := new(TransactionReceipt)

	err := eth.client.call("eth_getTransactionReceipt", transactionReceipt, hash)
	if err != nil {
		return nil, err
	}

	return transactionReceipt, nil
}

// EthGetCompilers returns a list of available compilers in the client.
func (eth *Eth) GetCompilers() ([]string, error) {
	compilers := []string{}

	err := eth.client.call("eth_getCompilers", &compilers)
	return compilers, err
}

// EthNewFilter creates a new filter object.
func (eth *Eth) NewFilter(params FilterParams) (string, error) {
	var filterID string
	err := eth.client.call("eth_newFilter", &filterID, params)
	return filterID, err
}

// EthNewBlockFilter creates a filter in the node, to notify when a new block arrives.
// To check if the state has changed, call EthGetFilterChanges.
func (eth *Eth) NewBlockFilter() (string, error) {
	var filterID string
	err := eth.client.call("eth_newBlockFilter", &filterID)
	return filterID, err
}

// EthNewPendingTransactionFilter creates a filter in the node, to notify when new pending transactions arrive.
// To check if the state has changed, call EthGetFilterChanges.
func (eth *Eth) NewPendingTransactionFilter() (string, error) {
	var filterID string
	err := eth.client.call("eth_newPendingTransactionFilter", &filterID)
	return filterID, err
}

// EthUninstallFilter uninstalls a filter with given id.
func (eth *Eth) UninstallFilter(filterID string) (bool, error) {
	var res bool
	err := eth.client.call("eth_uninstallFilter", &res, filterID)
	return res, err
}

// EthGetFilterChanges polling method for a filter, which returns an array of logs which occurred since last poll.
func (eth *Eth) GetFilterChanges(filterID string) ([]Log, error) {
	var logs = []Log{}
	err := eth.client.call("eth_getFilterChanges", &logs, filterID)
	return logs, err
}

// EthGetFilterLogs returns an array of all logs matching filter with given id.
func (eth *Eth) GetFilterLogs(filterID string) ([]Log, error) {
	var logs = []Log{}
	err := eth.client.call("eth_getFilterLogs", &logs, filterID)
	return logs, err
}

// EthGetLogs returns an array of all logs matching a given filter object.
func (eth *Eth) GetLogs(params FilterParams) ([]Log, error) {
	var logs = []Log{}
	err := eth.client.call("eth_getLogs", &logs, params)
	return logs, err
}

type Personal struct {
	rpc *Client
}
