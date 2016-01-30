package main

import (
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/btcsuite/btcrpcclient"
	"github.com/btcsuite/btcutil"
	"github.com/d4l3k/btcinterned"
)

var (
	host     = flag.String("host", "localhost:8334", "the btcd host")
	protocol = flag.String("protocol", "ws", "the protocol to talk to btcd on")
	user     = flag.String("user", "", "rpc username")
	pass     = flag.String("pass", "", "rpc password")
)

func main() {
	flag.Parse()

	// Connect to local btcd RPC server
	btcdHomeDir := btcutil.AppDataDir("btcd", false)
	certs, err := ioutil.ReadFile(filepath.Join(btcdHomeDir, "rpc.cert"))
	if err != nil {
		log.Fatal(err)
	}
	connCfg := &btcrpcclient.ConnConfig{
		Host:         *host,
		Endpoint:     *protocol,
		User:         *user,
		Pass:         *pass,
		Certificates: certs,
	}
	client, err := btcrpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	addr, err := btcutil.DecodeAddress("1JZJaDDC44DCKLnezDsbW43Zf8LspCKBYP", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(btcinterned.LookupAddress(client, addr, btcinterned.SixMonths))
}
