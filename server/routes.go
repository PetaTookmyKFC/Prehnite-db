package server

import (
	"net/http"
	"prehnite_light/keystore"
)

func (serv *Server) PowerKeyStore(WebRoute string, keys *keystore.KeyStore) {
	// GET
	// serv.Get("/get", ())
	// SET
	// DELETE
	// Make KeyStore
	serv.Get("/API", InsertHandler(keys))
}

func InsertHandler(keys *keystore.KeyStore) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		// var data struct {
		// 	Name  string
		// 	Value string
		// }

		// err := json.NewDecoder(req.Body).Decode(&data)
		// if err != nil {
		// 	res.Write([]byte("err.Error() FUCKING ERROR MAN!"))
		// 	res.WriteHeader(500)
		// 	return
		// }
		res.Write([]byte("Got it :SMILEYFACE:"))
		res.Header().Set("content-type", "application/TEXT")
		res.WriteHeader(200)
	}
}

// func InsertHandler(res http.ResponseWriter, req *http.Request) {
// 	res.WriteHeader(200)
// 	req.ParseForm()

// 	fmt.Println(req.Form)

// 	res.Write([]byte("Hello Worlds"))

// 	// COPIED CODE¬!!!!>?

// 	decoder := json.NewDecoder(req.Body)
//     err := decoder.Decode(&t)
//     if err != nil {
//         panic(err)
//     }
//     log.Println(t.Test)

// }
