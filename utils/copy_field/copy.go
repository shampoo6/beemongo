package copy_field

import (
	"github.com/globalsign/mgo/bson"
	"reflect"
	"strings"
)

func Copy(strPtr interface{}, distPtr interface{}) {
	_srcV := reflect.ValueOf(strPtr)
	_distV := reflect.ValueOf(distPtr)
	srcT := reflect.TypeOf(strPtr)
	distT := reflect.TypeOf(distPtr)
	if srcT.Kind() != reflect.Ptr || distT.Kind() != reflect.Ptr ||
		srcT.Elem().Kind() == reflect.Ptr || distT.Elem().Kind() == reflect.Ptr {
		panic("Fatal error:type of parameters must be Ptr of value")
	}
	if _srcV.IsNil() || _distV.IsNil() {
		panic("Fatal error:value of parameters should not be nil")
	}
	srcV := _srcV.Elem()
	dstV := _distV.Elem()
	srcFields := deepFields(reflect.ValueOf(strPtr).Elem().Type())
	for _, v := range srcFields {
		if v.Anonymous {
			continue
		}
		dst := dstV.FieldByName(v.Name)
		src := srcV.FieldByName(v.Name)
		if v.Name == "Id" {
			if src.Type().Name() == "string" && dst.Type().Name() == "ObjectId" {
				id := src.Interface().(string)
				id = strings.Trim(id, " ")
				if id == "" {
					continue
				}
				dst.Set(reflect.ValueOf(bson.ObjectIdHex(id)))
			} else if src.Type().Name() == "ObjectId" && dst.Type().Name() == "string" {
				id := src.Interface().(bson.ObjectId)
				if id == bson.ObjectIdHex("") {
					continue
				}
				dst.SetString(id.Hex())
			}
		}
		if !dst.IsValid() {
			continue
		}
		if src.Type() == dst.Type() && dst.CanSet() {
			dst.Set(src)
			continue
		}
		if src.Kind() == reflect.Ptr && !src.IsNil() && src.Type().Elem() == dst.Type() {
			dst.Set(src.Elem())
			continue
		}
		if dst.Kind() == reflect.Ptr && dst.Type().Elem() == src.Type() {
			dst.Set(reflect.New(src.Type()))
			dst.Elem().Set(src)
			continue
		}
	}
	return
}

func deepFields(iFaceType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField

	for i := 0; i < iFaceType.NumField(); i++ {
		v := iFaceType.Field(i)
		if v.Anonymous && v.Type.Kind() == reflect.Struct {
			fields = append(fields, deepFields(v.Type)...)
		} else {
			fields = append(fields, v)
		}
	}

	return fields
}
