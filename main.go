package main

import (
	"fmt"
	"sync"

	keystore "prehnite_light/keyStore"
)

func main() {
	fmt.Println("Started Program")

	store := keystore.Start_StringStore(5)
	store.Save()
	var wg sync.WaitGroup

	fmt.Println("Created KeyStore!")
	wg.Add(1)

	fmt.Println("Installed Routes")

	fmt.Println("Ended Program")

	wg.Wait()
}
