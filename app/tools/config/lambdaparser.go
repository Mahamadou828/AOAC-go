package config

import (
	"fmt"
	"os"
	"reflect"
)

func ParseLambdaCfg(cfg interface{}) error {
	v := reflect.ValueOf(cfg)
	if v.Kind() != reflect.Ptr {
		return ErrInvalidStruct
	}
	fields, err := extractFields(nil, cfg)

	if err != nil {
		return err
	}

	for _, field := range fields {
		if !field.Field.IsValid() || !field.Field.CanSet() {
			return fmt.Errorf("can't set the value of field %s", field.Field.String())
		}

		if len(field.Options.EnvName) == 0 {
			return fmt.Errorf("field %s missing tag env", field.Field.String())
		}

		val := os.Getenv(field.Options.EnvName)

		if len(val) == 0 {
			if field.Options.Required {
				return fmt.Errorf("can't get the value of the field %s", field.Field.String())
			}
			if len(field.Options.DefaultVal) > 0 {
				val = field.Options.DefaultVal
			}
		}

		if err := SetFieldValue(field, val); err != nil {
			return fmt.Errorf("can't set field value for %s: %v", field.Name, err)
		}
	}

	return nil
}
