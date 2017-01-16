package lib

import "reflect"
import "fmt"
import "encoding/json"

// Struct2Map function.
func Struct2Map(obj interface{}) (map[string]string, error) {
	data := make(map[string]string)

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Tag.Get("json")
		field := v.Field(i)

		switch field.Kind() {
		case reflect.String:
			data[key] = field.String()
		case reflect.Float32, reflect.Float64:
			data[key] = fmt.Sprintf("%f", field.Float())
		case reflect.Int, reflect.Uint8, reflect.Int16,
			reflect.Int32, reflect.Int64:
			data[key] = fmt.Sprintf("%d", field.Int())
		case reflect.Bool:
			data[key] = fmt.Sprintf("%t", field.Bool())
		case reflect.Map:
			jsonData, err := json.Marshal(field.Interface())
			if err != nil {
				return nil, err
			}
			data[key] = string(jsonData)
		}
	}

	return data, nil
}

func AdanceUnmarshal(data []byte, v interface{}) error {

	if value, ok := v.(StreamResponse); ok {
		value.ExtraInfo = make(map[string]interface{})
	}

	return json.Unmarshal(data, &v)
}
