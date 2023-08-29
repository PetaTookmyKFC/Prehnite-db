package keystore

type KeyStore struct {
	ActiveSize int
	CurrSize   int

	Storage   map[string]Data
	FirstItem string
	LastItem  string
}
type Data struct {
	Value string
	Prev  string
}

func Start_StringStore(size int) *KeyStore {
	rKeyStore := KeyStore{
		Storage:    make(map[string]Data),
		ActiveSize: size,
		FirstItem:  "",
		LastItem:   "",
	}
	return &rKeyStore
}

func (keys *KeyStore) getFirst() *Data {
	if keys.FirstItem == "" {
		return nil
	}
	d := keys.Storage[keys.FirstItem]
	return &d
}

func (keys *KeyStore) deleteLast() {
	if keys.LastItem != "" {
		NLI := keys.Storage[keys.LastItem].Prev
		delete(keys.Storage, keys.LastItem)
		keys.LastItem = NLI
		keys.CurrSize = keys.CurrSize - 1
	}
}

func (keys *KeyStore) InsertItem(key string, value string) {
	if keys.ActiveSize > keys.CurrSize {
		keys.deleteLast()
	}

	FI := keys.getFirst()

	if FI != nil {
		FI.Prev = key
	}
	keys.FirstItem = key

	keys.Storage[key] = Data{
		Value: value,
		Prev:  "",
	}
	keys.CurrSize = keys.CurrSize + 1
}

func (keys *KeyStore) GetItem (key string) *string {

  // Check if field exists
  if data, ok := keys.Storage[key]; ok {
    return &data.Value
  }
  // Return nil if not exis
  return nil;

}

