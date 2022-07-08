package config

import (
	"fmt"
	"reflect"
)

//ParseAdminCfg is a custom parser for handling configuration for the admin package.
func ParseAdminCfg(cfg any) error {
	t := reflect.TypeOf(cfg)
	if t.Kind() != reflect.Ptr {
		return fmt.Errorf("passing non-pointer value to ParseAdminCfg")
	}

	fs, err := extractFields(nil, cfg)
	if err != nil {
		return fmt.Errorf("can't extract fields from cfg: %v", err)
	}

	for _, f := range fs {
		if !f.Field.IsValid() || !f.Field.CanSet() {
			return fmt.Errorf("can't set the value of field %s", f.Field.String())
		}

		if len(f.Options.Help) == 0 {
			return fmt.Errorf("no help was provided for field %s", f.Field.String())
		}

		var val string
		fmt.Printf("%s: ", f.Options.Help)
		if _, err := fmt.Scan(&val); err != nil {
			return fmt.Errorf("invalid value for field %s", err)
		}

		if err := SetFieldValue(f, val); err != nil {
			return fmt.Errorf("can't set value for field %s", err)
		}
	}

	return nil
}
