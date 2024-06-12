package internal

import (
	"strconv"
	"testing"
)

func TestDeserializeGet(t *testing.T) {
	// Case 0: basic get
	expected := Command{
		action: "GET",
		key:    "test",
		value:  nil,
	}
	actual, err := deserialize(`*2\r\n$3\r\nGET\r\n$4\r\ntest\r\n`)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("Test GET case 0: expected: %v, actual: %v", expected, actual)
	}

	// Case 1: anything passed in after the key in a Get command is ignored
	actual, err = deserialize(`*3\r\n$3\r\nGET\r\n$4\r\ntest\r\n$5\r\nvalue\r\n`)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("Test GET case 1: expected: %v, actual: %v", expected, actual)
	}

	// Case 2: missing key should throw error
	_, err = deserialize(`*1\r\n$3\r\nget\r\n`)
	if err == nil {
		t.Error("Expected GET with no key to throw error")
	}
}

func TestDeserializeSet(t *testing.T) {
	// Case 0: String Value
	expected := Command{
		action: "SET",
		key:    "test",
		value:  "test_value",
	}
	actual, err := deserialize(`*3\r\n$3\r\nSET\r\n$4\r\ntest\r\n$10\r\ntest_value\r\n`)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("Test SET case 0: expected: %v, actual: %v", expected, actual)
	}

	// Case 1: Integer Value
	expected = Command{
		action: "SET",
		key:    "key",
		value:  5,
	}
	actual, err = deserialize(`*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n:5\r\n`)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("Test SET case 1: expected: %v, actual: %v", expected, actual)
	}

	// Case 2: Array value
	expected = Command{
		action: "SET",
		key:    "key",
		value:  []int{1, 2, 3},
	}
	actual, err = deserialize(`*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n*3\r\n:1\r\n:2\r\n:3\r\n`)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("Test SET case 1: expected: %v, actual: %v", expected, actual)
	}
}

func TestDeserializeDelete(t *testing.T) {
	// Case 0: Basic delete
	expected := Command{
		action: "DELETE",
		key:    "test",
		value:  nil,
	}
	actual, err := deserialize(`*2\r\n$3\r\nDELETE\r\n$4\r\ntest\r\n`)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("Test DELETE case 0: expected: %v, actual: %v", expected, actual)
	}

	// Case 1: input values should be ignored
	expected = Command{
		action: "DELETE",
		key:    "test",
		value:  nil,
	}
	actual, err = deserialize(`*2\r\n$3\r\nDELETE\r\n$4\r\ntest\r\n$10\r\ntest_value\r\n`)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("Test DELETE case 1: expected: %v, actual: %v", expected, actual)
	}

	// Case 2: missing key should throw an error
	_, err = deserialize(`*1\r\n$3\r\nDELETE\r\n`)
	if err == nil {
		t.Error("Test DELETE case 2: Expected DELETE with no key to throw error")
	}
}

var preString = "*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n"

func TestSerializeStrings(t *testing.T) {
	sut := "serialize"
	desc := "handle strings"

	// Case 0: key and value are bulk strings
	expected := preString + "$3\r\nbar\r\n"
	actual, err := serialize(Command{action: "SET", key: "foo", value: "bar"})

	if err != nil {
		handleErr(sut, desc, 0, err, t)
	}

	if actual != expected {
		handleAssertionError(sut, desc, 0, t, expected, actual)
	}

	// Case 1: key is bulk string and value is nil
	expected = "*2\r\n$3\r\nSET\r\n$3\r\nfoo\r\n"
	actual, err = serialize(Command{action: "SET", key: "foo"})

	if err != nil {
		handleErr(sut, desc, 1, err, t)
	}

	if actual != expected {
		handleAssertionError(sut, desc, 1, t, expected, actual)
	}

	// Case 2: missing key should throw an error
	_, err = serialize(Command{action: "SET", key: nil, value: nil})

	if err == nil {
		t.Error(sut + " - " + desc + " > case 2 : Expected Command where c.key == nil to throw")
	}
}

func TestSerializeIntegers(t *testing.T) {
	sut := "serialize"
	desc := "handle integers"

	// Case 0: key is a positive integer
	expected := "*2\r\n$3\r\nSET\r\n:+1\r\n"
	actual, err := serialize(Command{action: "SET", key: 1})

	if err != nil {
		handleErr(sut, desc, 0, err, t)
	}

	if actual != expected {
		handleAssertionError(sut, desc, 0, t, expected, actual)
	}

	// Case 1: key is a negative integer
	expected = "*2\r\n$3\r\nSET\r\n:-1\r\n"
	actual, err = serialize(Command{action: "SET", key: -1})

	if err != nil {
		handleErr(sut, desc, 1, err, t)
	}

	if actual != expected {
		handleAssertionError(sut, desc, 1, t, expected, actual)
	}

	// Case 2: value is a positive integer
	expected = preString + ":+1\r\n"
	actual, err = serialize(Command{action: "SET", key: "foo", value: 1})

	if err != nil {
		handleErr(sut, desc, 2, err, t)
	}

	if actual != expected {
		handleAssertionError(sut, desc, 2, t, expected, actual)
	}
}

func TestSerializeArrays(t *testing.T) {
	sut := "serialize"
	desc := "handle arrays"

	// case 0 : handle primitive array
	stringArr := []interface{}{"bar", "bax"}
	command := Command{action: "SET", key: "foo", value: stringArr}
	actual, err := serialize(command)
	expected := preString + "*2\r\n$3\r\nbar\r\n$3\r\nbax\r\n"

	if err != nil {
		handleErr(sut, desc, 0, err, t)
	}

	if actual != expected {
		handleAssertionError(sut, desc, 0, t, expected, actual)
	}

	// case 1 : handle nested array
	arrayArr := [][]string{{"bar"}, {"bax", "biz"}}
	command = Command{action: "SET", key: "foo", value: arrayArr}
	actual, err = serialize(command)
	expected = preString + "*2\r\n*1\r\n$3\r\nbar\r\n*2\r\n$3\r\nbax\r\n$3\r\nbiz\r\n"

	if err != nil {
		handleErr(sut, desc, 1, err, t)
	}

	if actual != expected {
		handleAssertionError(sut, desc, 0, t, expected, actual)
	}
}

func TestSerializeBooleans(t *testing.T) {
	sut := "serialize"
	desc := "handle booleans"

	// case 0 : handle true
	command := Command{action: "SET", key: "foo", value: true}
	actual, err := serialize(command)
	expected := preString + "#t\r\n"

	if err != nil {
		handleErr(sut, desc, 0, err, t)
	}

	if actual != expected {
		handleAssertionError(sut, desc, 0, t, expected, actual)
	}

	// case 1 : handle false
	command = Command{action: "SET", key: "foo", value: false}
	actual, err = serialize(command)
	expected = preString + "#f\r\n"

	if err != nil {
		handleErr(sut, desc, 1, err, t)
	}

	if actual != expected {
		handleAssertionError(sut, desc, 1, t, expected, actual)
	}
}

func handleErr(sut, desc string, c int, e error, t *testing.T) {
	t.Errorf(sut+" - "+desc+" > case "+strconv.Itoa(c)+": encountered an exception: %q", e)
}

func handleAssertionError(sut, desc string, c int, t *testing.T, expected, actual interface{}) {
	t.Errorf(sut+" - "+desc+" > case "+strconv.Itoa(c)+":\nEXPECTED:\n%q\nRECEIVED:\n%q\n", expected, actual)
}

func TestAppendCRLF(t *testing.T) {
	sut := "appendCRLF"
	desc := "basic use"

	// Case 0: should return the correctly appended string
	expected := "foo\r\n"
	actual := appendCRLF("foo")

	if actual != expected {
		handleAssertionError(sut, desc, 0, t, expected, actual)
	}

}

func TestGetCommandSerialization(t *testing.T) {
	sut := "getCommandSerialization"
	desc := "basic use"

	// Case 0: should return the correct serialization for the target command
	command := Command{action: "SET", key: "foo", value: "bar"}
	expected := "*3\r\n"
	actual := getCommandSerialization(command)

	if actual != expected {
		handleAssertionError(sut, desc, 0, t, expected, actual)
	}

	// Case 1: should return the correct serialization for the target command
	command = Command{action: "SET", key: "foo"}
	expected = "*2\r\n"
	actual = getCommandSerialization(command)

	if actual != expected {
		handleAssertionError(sut, desc, 1, t, expected, actual)
	}

	// Case 2: should return the correct serialization for the target command
	command = Command{action: "SET"}
	expected = "*1\r\n"
	actual = getCommandSerialization(command)

	if actual != expected {
		handleAssertionError(sut, desc, 2, t, expected, actual)
	}
}

func TestGetFieldSerialization(t *testing.T) {
	sut := "getFieldSerialization"
	desc := "handle string"

	// Case 0: string should return "$" + input length
	inputStr := "foo"
	expected := "$3\r\nfoo\r\n"
	actual := getFieldSerialization(inputStr)

	if actual != expected {
		handleAssertionError(sut, desc, 0, t, expected, actual)
	}

	// Case 1: positive integer should return ":+"
	desc = "handle positive integer"
	inputInt := 1
	expected = ":+1\r\n"
	actual = getFieldSerialization(inputInt)

	if actual != expected {
		handleAssertionError(sut, desc, 1, t, expected, actual)
	}

	// Case 2: negative integer should return ":-"
	desc = "handle negative integer"
	inputInt = -1
	expected = ":-1\r\n"
	actual = getFieldSerialization(inputInt)

	if actual != expected {
		handleAssertionError(sut, desc, 2, t, expected, actual)
	}

}
