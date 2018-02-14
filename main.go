package gojsonfieldmap

import (
	"reflect"
)

// GetJSONObjectFieldMap converts an interface{} into a json mapping string
func GetJSONObjectFieldMap(object interface{}) string {
	if object == nil {
		return "{}"
	}

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

			if field.Anonymous {
				//Embedded classes
				stringLen := len(enableValue)
				if stringLen > 3 {
					//Strip the {}'s off since this is considered part of this class now
					o += enableValue[1 : len(enableValue)-1]
				}
				//Anonymous is handled as the current field so just continue
				continue
			}
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
