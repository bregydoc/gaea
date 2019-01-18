package gaea

import "time"

// LogType implemet a kind of Log
type LogType string

// Mutation is a mutation kind log
var Mutation LogType = "mutation"

// Fixed is a mutation kind log
var Fixed LogType = "fixed"

// Info is a mutation kind log
var Info LogType = "info"

// Error is a mutation kind log
var Error LogType = "error"

// Log is a simple struct to implement how to log accions in the repository
type Log struct {
	At      time.Time         `json:"at" bson:"at"`
	Fields  map[string]string `json:"fields" bson:"fields"`
	Content string            `json:"content" bson:"content"`
	Type    LogType           `json:"type" bson:"type"`
}
