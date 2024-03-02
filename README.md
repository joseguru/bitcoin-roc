# bitcoin-rpc
Connect to bitcoind using btcsuite rpcclient

    $ go get github.com/btcsuite/btcd/rpcclient
    $ go get github.com/joho/godotenv

Next create your .env file by
    cp .env.example .env

Next, modify the .env source to specify the correct RPC username and password for the RPC server:

	User: "yourrpcuser",
	Pass: "yourrpcpass",

Finally, navigate to the example's directory and run it with:


    $ go run main.go