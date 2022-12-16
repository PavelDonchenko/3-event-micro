package common

import (
	"encoding/json"
	"time"
)

type Tweet struct {
	Creator   string    `json:"creator,omitempty" binding:"required"`
	Body      string    `json:"body,omitempty" binding:"required"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func JsonMarshal(data interface{}) []byte {
	_json, err := json.Marshal(data)
	HandleError(err, "Error encoding JSON")
	return _json
}
