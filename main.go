package main

import (
	"fmt"
	"log"
	"prehnite_light/config"
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

	fmt.Println("Starting KeyStore")

	// store := keystore.Start_StringStore(5, path.Join())
	// store.Save()
	// var wg sync.WaitGroup

	// fmt.Println("Created KeyStore!")
	// wg.Add(1)

	// fmt.Println("Installed Routes")

	// Inst_Server := server.StartServer(":3080")

	// Inst_Server.Static("/", "./Public")

	// fmt.Println("Ended Program")

	// wg.Wait()

	// fmt.Print("Creating Settings!")

	// config.LoadSystemConfig()

}
