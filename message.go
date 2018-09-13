package mmq

import "encoding/json"

type MsgID [16]byte

type Msg struct {
	ID   MsgID           `json:"id"`
	Body json.RawMessage `json:"body"`
	TS   int64           `json:"ts"`
}
