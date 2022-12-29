package msgbus

var (
	ErrEventIdWasSet     = errors.New("msgbuss: event id has set already")
	ErrEventTsWasSet     = errors.New("msgbuss: event timestamps has set already")
	ErrEventBucketWasSet = errors.New("msgbuss: event bucket has set already")
	ErrEventDataWasSet   = errors.New("msgbuss: event data has set already")
)
