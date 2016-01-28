package models

type JsonTree struct {
	Id       int        `json:"id"`
	Text     string     `json:"text"`
	IconCls  string     `json:"iconCls"`
	State    string     `json:"state"`
	Checked  bool       `json:"checked"`
	Children []JsonTree `json:"children"`
}
