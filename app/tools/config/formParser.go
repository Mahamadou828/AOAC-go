package config

import (
	"fmt"
	"github.com/Mahamadou828/AOAC/foundation/lambda"
	"github.com/aws/aws-lambda-go/events"
	"io"
	"reflect"
)

// ParseForm parse all field of a form data
func ParseForm(form any, r events.APIGatewayProxyRequest) error {
	t := reflect.TypeOf(form)
	if t.Kind() != reflect.Ptr {
		return fmt.Errorf("passing non-pointer type to parser: %v", form)
	}

	fs, err := extractFields(nil, form)
	if err != nil {
		return fmt.Errorf("can't extract fields: %v", err)
	}

	mapFormValue := map[string]string{}
	reader, err := lambda.NewReaderMultipart(r)
	if err != nil {
		return fmt.Errorf("can't create reader: %v from request", err)
	}

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read part %s: %v", part.FormName(), err)
		}
		defer part.Close()
		content, err := io.ReadAll(part)
		if err != nil {
			return fmt.Errorf("can't read part %s: %v", part.FormName(), err)
		}
		mapFormValue[part.FormName()] = string(content)
	}

	for _, f := range fs {
		if !f.Field.IsValid() || !f.Field.CanSet() {
			return fmt.Errorf("can't set field: %v, %v", f.Name, err)
		}

		value, ok := mapFormValue[f.Name]
		if !ok {
			continue
		}

		if err := SetFieldValue(f, value); err != nil {
			return fmt.Errorf("can't set field: %v, %v", f.Name, err)
		}
	}

	return nil
}
