package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int `json:"n"`
	Balance int `json:"b"`
}

func main() {
	account := Account{1, 100}
	str, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	println(string(str))

	// With encoder
	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(account)
	if err != nil {
		panic(err)
	}

	accountStr := []byte(`{"n":1,"b":500}`)
	var accountX Account
	err = json.Unmarshal(accountStr, &accountX)
	if err != nil {
		panic(err)
	}
	println(accountX.Balance)
}
