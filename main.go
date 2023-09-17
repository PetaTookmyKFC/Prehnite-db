package main

import (
	"fmt"
	"log"
	"prehnite_light/config"
	"prehnite_light/keystore"
	"prehnite_light/server"
	"sync"
)

func main() {

	fmt.Println("Prehnite Starting ✅")

	fmt.Print("Loading Config ")
	conf, err := config.LoadSystemConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("✅")
	fmt.Printf("Keystore Located at : %s", conf.StoreDirectory)

	fmt.Print("Starting KeyStore  ")

	store := keystore.Start_StringStore(5, conf.StoreDirectory)
	fmt.Print(store)
	// store := keystore.Start_StringStore(5, path.Join())
	// store.Save()

	fmt.Println("✅")
	fmt.Print("Starting Server ")
	var wg sync.WaitGroup
	wg.Add(1)

	// fmt.Println("Installed Routes")

	Inst_Server := server.StartServer(":3080")
	Inst_Server.Static("/", "./Public")
	Inst_Server.PowerKeyStore("/API", store)
	fmt.Println("✅")
	// fmt.Println("Ended Program")

	wg.Wait()

	// fmt.Print("Creating Settings!")

	// config.LoadSystemConfig()

}
