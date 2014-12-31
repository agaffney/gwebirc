package types

import (
	"bytes"
	"fmt"
)

type Event struct {
	Id         uint64 `json:"id"`
	Type       string `json:"type"`
	Connection string `json:"connection"`
	Attribs    map[string]string
	Args       []string `json:"-"`
}

func (e *Event) Init() {
	e.Attribs = make(map[string]string, 0)
}

func (e *Event) MarshalJSON() ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf(`{"id":"%d","type":"%s","connection":"%s"`, e.Id, e.Type, e.Connection))
	for key, value := range e.Attribs {
		if key[0] != '_' {
			buffer.WriteString(fmt.Sprintf(`,"%s":"%s"`, key, value))
		}
	}
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}
