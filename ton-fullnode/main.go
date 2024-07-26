package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
)

const (
	mainnetConfigPath = "https://ton.org/global.config.json"
	testnetConfigPath = "https://ton-blockchain.github.io/testnet-global.config.json"
)

func main() {
	client := liteclient.NewConnectionPool()

	err := client.AddConnectionsFromConfigUrl(context.Background(), mainnetConfigPath)
	if err != nil {
		log.Fatalln("Failed to add connections from config:", err)
		return
	}

	api := ton.NewAPIClient(client)
	master, err := api.GetMasterchainInfo(context.Background())
	if err != nil {
		log.Fatalln("Failed to get masterchain info:", err)
		return
	}

	message, err := json.Marshal(master)
	if err != nil {
		log.Fatalln("Failed to marshal masterchain info:", err)
		return
	}

	log.Println(string(message))
}
