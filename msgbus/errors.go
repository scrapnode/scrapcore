package msgbus

import "errors"

var (
	ErrEventIdWasSet   = errors.New("msgbuss: event id has set already")
	ErrEventDataWasSet = errors.New("msgbuss: event data has set already")
	ErrEventDataEmpty  = errors.New("msgbuss: event data is empty")
)
