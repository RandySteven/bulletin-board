package utils

import (
	"fmt"
	"reflect"
	"time"
)

func UserFullName(firstName, lastName string) string {
	return fmt.Sprintf("%s %s", firstName, lastName)
}

func StringToDate(dateStr string) (time.Time, error) {
	dateTime, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return dateTime, nil
}

func HashImageFile(imageFile string) string {
	return imageFile
}

func GetFieldsOfObject(object interface{}) []string {
	fields := []string{}
	typ := reflect.TypeOf(object).Elem()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i).Name
		fields = append(fields, field)
	}
	return fields
}
