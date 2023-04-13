package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	FirstName string `required:"true"`
	LastName  string `required:"false"`
	Age       uint8  `required:"true"`
}

func main() {
	person := Person{
		FirstName: "Tiago",
		Age:       32,
	}

	err := validateFields(person)

	fmt.Println(err)
}

func validateFields(stc any) error {
	obj := reflect.TypeOf(stc)
	value := reflect.ValueOf(stc)

	for i := 0; i < obj.NumField(); i++ {
		f := obj.Field(i)
		v := value.Field(i)

		required := f.Tag.Get("required")
		if required == "false" {
			continue
		}

		switch v.Kind() {
		case reflect.String:
			if v.String() == "" {
				return fmt.Errorf("Field %s is required", f.Name)
			}
		case reflect.Uint8:
			if v.Uint() == 0 {
				return fmt.Errorf("Field %s is required", f.Name)
			}
		}
	}

	return nil
}
