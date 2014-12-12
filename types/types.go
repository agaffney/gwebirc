package types

type IrcEvent struct {
	Id          uint64   `json:"id"`
	Source      string   `json:"source"`
	Source_nick string   `json:"source_nick"`
	Code        string   `json:"code"`
	Args        []string `json:"-"`
	Msg         string   `json:"msg"`
	Connection  string   `json:"connection"`
	Channel     string   `json:"channel"`
	Raw         string   `json:"-"`
}
