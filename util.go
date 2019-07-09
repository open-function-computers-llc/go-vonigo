package vonigo

import (
	"errors"
)

func createFields(keys, values interface{}) ([]map[string]interface{}, error) {
	var fields []map[string]interface{}

	keySlice, keysIsSlice := keys.([]interface{})
	valSlice, valuesIsSlice := values.([]interface{})

	if keysIsSlice && valuesIsSlice {
		if len(keySlice) != len(valSlice) {
			return fields, errors.New("The number of keys and values must match")
		}

		for k, v := range keySlice {
			item := map[string]interface{}{
				"fieldID":    v,
				"fieldValue": valSlice[k],
			}
			fields = append(fields, item)
		}

		return fields, nil
	}

	keyString, keysIsString := keys.(string)
	valString, valuesIsString := values.(string)

	if keysIsString && valuesIsString {
		item := map[string]interface{}{
			"fieldID":    keyString,
			"fieldValue": valString,
		}
		fields = append(fields, item)
		return fields, nil
	}

	return fields, errors.New("Parameter types did not match")
}
