package convert

import "reflect"

func StructToMap(item interface{}) map[string]interface{} {
	values := make(map[string]interface{})

	itemValue := reflect.ValueOf(item)
	itemType := reflect.TypeOf(item)

	for i := 0; i < itemValue.NumField(); i++ {
		field := itemType.Field(i)
		value := itemValue.Field(i).Interface()

		// Check if the field is a pointer
		if field.Type.Kind() == reflect.Ptr {
			// If the pointer is nil, set the value to nil
			if itemValue.Field(i).IsNil() {
				value = nil
			} else {
				// If the pointer is not nil, get the value it points to
				value = itemValue.Field(i).Elem().Interface()
			}
		}

		values[field.Name] = value
	}

	return values
}
