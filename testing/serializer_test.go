package internal

import (
	"fmt"
	"testing"
	"wijohnst/tp-redis-server/internal"
)

/*
	//INTEGER VALUE
	c1 := internal.Command{
		Action: "gET",
		Key:    "key1",
		Value:  5,
	}
	c2 := internal.Command{
		Action: "sET",
		Key:    "key1",
		Value:  104,
	}
	//SLICE VALUE
	c3 := internal.Command{
		Action: "gEt",
		Key:    "key2",
		Value:  []int{1, 2, 3},
	}
	c4 := internal.Command{
		Action: "deLetE",
		Key:    "key2",
		Value:  []string{"hello", "world"},
	}
*/

func TestIsValidAction(t *testing.T) {
	//Case 1: Valid - GET
	command := internal.Command{
		Action: "GET",
	}
	expected := true
	actual := internal.IsValidAction(command)
	if actual != expected {
		t.Errorf("Case 1/4: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 1/4: PASS")
	}

	//Case 2: Valid - SET
	command = internal.Command{
		Action: "set",
		Key:    "key",
	}
	expected = true
	actual = internal.IsValidAction(command)
	if actual != expected {
		t.Errorf("Case 2/4: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 2/4: PASS")
	}

	//Case 3: Valid - DELETE
	command = internal.Command{
		Action: "dELeTe",
		Key:    "key",
		Value:  "value",
	}
	expected = true
	actual = internal.IsValidAction(command)
	if actual != expected {
		t.Errorf("Case 3/4: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 3/4: PASS")
	}

	//Case 4: Invalid
	command = internal.Command{
		Action: "invalid",
	}
	expected = false
	actual = internal.IsValidAction(command)
	if actual != expected {
		t.Errorf("Case 4/4: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 4/4: PASS")
	}
	fmt.Println("----------------------------------------")
}

func TestIsValidCommand(t *testing.T) {
	//Case 1: Valid - GET
	command := internal.Command{
		Action: "GET",
		Key:    "key",
	}
	expected := true
	actual := internal.IsValidCommand(command)
	if actual != expected {
		t.Errorf("Case 1/6: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 1/6: PASS")
	}

	//Case 2: Valid - SET
	command = internal.Command{
		Action: "set",
		Key:    "key",
	}
	expected = true
	actual = internal.IsValidCommand(command)
	if actual != expected {
		t.Errorf("Case 2/6: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 2/6: PASS")
	}

	//Case 3: Valid - DELETE
	command = internal.Command{
		Action: "dELeTe",
		Key:    "key",
		Value:  "value",
	}
	expected = true
	actual = internal.IsValidCommand(command)
	if actual != expected {
		t.Errorf("Case 3/6: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 3/6: PASS")
	}

	//Case 4: Invalid
	command = internal.Command{
		Action: "GET",
	}
	expected = false
	actual = internal.IsValidCommand(command)
	if actual != expected {
		t.Errorf("Case 4/6: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 4/6: PASS")
	}

	//Case 5: Invalid
	command = internal.Command{
		Action: "SET",
		Key:    "",
		Value:  "testvalue",
	}
	expected = false
	actual = internal.IsValidCommand(command)
	if actual != expected {
		t.Errorf("Case 5/6: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 5/6: PASS")
	}

	//Case 6: Invalid
	command = internal.Command{
		Action: "DELETE",
		Value:  "testvalue",
	}
	expected = false
	actual = internal.IsValidCommand(command)
	if actual != expected {
		t.Errorf("Case 6/6: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 6/6: PASS")
	}
	fmt.Println("----------------------------------------")
}

func TestGetReturnArraySize(t *testing.T) {
	//Case 1: Action
	command := internal.Command{
		Action: "get",
	}
	expected := "2"
	actual := internal.GetReturnArraySize(command)
	if actual != expected {
		t.Errorf("Case 1/4: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 1/4: PASS")
	}

	//Case 2: Action, Key
	command = internal.Command{
		Action: "dELeTe",
		Key:    "key",
	}
	expected = "2"
	actual = internal.GetReturnArraySize(command)
	if actual != expected {
		t.Errorf("Case 2/4: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 2/4: PASS")
	}

	//Case 3: Action, Key, Value
	command = internal.Command{
		Action: "SEt",
		Key:    "key",
		Value:  "value",
	}
	expected = "3"
	actual = internal.GetReturnArraySize(command)
	if actual != expected {
		t.Errorf("Case 3/4: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 3/4: PASS")
	}

	//Case 4: Action, Value
	command = internal.Command{
		Action: "testCommand",
		Key:    "",
		Value:  "value",
	}
	expected = "Invalid input Command"
	actual = internal.GetReturnArraySize(command)
	if actual != expected {
		t.Errorf("Case 4/4: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 4/4: PASS")
	}
	fmt.Println("----------------------------------------")
}

func TestActionIsSet(t *testing.T) {
	//Case 1: Get, False
	command := internal.Command{
		Action: "get",
	}
	expected := false
	actual := internal.ActionIsSet(command)
	if actual != expected {
		t.Errorf("Case 1/4: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 1/4: PASS")
	}

	//Case 2: Delete, False
	command = internal.Command{
		Action: "dELeTe",
		Key:    "key",
	}
	expected = false
	actual = internal.ActionIsSet(command)
	if actual != expected {
		t.Errorf("Case 2/4: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 2/4: PASS")
	}

	//Case 3: Set, True
	command = internal.Command{
		Action: "SEt",
		Key:    "key",
		Value:  "value",
	}
	expected = true
	actual = internal.ActionIsSet(command)
	if actual != expected {
		t.Errorf("Case 3/4: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 3/4: PASS")
	}

	//Case 4: Invalid, False
	command = internal.Command{
		Action: "testCommand",
		Key:    "",
		Value:  "value",
	}
	expected = false
	actual = internal.ActionIsSet(command)
	if actual != expected {
		t.Errorf("Case 4/4: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 4/4: PASS")
	}
	fmt.Println("----------------------------------------")
}

func TestSerializeAction(t *testing.T) {
	//Case 1: All lowercase
	command := internal.Command{
		Action: "get",
	}
	expected := "$3\\r\\nGET\\r\\n"
	actual := internal.SerializeAction(command)
	if actual != expected {
		t.Errorf("Case 1/3: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 1/3: PASS")
	}

	//Case 2: Upper/lowercase mix
	command = internal.Command{
		Action: "dELeTe",
		Key:    "key",
	}
	expected = "$6\\r\\nDELETE\\r\\n"
	actual = internal.SerializeAction(command)
	if actual != expected {
		t.Errorf("Case 2/3: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 2/3: PASS")
	}

	//Case 3: All uppercase
	command = internal.Command{
		Action: "TESTCOMMAND",
		Key:    "key",
		Value:  "value",
	}
	expected = "$11\\r\\nTESTCOMMAND\\r\\n"
	actual = internal.SerializeAction(command)
	if actual != expected {
		t.Errorf("Case 3/3: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 3/3: PASS")
	}
	fmt.Println("----------------------------------------")
}

func TestSerializeKey(t *testing.T) {
	//Case 1: Empty key
	command := internal.Command{
		Action: "get",
	}
	expected := "$0\\r\\n\\r\\n"
	actual := internal.SerializeKey(command)
	if actual != expected {
		t.Errorf("Case 1/3: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 1/3: PASS")
	}

	//Case 2: Short key
	command = internal.Command{
		Action: "get",
		Key:    "key",
	}
	expected = "$3\\r\\nkey\\r\\n"
	actual = internal.SerializeKey(command)
	if actual != expected {
		t.Errorf("Case 2/3: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 2/3: PASS")
	}

	//Case 3: Long key
	command = internal.Command{
		Action: "dELeTe",
		Key:    "TESTLONGKEYVALUE",
	}
	expected = "$16\\r\\nTESTLONGKEYVALUE\\r\\n"
	actual = internal.SerializeKey(command)
	if actual != expected {
		t.Errorf("Case 3/3: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 3/3: PASS")
	}
	fmt.Println("----------------------------------------")
}

func TestSerializeValue(t *testing.T) {
	//Case 1: Empty
	command := internal.Command{
		Action: "set",
		Key:    "key1",
	}
	expected := "$0\\r\\n\\r\\n"
	actual := internal.SerializeValue(command)
	if actual != expected {
		t.Errorf("Case 1/4: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 1/4: PASS")
	}

	//Case 2: No value
	command = internal.Command{
		Action: "set",
		Key:    "key2",
		Value:  "",
	}
	expected = "$0\\r\\n\\r\\n"
	actual = internal.SerializeValue(command)
	if actual != expected {
		t.Errorf("Case 2/4: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 2/4: PASS")
	}

	//Case 3: Short value
	command = internal.Command{
		Action: "invalidCommand",
		Key:    "key",
		Value:  "value",
	}
	expected = "$5\\r\\nvalue\\r\\n"
	actual = internal.SerializeValue(command)
	if actual != expected {
		t.Errorf("Case 3/4: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 3/4: PASS")
	}

	//Case 4: Long value
	command = internal.Command{
		Action: "invalidCommand",
		Key:    "key",
		Value:  "TestLongValue",
	}
	expected = "$13\\r\\nTestLongValue\\r\\n"
	actual = internal.SerializeValue(command)
	if actual != expected {
		t.Errorf("Case 4/4: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 4/4: PASS")
	}
	fmt.Println("----------------------------------------")
}

func TestSerializer(t *testing.T) {
	//Case 1: Empty command
	command := internal.Command{}
	expected := ""
	actual, err := internal.Serialize(command)
	if err != nil {
		t.Error(err)
	} else if actual != expected {
		t.Errorf("Case 1/14: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 1/14: PASS")
	}

	//Case 2: Get - Action
	command = internal.Command{
		Action: "get",
	}
	expected = ""
	actual, err = internal.Serialize(command)
	if err != nil {
		t.Error(err)
	} else if actual != expected {
		t.Errorf("Case 2/14: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 2/14: PASS")
	}

	//Case 3: Get - Action, Key
	command = internal.Command{
		Action: "get",
		Key:    "key1",
	}
	expected = "*2\\r\\n$3\\r\\nGET\\r\\n$4\\r\\nkey1\\r\\n"
	actual, err = internal.Serialize(command)
	if err != nil {
		t.Error(err)
	} else if actual != expected {
		t.Errorf("Case 3/14: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 3/14: PASS")
	}

	//Case 4: Get - Action, Key, Value
	command = internal.Command{
		Action: "get",
		Key:    "key2",
		Value:  "test_value",
	}
	expected = "*2\\r\\n$3\\r\\nGET\\r\\n$4\\r\\nkey2\\r\\n"
	actual, err = internal.Serialize(command)
	if err != nil {
		t.Error(err)
	} else if actual != expected {
		t.Errorf("Case 4/14: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 4/14: PASS")
	}

	//Case 5: Get - Action, Value
	command = internal.Command{
		Action: "get",
		Value:  "test_value",
	}
	expected = ""
	actual, err = internal.Serialize(command)
	if err != nil {
		t.Error(err)
	} else if actual != expected {
		t.Errorf("Case 5/14: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 5/14: PASS")
	}

	//Case 6: Set - Action
	command = internal.Command{
		Action: "sET",
		Key:    "",
		Value:  "",
	}
	expected = ""
	actual, err = internal.Serialize(command)
	if err != nil {
		t.Error(err)
	} else if actual != expected {
		t.Errorf("Case 6/14: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 6/14: PASS")
	}

	//Case 7: Set - Action, Key
	command = internal.Command{
		Action: "sEt",
		Key:    "key1",
	}
	expected = "*3\\r\\n$3\\r\\nSET\\r\\n$4\\r\\nkey1\\r\\n$0\\r\\n\\r\\n"
	actual, err = internal.Serialize(command)
	if err != nil {
		t.Error(err)
	} else if actual != expected {
		t.Errorf("Case 7/14: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 7/14: PASS")
	}

	//Case 8: Set - Action, Key, Value
	command = internal.Command{
		Action: "SEt",
		Key:    "key2",
		Value:  "test_value",
	}
	expected = "*3\\r\\n$3\\r\\nSET\\r\\n$4\\r\\nkey2\\r\\n$10\\r\\ntest_value\\r\\n"
	actual, err = internal.Serialize(command)
	if err != nil {
		t.Error(err)
	} else if actual != expected {
		t.Errorf("Case 8/14: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 8/14: PASS")
	}

	//Case 9: Set - Action, Value
	command = internal.Command{
		Action: "SeT",
		Value:  "test_value",
	}
	expected = ""
	actual, err = internal.Serialize(command)
	if err != nil {
		t.Error(err)
	} else if actual != expected {
		t.Errorf("Case 9/14: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 9/14: PASS")
	}

	//Case 10: Delete - Action
	command = internal.Command{
		Action: "DELETE",
		Key:    "",
		Value:  "",
	}
	expected = ""
	actual, err = internal.Serialize(command)
	if err != nil {
		t.Error(err)
	} else if actual != expected {
		t.Errorf("Case 10/14: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 10/14: PASS")
	}

	//Case 11: Delete - Action, Key
	command = internal.Command{
		Action: "DELETE",
		Key:    "key1",
	}
	expected = "*2\\r\\n$6\\r\\nDELETE\\r\\n$4\\r\\nkey1\\r\\n"
	actual, err = internal.Serialize(command)
	if err != nil {
		t.Error(err)
	} else if actual != expected {
		t.Errorf("Case 11/14: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 11/14: PASS")
	}

	//Case 12: Delete - Action, Key, Value
	command = internal.Command{
		Action: "deLeTe",
		Key:    "key2",
		Value:  "test_value",
	}
	expected = "*2\\r\\n$6\\r\\nDELETE\\r\\n$4\\r\\nkey2\\r\\n"
	actual, err = internal.Serialize(command)
	if err != nil {
		t.Error(err)
	} else if actual != expected {
		t.Errorf("Case 12/14: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 12/14: PASS")
	}

	//Case 13: Delete - Action, Value
	command = internal.Command{
		Action: "delete",
		Value:  "test_value",
	}
	expected = ""
	actual, err = internal.Serialize(command)
	if err != nil {
		t.Error(err)
	} else if actual != expected {
		t.Errorf("Case 13/14: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 13/14: PASS")
	}

	//Case 14: Set - Action, Key, Value
	command = internal.Command{
		Action: "SEt",
		Key:    "test_key",
		Value:  "5",
	}
	expected = "*3\\r\\n$3\\r\\nSET\\r\\n$8\\r\\ntest_key\\r\\n$1\\r\\n5\\r\\n"
	actual, err = internal.Serialize(command)
	if err != nil {
		t.Error(err)
	} else if actual != expected {
		t.Errorf("Case 14/14: FAIL\texpected: %v, actual: %v", expected, actual)
	} else {
		fmt.Println("Case 14/14: PASS")
	}
	fmt.Println("----------------------------------------")
}
