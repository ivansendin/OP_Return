

// ./btcd -u=teste -P=teste --notls
// change chainsvrresults.go... bug?

package main

import (
	"fmt"
	"strings"
	rpcclient "github.com/btcsuite/btcd/rpcclient"
	"log"
	"github.com/btcsuite/btcd/txscript"
	"encoding/hex"
)

func main() {
     first := 600000
     last := first+1
     connCfg := &rpcclient.ConnConfig{
     	     	Host:         "localhost:8334",
		User:         "teste",
		Pass:         "teste",
		HTTPPostMode: true, 
		DisableTLS:   true, 
		}
	
     client, err := rpcclient.New(connCfg, nil)
     if err != nil {
	log.Fatal(err)
     }
     defer client.Shutdown()

     for n:=first; n <=last;n++ {
     	 blockHash, _ := client.GetBlockHash(int64(n))
	 fmt.Printf("Block Hash: %s\n",blockHash)
	 txs, err := client.GetBlockVerboseTx(blockHash)
	 if err != nil {
	    log.Fatal(err)
	 }
	    
	 for _,_tx := range(txs.Tx) {
	     script, err := hex.DecodeString(_tx.Hex)
     	     if err != nil {	
     	     	log.Fatal(err)
     	     }
	     ds,_ := txscript.DisasmString(script)
	     if (strings.Contains(ds,"OP_RETURN")) {
	     	fmt.Printf("%s\n",_tx.Hash)
	     }
	}
     }
}
