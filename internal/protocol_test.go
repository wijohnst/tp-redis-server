package internal

import "testing"

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

/*
************************************************************************************
							SERIALIZE FUNCTION TESTS
************************************************************************************
*/
func TestSerializeGET(t *testing.T) {
	expected := "*2\r\n$3\r\nGET\r\n$3\r\nfoo\r\n"
	command := Command{
		action: "GET",
		key:    "foo",
		value:  nil,
	}
	actual, err := serialize(command)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestSerializeSET(t *testing.T) {
	expected := "*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"
	command := Command{
		action: "SET",
		key:    "foo",
		value:  "bar",
	}
	actual, err := serialize(command)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestSerializeDELETE(t *testing.T) {
	expected := "*2\r\n$6\r\nDELETE\r\n$3\r\nfoo\r\n"
	command := Command{
		action: "DELETE",
		key:    "foo",
		value:  nil,
	}
	actual, err := serialize(command)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

/*
************************************************************************************
							SERIALIZE RESP FUNCTION TESTS
************************************************************************************
*/

func TestSerializeSimpleStringRESP(t *testing.T) {
	expected := "+OK\r\n"
	actual, err := serializeRESP("+OK")
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestSerializeErrorRESP(t *testing.T) {
	expected := "-Error message\r\n"
	actual, err := serializeRESP("-Error message")
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestSerializeIntegerRESP(t *testing.T) {
	expected := ":123\r\n"
	actual, err := serializeRESP(123)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestSerializeBulkStringRESP(t *testing.T) {
	expected := "$11\r\nhello world\r\n"
	actual, err := serializeRESP("hello world")
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}

func TestSerializeArrayRESP(t *testing.T) {
	expected := "*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"
	actual, err := serializeRESP([]interface{}{"foo", "bar"})
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("expected: %s, actual: %s", expected, actual)
	}
}