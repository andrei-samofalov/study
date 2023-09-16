package models

import (
	"encoding/json"
	"reflect"
)

type Model interface{}

func ToJsonString(m Model) string {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func ToModel(jsonData string, model *Model) (*Model, error) {
	err := json.Unmarshal([]byte(jsonData), model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func setStructFields(s interface{}, data map[string]interface{}) {
	v := reflect.ValueOf(s).Elem()
	for fieldName, fieldValue := range data {
		field := v.FieldByName(fieldName)
		if field.IsValid() {
			if field.Type().AssignableTo(reflect.TypeOf(fieldValue)) {
				field.Set(reflect.ValueOf(fieldValue))
			}
		}
	}
}
