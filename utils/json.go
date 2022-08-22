package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Method int

const (
	POST Method = iota
	GET
	DELETE
)

func (e Method) String() string {
	switch e {
	case POST:
		return "POST"
	case GET:
		return "GET"
	case DELETE:
		return "DELETE"
	default:
		return "XX"
	}
}

func GetJson(url string, method Method, res interface{}, jsonBody string) error {

	var r *http.Response
	var err error

	clienteHttp := &http.Client{}
	req, err := http.NewRequest(method.String(), url, bytes.NewBuffer([]byte(jsonBody)))

	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	r, err = clienteHttp.Do(req)

	if err != nil {
		return err
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(body, res)
	if r.StatusCode > 299 {
		return errors.New(r.Status)
	}

	return nil
}
