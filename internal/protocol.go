package internal

import (
	"fmt"
	"strings"
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

// serialize function converts a Command struct into a RESP-compatible array and serializes it
func serialize(c Command) (string, error) {
	var data interface{}
	switch c.action {
	case "GET":
		data = []interface{}{"GET", c.key}
	case "SET":
		data = []interface{}{"SET", c.key, c.value}
	case "DELETE":
		data = []interface{}{"DELETE", c.key}
	default:
		return "", fmt.Errorf("unsupported command: %s", c.action)
	}
	return serializeRESP(data)
}

// serializeRESP handles serialization of various Redis Serialization Protocol (RESP) data types
func serializeRESP(data interface{}) (string, error) {
	switch dataTypeValue := data.(type) {
	case string:
		if strings.HasPrefix(dataTypeValue, "-") {
			return fmt.Sprintf("-%s\r\n", dataTypeValue[1:]), nil // Prints Error
		} else if strings.HasPrefix(dataTypeValue, "+") {
			return fmt.Sprintf("+%s\r\n", dataTypeValue[1:]), nil // PrintsSimple String
		}
		return fmt.Sprintf("$%d\r\n%s\r\n", len(dataTypeValue), dataTypeValue), nil // Prints Bulk String
	case int:
		return fmt.Sprintf(":%d\r\n", dataTypeValue), nil // Prints Integer
	case []interface{}:
		var builder strings.Builder
		builder.WriteString(fmt.Sprintf("*%d\r\n", len(dataTypeValue))) // Prints Array
		for _, elem := range dataTypeValue {
			serializedElem, err := serializeRESP(elem)
			if err != nil {
				return "", err
			}
			builder.WriteString(serializedElem)
		}
		return builder.String(), nil
	default:
		return "", fmt.Errorf("unsupported data type: %T", dataTypeValue)
	}
}