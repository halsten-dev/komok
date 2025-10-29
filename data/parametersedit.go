package data

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type ParamListOwner interface {
	SetParams(json string)
	GetParams() string
}

// Param represents a single param
type Param struct {
	Label string
	Value interface{}
	Type  reflect.Kind
}

type ParamString struct {
	Label string
	Value string
}

// ParamListData represents the list of single params and hold the struct definition of the params
type ParamListData struct {
	Owner  ParamListOwner
	Struct interface{}
	Params []Param
}

type ParamStringListData []*ParamString

func NewParamListData(owner ParamListOwner, paramStruct interface{}) (*ParamListData, error) {
	pld := &ParamListData{}

	pld.Owner = owner
	pld.Struct = paramStruct

	if !isPointer(reflect.ValueOf(paramStruct)) {
		return nil, fmt.Errorf("paramStruct must be a pointer to a struct")
	}

	paramStr := pld.Owner.GetParams()

	if paramStr != "" {
		bJson := []byte(paramStr)

		if err := json.Unmarshal(bJson, pld.Struct); err != nil {
			return nil, err
		}
	}

	pld.Params = initFromStruct(pld.Struct)

	return pld, nil
}

func (pld *ParamListData) Save() error {
	err := convertToStruct(pld.Params, pld.Struct)

	if err != nil {
		return err
	}

	bJson, err := json.MarshalIndent(pld.Struct, "", "    ")

	if err != nil {
		return err
	}

	strJson := string(bJson)

	pld.Owner.SetParams(strJson)

	return nil
}

// convertToStruct is a private function that helps to convert a []Param to the target struct
func convertToStruct(params []Param, target interface{}) error {
	targetValue := reflect.ValueOf(target)
	if !isPointer(targetValue) {
		return fmt.Errorf("target must be a pointer to a struct")
	}
	targetValue = targetValue.Elem()

	for _, param := range params {
		field := targetValue.FieldByName(param.Label)
		if !field.IsValid() {
			continue // Ignore fields that don't exist in the target struct
		}
		if !field.CanSet() {
			return fmt.Errorf("cannot set field %s", param.Label)
		}

		paramValue := reflect.ValueOf(param.Value)
		if !paramValue.Type().ConvertibleTo(field.Type()) {
			return fmt.Errorf("cannot convert %s to type of field %s", param.Value, param.Label)
		}

		field.Set(paramValue.Convert(field.Type()))
	}

	return nil
}

// initFromStruct is a private function that return a []Param from the given struct
func initFromStruct(param interface{}) []Param {
	var items []Param

	items = make([]Param, 0)

	v := reflect.ValueOf(param)
	t := v.Type()

	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = v.Type()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Name
		fieldType := field.Type.Kind()

		item := Param{
			Label: fieldName,
			Type:  fieldType,
			Value: v.Field(i).Interface(),
		}
		items = append(items, item)
	}

	return items
}

func isPointer(v reflect.Value) bool {
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return false
	}

	return true
}
