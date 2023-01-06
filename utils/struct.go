package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"reflect"
	"strings"
)

func IsStruct(who interface{}) bool {
	//reflect.Indirect(reflect.ValueOf(who)).Kind() // actually i haven't try indirect with non-pointer struct yet and don't know what's gonna happen when i use it.
	return (reflect.ValueOf(who).Kind() == reflect.Ptr && reflect.ValueOf(who).Elem().Kind() == reflect.Struct) || (reflect.ValueOf(who).Kind() == reflect.Struct)
}

func StructValueByKey(out interface{}, notation string) interface{} {
	if notation != "" && IsStruct(out) {

		var notations []string
		notations = strings.Split(notation, ".")

		if len(notations) == 0 && notation != "" {
			notations = append(notations, notation)
		}

		var r reflect.Value
		var newNotation string

		if reflect.TypeOf(out).Kind() == reflect.Ptr {
			r = reflect.ValueOf(out).Elem()
		} else if reflect.TypeOf(reflect.ValueOf(out)).Kind() == reflect.Struct {
			r = reflect.ValueOf(out)
		}

		if strings.Contains(notation, ".") {
			newNotation = strings.ReplaceAll(notation, notations[0]+".", "")
		}

		value := r.FieldByName(cases.Title(language.English, cases.NoLower).String(notations[0]))
		if !value.IsValid() {
			return nil
		}

		return StructValueByKey(value.Interface(), newNotation)
	}

	return out
}
