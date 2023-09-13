package main

import (
	"fmt"
	"sync"

	"prehnite_light/keystore"
	"prehnite_light/server"
)

func main() {
	fmt.Println("Started Program")

	store := keystore.Start_StringStore(5)
	store.Save()
	var wg sync.WaitGroup

	fmt.Println("Created KeyStore!")
	wg.Add(1)

	fmt.Println("Installed Routes")

	Inst_Server := server.StartServer(":3080")

	Inst_Server.Static("/", "./Public")

	fmt.Println("Ended Program")

	wg.Wait()
}
