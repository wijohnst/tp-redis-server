package internal

import (
	"errors"
	"reflect"
	"strconv"
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
func serialize(c Command) (string, error) {
	// short-circuit
	if c.key == nil {
		return "", errors.New("key is not defined")
	}

	elements := []string{getCommandSerialization(c)}
	elements = append(elements, getFieldSerialization(c.action))
	elements = append(elements, getFieldSerialization(c.key))

	if c.value != nil {
		elements = append(elements, getFieldSerialization(c.value))
	}

	out := strings.Join(elements, "")
	return out, nil
}

func appendCRLF(s string) string {
	CRLF := "\r\n"

	return s + CRLF
}

func getCommandSerialization(c Command) string {
	out := "*"
	nonNilFields := 1

	if c.key != nil {
		nonNilFields++
	}

	if c.value != nil {
		nonNilFields++
	}

	return appendCRLF(out + strconv.Itoa(nonNilFields))
}

func getFieldSerialization(f interface{}) string {
	switch f := f.(type) {

	case string:
		chars := strconv.Itoa(len(f))
		meta := appendCRLF("$" + chars)

		return meta + appendCRLF(f)

	case int:
		s := strconv.Itoa(f)
		if f < 0 {
			return appendCRLF(":" + s)
		}

		return appendCRLF(":+" + s)

	case bool:
		out := "#f"

		if f {
			out = "#t"
		}

		return appendCRLF(out)

	default:
		field := reflect.ValueOf(f)

		if field.Kind() == reflect.Slice {
			out := appendCRLF("*" + strconv.Itoa(field.Len()))

			for i := 0; i < field.Len(); i++ {
				out += getFieldSerialization(field.Index(i).Interface())
			}

			return out
		}

		return ""
	}

}
