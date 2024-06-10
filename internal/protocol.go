package internal

// / We'll likely want to make our own custom KeyValue interface that we only implement for the types we want to support
// / For now, I'll set these up with empty interfaces for you guys to get started
type Command struct {
	action string
	key    interface{}
	value  interface{}
}

// This should decode RESP bulk string arrays into Commands
func deserialize(s string) (Command, error) {
	return Command{
		action: "",
		key:    "",
		value:  nil,
	}, nil
}

// This should encode the `value` field of the input command as a RESP response
func serialize(c Command) (string, error) {
	return `-Error: Not implemented yet\r\n`, nil
}
