package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetJson(url string, res interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(body, res)

	return nil
}
