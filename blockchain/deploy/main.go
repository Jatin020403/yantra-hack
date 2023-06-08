// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/ethclient"
// )

// func main() {
// 	client, err := ethclient.Dial("https://polygon-mumbai.g.alchemy.com/v2/-Sw3zWYIhrJKYyKnauBOMZWe3VWUKJSS")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Get the balance of an account
// 	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
// 	balance, err := client.BalanceAt(context.Background(), account, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Account balance:", balance) // 25893180161173005034

// 	// Get the latest known block
// 	block, err := client.BlockByNumber(context.Background(), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Latest block:", block.Number().Uint64())
// }

package main

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

// const (
// 	contractAddress =     // Insert the address of your deployed contract on the Mumbai Testnet here
// 	contractABI     = `[{"constant":true,"inputs":[...}]` // Insert the ABI of your smart contract here
// 	rpcURL          = "https://rpc-mumbai.matic.today"    // Mumbai Testnet RPC endpoint
// )

func connect(string *caller) {

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	contractAddress := os.Getenv("CONTRACT_ADDRESS")

	contractABI := `[
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "string",
				"name": "id",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "location",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "bool",
				"name": "isDamaged",
				"type": "bool"
			}
		],
		"name": "LocationUpdate",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "string",
				"name": "id",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "origin",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "tom",
				"type": "string"
			}
		],
		"name": "NewEntry",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "_id",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "_origin",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "_tom",
				"type": "string"
			}
		],
		"name": "newEntry",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "_id",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "_location",
				"type": "string"
			},
			{
				"internalType": "bool",
				"name": "_isDamaged",
				"type": "bool"
			}
		],
		"name": "updateLocation",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "_id",
				"type": "string"
			}
		],
		"name": "getProduct",
		"outputs": [
			{
				"internalType": "string",
				"name": "origin",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "tom",
				"type": "string"
			},
			{
				"internalType": "string[]",
				"name": "locations",
				"type": "string[]"
			},
			{
				"internalType": "bool",
				"name": "isDamaged",
				"type": "bool"
			}
		],
		"stateMutability": "view",
		"type": "function"
	}
]`
	rpcURL := os.Getenv("POLYGON_RPC_URL")

	accountAddress := os.Getenv("ACCOUNT_ADDRESS")

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal(err)
	}

	contractAddr := common.HexToAddress(contractAddress)

	contract, err := NewContract(contractAddr, client)
	if err != nil {
		log.Fatal(err)
	}

	// Example function call: newEntry
	txOpts := getTxOpts()
	tx, err := contract.NewEntry(txOpts, "123", "Origin 1", "ToM 1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("New entry transaction hash:", tx.Hash().Hex())

	// Example function call: updateLocation
	txOpts = getTxOpts()
	tx, err = contract.UpdateLocation(txOpts, "123", "Location 1", false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location update transaction hash:", tx.Hash().Hex())

	// Example function call: getProduct
	product, err := contract.GetProduct(nil, "123")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product details:", product)
}

func getTxOpts() *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA("...") // Insert your private key here
	if err != nil {
		log.Fatal(err)
	}

	chainID := big.NewInt(80001) // Chain ID of the Mumbai Testnet
	auth := bind.NewKeyedTransactor(privateKey)
	auth.chainID = chainID
	auth.Nonce = big.NewInt(0)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = big.NewInt(1000000000) // Replace with the appropriate gas price

	return auth
}
