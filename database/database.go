package database

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Deepanshu-Sharma-18/api/utils"
)

func StoreJokes() {
	jsonfile, err := os.Open("data/jokes.json")
	defer jsonfile.Close()

	if err != nil {

		panic(err)
	}

	bytes, err := ioutil.ReadAll(jsonfile)

	json.Unmarshal(bytes, &utils.DadJokes)

}
