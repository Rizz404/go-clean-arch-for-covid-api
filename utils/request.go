package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func ParseRequestBody(r *http.Request, dst any) error {
	contentType := r.Header.Get("Content-Type")
	contentType = strings.ToLower(strings.TrimSpace(contentType))

	if strings.Contains(contentType, "application/json") {
		return json.NewDecoder(r.Body).Decode(dst)
	} else if strings.Contains(contentType, "application/x-www-form-urlencoded") {
		return parseFormToStruct(r, dst)
	}

	return fmt.Errorf("unsupported content type: %s", contentType)
}

func parseFormToStruct(r *http.Request, dst any) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	v := reflect.ValueOf(dst).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		tag := fieldType.Tag.Get("form")
		if tag == "" {
			tag = fieldType.Tag.Get("json")
		}
		if tag == "" {
			tag = strings.ToLower(fieldType.Name)
		}

		formValue := r.FormValue(tag)
		if formValue == "" {
			continue
		}

		// Set nilai berdasarkan tipe field
		switch field.Kind() {
		case reflect.String:
			field.SetString(formValue)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if intVal, err := strconv.ParseInt(formValue, 10, 64); err == nil {
				field.SetInt(intVal)
			}
		case reflect.Float32, reflect.Float64:
			if floatVal, err := strconv.ParseFloat(formValue, 64); err == nil {
				field.SetFloat(floatVal)
			}
		case reflect.Bool:
			if boolVal, err := strconv.ParseBool(formValue); err == nil {
				field.SetBool(boolVal)
			}
		}
	}

	return nil
}
