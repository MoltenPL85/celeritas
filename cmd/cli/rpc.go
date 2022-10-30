package main

import (
	"fmt"
	"net/rpc"
	"os"

	"github.com/fatih/color"
)

func rpcClient(inMaintenance bool) {
	rpcPort := os.Getenv("RPC_PORT")
	c, err := rpc.Dial("tcp", "127.0.0.1:"+rpcPort)
	if err != nil {
		exitGracefully(err)
	}

	fmt.Println("Connected...")
	var result string
	err = c.Call("RPCServer.MaintenanceMode", inMaintenance, &result)
	if err != nil {
		exitGracefully(err)
	}

	color.Yellow(result)
}
