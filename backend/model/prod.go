package model

import (
	"encoding/json"
	"io"

	"github.com/Jim-Lin/bee-bee-alert/backend/util"
)

type Prod struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Url   string `json:"url"`
}

func DecodeJson(r io.Reader) Prod {
	var prod Prod

	decoder := json.NewDecoder(r)
	err := decoder.Decode(&prod)
	util.CheckError(err)

	return prod
}
