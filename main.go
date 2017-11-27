package gojsonfieldmap

import "reflect"

// GetJSONObjectFieldMap converts an interface{} into a json mapping string
func GetJSONObjectFieldMap(object interface{}) string {
	o := "{"

	pointerType := reflect.TypeOf(object)
	pointerKind := pointerType.Kind()

	if pointerKind == reflect.Ptr {
		pointerType = pointerType.Elem()
		pointerKind = pointerType.Kind()
	}

	if pointerKind == reflect.Slice {
		pointerType = pointerType.Elem()
		pointerKind = pointerType.Kind()
	}

	if pointerType.Kind() != reflect.Struct {
		return "1"
	}

	for i := 0; i < pointerType.NumField(); i++ {
		field := pointerType.Field(i)

		enableValue := "1"

		fieldType := field.Type.String()

		if fieldType != "string" && fieldType != "int" {
			subObject := reflect.New(field.Type).Interface()
			enableValue = GetJSONObjectFieldMap(subObject)
		}

		ftype := field.Tag.Get("json")

		if ftype == "" {
			continue
		}

		if i != 0 {
			o += ","
		}
		o += "\"" + ftype + "\":" + enableValue
	}
	o += "}"

	return o
}
