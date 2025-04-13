package service

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	selection "github.com/GoAdminGroup/go-admin/template/types/form/select"
	template2 "html/template"
	"reflect"
	"strconv"
)

// RequestData 定义请求数据结构
type RequestData struct {
	Value string `json:"value"`
}

// ResponseData 定义响应数据结构
type ResponseData struct {
	Message string `json:"message"`
}

var myDb db.Connection

func SetDb(appDb db.Connection) {
	myDb = appDb
}
func GetDb() db.Connection {
	return myDb
}

func TransFieldOptions(result interface{}, text, value string) types.FieldOptions {
	rV := reflect.ValueOf(result)
	rT := reflect.TypeOf(result)
	var myOptions types.FieldOptions
	if rT.Kind() == reflect.Slice {
		for i := 0; i < rV.Len(); i++ {
			rVItem := rV.Index(i)
			switch rVItem.Kind() {
			case reflect.Struct:
				textValue := rVItem.FieldByName(text)
				valueValue := rVItem.FieldByName(value)
				myOptions = append(myOptions, types.FieldOption{
					Text:  TransStr(textValue),
					Value: TransStr(valueValue),
				})
			case reflect.Map:
				textValue := rVItem.MapIndex(reflect.ValueOf(text))
				valueValue := rVItem.MapIndex(reflect.ValueOf(value))
				myOptions = append(myOptions, types.FieldOption{
					Text:  TransStr(textValue),
					Value: TransStr(valueValue),
				})
			default:
				return myOptions
			}

		}
	}
	return myOptions
}

func TransSelectionOptions(result interface{}, text, value string) selection.Options {
	rV := reflect.ValueOf(result)
	rT := reflect.TypeOf(result)
	var myOptions selection.Options
	if rT.Kind() == reflect.Slice {
		for i := 0; i < rV.Len(); i++ {
			rVItem := rV.Index(i)
			switch rVItem.Kind() {
			case reflect.Struct:
				textValue := rVItem.FieldByName(text)
				valueValue := rVItem.FieldByName(value)
				myOptions = append(myOptions, selection.Option{
					Text: TransStr(textValue),
					ID:   TransStr(valueValue),
				})
			case reflect.Map:
				textValue := rVItem.MapIndex(reflect.ValueOf(text))
				valueValue := rVItem.MapIndex(reflect.ValueOf(value))
				myOptions = append(myOptions, selection.Option{
					Text: TransStr(textValue),
					ID:   TransStr(valueValue),
				})
			default:
				return myOptions
			}

		}
	}
	return myOptions
}

func TransStr(t reflect.Value) string {
	switch t.Kind() {
	case reflect.Int64, reflect.Int:
		return strconv.FormatInt(t.Int(), 10)
	case reflect.String:
		return t.String()
	case reflect.Interface:
		return TransStr(reflect.ValueOf(t.Interface()))
	default:
		return ""
	}
}

func Int64ToTmp(v int64) template2.HTML {
	return template.HTML(strconv.FormatInt(v, 10))
}
