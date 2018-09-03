package mmq

import "encoding/json"

type MsgID [16]byte

type Msg struct {
	ID   MsgID
	Body json.RawMessage
}
