package colorjson

import (
	"encoding/json"

	"github.com/TylerBrock/colorjson"
)

type CJ struct {
}

func (cj *CJ) Colorify(raw json.RawMessage) (json.RawMessage, error) {
	var obj map[string]interface{}
	if err := json.Unmarshal(raw, &obj); err != nil {
		return nil, err
	}

	f := colorjson.NewFormatter()
	f.Indent = 4

	s, err := f.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return s, nil
}
