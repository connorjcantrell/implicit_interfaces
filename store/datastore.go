package store

type MyDataStore struct {
	userData map[string]string
}

func (sds MyDataStore) UserNameByID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

// Factory function
func New() MyDataStore {
	return MyDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
		},
	}
}
