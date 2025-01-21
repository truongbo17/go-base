package utils

import (
	"encoding/json"
	"reflect"
	"strings"
)

func JSONHelper(input interface{}) (interface{}, error) {
	val := reflect.ValueOf(input)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, nil
	}

	output := make(map[string]interface{})
	typeOfT := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typeOfT.Field(i).Tag.Get("json")
		if fieldName == "" || fieldName == "-" {
			fieldName = typeOfT.Field(i).Name
		}

		if field.IsZero() {
			output[fieldName] = nil
		} else {
			output[fieldName] = field.Interface()
		}
	}
	return output, nil
}

func JSONResponse(data interface{}) (map[string]interface{}, error) {
	normalizedData, err := JSONHelper(data)
	if err != nil {
		return nil, err
	}
	bytes, err := json.Marshal(normalizedData)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func FormatStackTrace(stack []byte) string {
	lines := strings.Split(string(stack), "\n")
	return strings.Join(lines, "\n")
}
