package keystore

type Yard struct {
	containers map[string]KeyStore
	length     int
}

func InitiliseSiloStore() *Yard {

	yrd := Yard{
		containers: make(map[string]KeyStore),
		length:     0,
	}
	return &yrd
}
