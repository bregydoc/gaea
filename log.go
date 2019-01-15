package gaea

import "time"

type LogType string

var Mutation LogType = "mutation"
var Fixed LogType = "fixed"
var Info LogType = "info"
var Error LogType = "error"

type Log struct {
	At      time.Time         `json:"at" bson:"at"`
	Fields  map[string]string `json:"fields" bson:"fields"`
	Content string            `json:"content" bson:"content"`
	Type    LogType           `json:"type" bson:"type"`
}
