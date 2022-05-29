package api

import (
	// "errors"
	// "fmt"
	// "reflect"
)
const(
	queue = "Informer"
)

type Header struct {
	Queue string
}

func GetServiceName() string {
	return queue
}

func GetParams(map[string]interface{}) {

}

// func SetField(obj interface{}, name string, value interface{}) error {
//     structValue := reflect.ValueOf(obj).Elem()
//     structFieldValue := structValue.FieldByName(name)

//     if !structFieldValue.IsValid() {
//         return fmt.Errorf("No such field: %s in obj", name)
//     }

//     if !structFieldValue.CanSet() {
//         return fmt.Errorf("Cannot set %s field value", name)
//     }

//     structFieldType := structFieldValue.Type()
//     val := reflect.ValueOf(value)
//     if structFieldType != val.Type() {
//         return errors.New("Provided value type didn't match obj field type")
//     }

//     structFieldValue.Set(val)
//     return nil
// }

// func (h *Header) FillStruct(m map[string]interface{}) error {
//     for k, v := range m {
//         err := SetField(h, k, v)
//         if err != nil {
//             return err
//         }
//     }
//     return nil
// }