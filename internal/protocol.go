package internal

import (
	"fmt"
	"strconv"
	"strings"
)

// / We'll likely want to make our own custom KeyValue interface that we only implement for the types we want to support
// / For now, I'll set these up with empty interfaces for you guys to get started
type Command struct {
	Action string
	Key    string //@@interface{}
	Value  string //@@interface{}
}

const CRLF string = "\\r\\n"

// This should decode RESP bulk string arrays into Commands
func deserialize(s string) (Command, error) {
	return Command{
		Action: "",
		Key:    "",
		Value:  "", //@@nil,
	}, nil
}

// Determine if the given Command action is valid
func IsValidAction(c Command) bool {
	action := strings.ToUpper(c.Action)
	return action == "GET" || action == "SET" || action == "DELETE"
}

// Determine if the given Command is valid
func IsValidCommand(c Command) bool {
	/* Valid Command:
	- Action is GET, SET, or DELETE
	- Key exists
	*/
	if !IsValidAction(c) {
		return false
	} else if c.Key == "" {
		return false
	} else {
		return true
	}
}

// Determine size of serialized array
func GetReturnArraySize(c Command) string {
	if strings.ToUpper(c.Action) == "GET" || strings.ToUpper(c.Action) == "DELETE" { //can ignore Command value even if it exists
		return "2"
	} else if strings.ToUpper(c.Action) == "SET" {
		return "3"
	} else {
		return "Invalid input Command"
	}
}

// Checks if the given Command action is SET
func ActionIsSet(c Command) bool {
	return strings.ToUpper(c.Action) == "SET"
}

// Serialize Command action
func SerializeAction(c Command) string {
	//ASSUMPTION: Action will always be a bulk string
	var actionLength = strconv.Itoa(len(c.Action))
	return "$" + actionLength + CRLF + strings.ToUpper(c.Action) + CRLF
}

// Serialize Command key
func SerializeKey(c Command) string {
	//ASSUMPTION: Key will always be a bulk string
	var keyLength = strconv.Itoa(len(c.Key))
	return "$" + keyLength + CRLF + c.Key + CRLF
}

// Serialize Command value
func SerializeValue(c Command) string {
	//ASSUMPTION: Value will always be a bulk string
	var valueLength = strconv.Itoa(len(c.Value))
	return "$" + valueLength + CRLF + c.Value + CRLF
}

// This should encode the `value` field of the input command as a RESP response
func Serialize(c Command) (result string, errMessage error) {
	//Check validity of input command
	if !IsValidCommand(c) {
		fmt.Println("Invalid command")
		return "", nil
	}

	//ASSUMPTION: Input Command will always be an array
	result = "*" + GetReturnArraySize(c) + CRLF

	result = result + SerializeAction(c)
	result = result + SerializeKey(c)
	if ActionIsSet(c) { //value only needs to be serialized if the Command action is SET
		result = result + SerializeValue(c)
	}
	return result, nil
}
