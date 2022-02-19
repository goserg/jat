package colorify

import "encoding/json"

type Colorifier interface {
	Colorify(json.RawMessage) (json.RawMessage, error)
}
