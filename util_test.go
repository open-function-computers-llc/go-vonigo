package vonigo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanCreateFieldsFromStrings(t *testing.T) {
	keysString := "12345"
	valuesString := "This is a fake value"
	expectedValue := map[string]interface{}{
		"fieldID":    "12345",
		"fieldValue": "This is a fake value",
	}

	expectedSlice := append([]map[string]interface{}{}, expectedValue)

	fields, _ := createFields(keysString, valuesString)

	assert.Equal(t, expectedSlice, fields)
}

func TestCanCreateFieldsFromSlices(t *testing.T) {
	keySlice := make([]interface{}, 2)
	keySlice[0] = "key numero uno"
	keySlice[1] = "key numero dos"

	valueSlice := make([]interface{}, 2)
	valueSlice[0] = "This is a fake value"
	valueSlice[1] = map[string]interface{}{
		"key": "prop",
	}

	expectedValue1 := map[string]interface{}{
		"fieldID":    keySlice[0],
		"fieldValue": valueSlice[0],
	}

	expectedValue2 := map[string]interface{}{
		"fieldID":    keySlice[1],
		"fieldValue": valueSlice[1],
	}

	expectedSlice := append([]map[string]interface{}{}, expectedValue1, expectedValue2)

	fields, _ := createFields(keySlice, valueSlice)

	assert.Equal(t, expectedSlice, fields)
}
