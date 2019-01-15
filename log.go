package gaea

import "time"

type LogType string

var Mutation LogType = "mutation"
var Fixed LogType = "fixed"
var Info LogType = "info"
var Error LogType = "error"

type Log struct {
	At      time.Time         `json:"at"`
	Fields  map[string]string `json:"fields"`
	Content string            `json:"content"`
	Type    LogType           `json:"type"`
}
