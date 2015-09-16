package json

import (
	"encoding/json"
	"net/http"
)

func DecodeRequest(r *http.Request, v interface{}) {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&v); err != nil {
		panic(err)
	}
}

func EncodeResponse(w http.ResponseWriter, item interface{}) {
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(item); err != nil {
		panic(err)
	}
}
