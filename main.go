package main

import (
	"fmt"
	keystore "prehnite_light/keyStore"
)

func main() {
	fmt.Println("Started Program")

	store := keystore.Start_StringStore(5)

	store.InsertItem("TEST", "Example String")
	fmt.Println("Inserted Item")

	val := store.GetItem("TEST")
	fmt.Printf("Got item %s ", *val)

	store.Save()
	fmt.Println("Saved Items")

	fmt.Println("Ended Program")
}
