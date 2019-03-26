package vonigo

import (
	"testing"
)

func TestCanGetBaseParameter(t *testing.T) {

	// Passing valid action for purpose of test
	resultString, _ := getBaseParams("update")

	if resultString["securityToken"] != securityToken {
		t.Error("The secuity token was not added to the base parameters")
	}

	// Testing for getting/retrieving items from API
	resultString, _ = getBaseParams("retrieve")

	if resultString["method"] != "1" {
		t.Error("The method action parameter was not correctly set in the base parameter")
	}

	// Testing for editing/updating item in API
	resultString, _ = getBaseParams("update")

	if resultString["method"] != "2" {
		t.Error("The method action parameter was not correctly set in the base parameter")
	}

	// Testing for adding/creating items from API
	resultString, _ = getBaseParams("create")

	if resultString["method"] != "3" {
		t.Error("The method action parameter was not correctly set in the base parameter")
	}

	_, err := getBaseParams("randomstring")

	if err == nil {
		t.Error("getBaseParams should have thrown an error when an invalid string is passed")
	}

}
