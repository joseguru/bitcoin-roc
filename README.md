# bitcoin-rpc
Connect to bitcoind using btcsuite rpcclient

$ go get github.com/btcsuite/btcd/rpcclient

Next, modify the main.go source to specify the correct RPC username and password for the RPC server:

	User: "yourrpcuser",
	Pass: "yourrpcpass",

Finally, navigate to the example's directory and run it with:


    $ go run main.go