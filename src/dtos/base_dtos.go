package dtos

import (
	"net/http"
	"reflect"
)

type Response struct {
	Code    int         `json:"-"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

func NewResponse(data interface{}, err interface{}, message string) (response Response) {
	code := http.StatusOK

	rt := reflect.TypeOf(data)
	var defaultData interface{}

	if rt != nil {
		switch expression := rt.Kind(); expression {
		case reflect.Struct:
			defaultData = make(map[string]interface{}, 0)
		case reflect.Ptr:
			dataInterface := reflect.New(rt.Elem()).Interface()
			dataKind := reflect.TypeOf(dataInterface).Kind()
			switch dataKind {
			case reflect.Struct:
				defaultData = make(map[string]interface{}, 0)
			default:
				defaultData = make([]interface{}, 0)
			}
		default:
			defaultData = make([]interface{}, 0)
		}
	} else {
		defaultData = make(map[string]interface{}, 0)
	}

	if code >= 300 {
		return Response{
			Code:    code,
			Message: message,
			Data:    defaultData,
			Error:   err,
		}
	} else {
		return Response{
			Code:    code,
			Message: message,
			Data:    data,
			Error:   err,
		}
	}
}
