package config

import (
	"reflect"
)

type ClaimMap map[string]interface{}

func GetClaimMap[T any](claims T, claimMap ClaimMap) ClaimMap {
	if claimMap == nil {
		claimMap = ClaimMap{}
	}

	t := reflect.TypeOf(claims)
	v := reflect.ValueOf(claims)
	length := t.NumField()
	for index := 0; index < length; index++ {
		typeField := t.Field(index)
		valueField := v.Field(index)
		tag := typeField.Tag
		if isNonZero(valueField) {
			value := valueField.Interface()
			claim := tag.Get("claim")
			if claim != "" {
				claimMap[claim] = value
			}
		}
	}

	return claimMap
}

func GetConfigMap[T any](claims T) T {
	t := reflect.TypeOf(claims)
	v := reflect.ValueOf(&claims).Elem()
	length := t.NumField()
	for index := 0; index < length; index++ {
		typeField := t.Field(index)
		valueField := v.Field(index)
		tag := typeField.Tag
		if isNonZero(valueField) {
			value := valueField.Interface()
			claim := tag.Get("config")
			if claim == "" {
				continue
			}
			switch val := value.(type) {
			case string:
				valueField.SetString(GetEnvString(claim, val))
			case int:
				valueField.SetInt(int64(GetEnvInt(claim, &val)))
			}
		}
	}
	return claims
}

func isNonZero(value reflect.Value) bool {
	if value.IsZero() {
		return false
	}

	switch value.Kind() {
	case reflect.Chan:
		fallthrough
	case reflect.Func:
		fallthrough
	case reflect.Interface:
		fallthrough
	case reflect.Map:
		fallthrough
	case reflect.Pointer:
		fallthrough
	case reflect.Slice:
		return !value.IsNil()
	default:
		return true
	}
}
