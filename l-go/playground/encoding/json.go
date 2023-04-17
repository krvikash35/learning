package encoding

import (
	"encoding/json"
	"log"
)

type change struct {
	Field    string `json:"field"`
	OldValue any    `json:"old_value"`
	NewValue any    `json:"new_value"`
}

type changeType struct {
	Field    string
	OldValue string
	NewValue string
}

type changes = []change

var changesJson = `[
	{"field": "description.en", "old_value": "a", "new_value": "b"},
	{"field": "active", "old_value": true, "new_value": false}
	]`

func Run() {
	var changes changes
	err := json.Unmarshal([]byte(changesJson), &changes)
	if err != nil {
		panic(err)
	}
	log.Printf("unmarshal %+v", changes)
}
