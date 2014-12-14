package lz4json

import (
	"encoding/json"
	"github.com/bkaradzic/go-lz4"
)

func Marshal(v interface{}) ([]byte, error) {
	buf, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	buf, err = lz4.Encode(nil, buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func Unmarshal(data []byte, v interface{}) error {
	buf, err := lz4.Decode(nil, data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, v)
	if err != nil {
		return err
	}

	return nil
}
