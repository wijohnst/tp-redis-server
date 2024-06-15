package internal

import (
	"fmt"
)

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
    switch c.action {
    case "GET":
		if c.key == nil {
            return "", fmt.Errorf("key cannot be nil for GET command")
        }
        if key, ok := c.key.(string); ok {
            return fmt.Sprintf("*2\r\n$3\r\nGET\r\n$%d\r\n%s\r\n", len(key), key), nil
        }
        return "", fmt.Errorf("invalid key type for GET command")
    case "SET":
		if c.key == nil {
            return "", fmt.Errorf("key cannot be nil for SET command")
        }
		if c.value == nil {
            return "", fmt.Errorf("value cannot be nil for SET command")
        }
        if key, ok := c.key.(string); ok {
            if value, ok := c.value.(string); ok {
                return fmt.Sprintf("*3\r\n$3\r\nSET\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n", len(key), key, len(value), value), nil
            }
            return "", fmt.Errorf("invalid value type for SET command")
        }
        return "", fmt.Errorf("invalid key type for SET command")
    case "DELETE":
		if c.key == nil {
            return "", fmt.Errorf("key cannot be nil for DELETE command")
        }
        if key, ok := c.key.(string); ok {
            return fmt.Sprintf("*2\r\n$6\r\nDELETE\r\n$%d\r\n%s\r\n", len(key), key), nil
        }
        return "", fmt.Errorf("invalid key type for DELETE command")
    default:
        return "", fmt.Errorf("unsupported command: %s", c.action)
    }
}
