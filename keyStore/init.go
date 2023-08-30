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
	Next  string
}

func Start_StringStore(size int) *KeyStore {
	rKeyStore := KeyStore{
		Storage:    make(map[string]Data),
		ActiveSize: size,
		FirstItem:  "",
		LastItem:   "",
	}
	rKeyStore.LoadPrev()
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
func (keys *KeyStore) deleteItem(key string) bool {
	if val, ok := keys.Storage[key]; ok {
		var linKey *Data
		if val.Next != "" {
			linKey = keys.GetItem(val.Next)
			if linKey != nil {
				linKey.Prev = val.Prev
			}
		}
		if val.Prev != "" {
			linKey = keys.GetItem(val.Prev)
			if linKey != nil {
				linKey.Next = val.Next
			}
		}

		delete(keys.Storage, key)
		keys.CurrSize = keys.CurrSize - 1
		return true
	}
	return false
}
func (keys *KeyStore) InsertItem(key string, value string) {
	if keys.ActiveSize > keys.CurrSize {
		keys.deleteLast()
	}

	FI := keys.getFirst()

	if FI != nil {
		FI.Prev = key
	}
	nwPair := Data{
		Value: value,
		Prev:  "",
		Next:  keys.FirstItem,
	}

	keys.FirstItem = key

	keys.Storage[key] = nwPair

	keys.CurrSize = keys.CurrSize + 1
	keys.Save()
}

func (keys *KeyStore) GetItem(key string) *Data {
	// Check if field exists
	if data, ok := keys.Storage[key]; ok {
		return &data
	}
	// Return nil if not exis
	return nil
}
