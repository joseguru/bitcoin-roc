package main

import (
	"log"

	"github.com/btcsuite/btcd/rpcclient"
)

func main(){
connConf:= &rpcclient.ConnConfig{

	Host:         "localhost:38332",
		User:         "joseguru",
		Pass:         "_Uzt0IvkJFWHYXH_KoX5IygcJVrnGrp1N6YenFBgHfo",
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true,
}
// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpcclient.New(connConf, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()
	info,err :=client.GetBlockChainInfo()
	if err != nil {
		log.Fatal(err)
	}

	networkInfo, err:= client.GetNetworkInfo()
	if err != nil {
		log.Fatal(err)
	}
	// Get the current block count.
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	

    
	log.Printf("Block count: %d", blockCount)
	log.Println("Block chain info:", info.Chain)
	log.Printf("Block version info: %d", networkInfo.Version)

}