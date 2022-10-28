package lambda

import (
	"encoding/json"
	"strings"
)

func DecodeBody(body string, val any) error {
	decoder := json.NewDecoder(strings.NewReader(body))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(val); err != nil {
		return err
	}

	return nil
}
