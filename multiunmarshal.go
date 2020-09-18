package multiunmarshal

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// UnmarshalMultiple takes multiple json documents and unmarshals them into
// the destination interface
func UnmarshalMultiple(j [][]byte, dest interface{}) error {
	// get the value of the target interface
	val := reflect.ValueOf(dest)
	// only accept pointer values
	if val.Kind() != reflect.Ptr {
		return fmt.Errorf("destination must be a pointer")
	}

	elem := val.Elem()
	if !elem.CanAddr() {
		return fmt.Errorf("destination must be addressable (a pointer)")
	}

	// since we're unmarshaling multiple docs, we can only unmarshal
	// into slice types
	if elem.Kind() != reflect.Slice {
		fmt.Errorf("destination must be a slice")
	}

	// figure out base slice type
	// TODO: handle case where the slice base type could be a pointer value
	baseType := elem.Type().Elem()

	for _, b := range j {
		newBase := reflect.New(baseType)
		if err := json.Unmarshal(b, newBase.Interface()); err != nil {
			return err
		}
		// set the value of the target to equal the value
		elem.Set(reflect.Append(elem, newBase.Elem()))
	}
	return nil
}
