package meerkat

import "net/http"
import "encoding/json"
import "encoding/xml"
import "errors"
import "strings"
import "reflect"
import "strconv"

type BinderFunc func(*Context, interface{})error


func   JSONBind(context *Context, obj interface{}) error {
	req := context.request;
	if req == nil || req.Body == nil {
		errors.New("req is null.")
	}
	return json.NewDecoder(req.Body).Decode(obj);
}

func   XMLBind(context *Context, obj interface{}) error {
	req := context.request;
	decoder := xml.NewDecoder(req.Body)
	return decoder.Decode(obj)
}

func   FormBind(context *Context, obj interface{}) error {
	param,err := context.FormParams();
	if err!=nil{
		return err
	}
	return bindData(obj,param,"form")
}

func   QueryBind(context *Context, obj interface{}) error {
	return bindData(obj,context.QueryParams(),"query")
}


func DefaultBinder(req *http.Request) BinderFunc {
	if req.Method == http.MethodGet  || req.Method == http.MethodDelete{
		return QueryBind
	}
	contentType := req.Header.Get("Content-Type")
	switch  {
	case strings.HasPrefix(contentType,ContentTypeJSON):
		return JSONBind
	case strings.HasPrefix(contentType,ContentTypeXML), strings.HasPrefix(contentType,ContentTypeXML2):
		return XMLBind
	case strings.HasPrefix(contentType,ContentTypePOSTForm),strings.HasPrefix(contentType,ContentTypeMultipartPOSTForm):
		return FormBind
	default:
		return nil
	}
}

func   bindData(obj interface{} ,data map[string][]string, tag string) error {
	values := reflect.ValueOf(obj).Elem()
	types := reflect.TypeOf(obj).Elem()
	if types.Kind() != reflect.Struct {
		return errors.New("binding element must be a struct")
	}

	for i := 0; i < types.NumField(); i++ {
		typeField := types.Field(i)
		valueField := values.Field(i)
		if !valueField.CanSet() {
			continue
		}
		name := typeField.Tag.Get(tag)
		inputValue, exists := data[name]
		if exists{
			err := setbindValue(valueField,inputValue[0])
			if err !=nil {
				return err
			}
		}
	}
	return nil
}

func setbindValue( value reflect.Value,valueStr string) error{
	valueKind := value.Kind()
	switch valueKind {
	case reflect.Int:
		return setIntField(valueStr, 0, value)
	case reflect.Int8:
		return setIntField(valueStr, 8, value)
	case reflect.Int16:
		return setIntField(valueStr, 16, value)
	case reflect.Int32:
		return setIntField(valueStr, 32, value)
	case reflect.Int64:
		return setIntField(valueStr, 64, value)
	case reflect.Uint:
		return setUintField(valueStr, 0, value)
	case reflect.Uint8:
		return setUintField(valueStr, 8, value)
	case reflect.Uint16:
		return setUintField(valueStr, 16, value)
	case reflect.Uint32:
		return setUintField(valueStr, 32, value)
	case reflect.Uint64:
		return setUintField(valueStr, 64, value)
	case reflect.Bool:
		return setBoolField(valueStr, value)
	case reflect.Float32:
		return setFloatField(valueStr, 32, value)
	case reflect.Float64:
		return setFloatField(valueStr, 64, value)
	case reflect.String:
		value.SetString(valueStr)
	default:
		return errors.New("unknown type")
	}
	return nil
}


func setIntField(value string, bitSize int, field reflect.Value) error {
	if value == "" {
		value = "0"
	}
	intVal, err := strconv.ParseInt(value, 10, bitSize)
	if err == nil {
		field.SetInt(intVal)
	}
	return err
}

func setUintField(value string, bitSize int, field reflect.Value) error {
	if value == "" {
		value = "0"
	}
	uintVal, err := strconv.ParseUint(value, 10, bitSize)
	if err == nil {
		field.SetUint(uintVal)
	}
	return err
}

func setBoolField(value string, field reflect.Value) error {
	if value == "" {
		value = "false"
	}
	boolVal, err := strconv.ParseBool(value)
	if err == nil {
		field.SetBool(boolVal)
	}
	return err
}

func setFloatField(value string, bitSize int, field reflect.Value) error {
	if value == "" {
		value = "0.0"
	}
	floatVal, err := strconv.ParseFloat(value, bitSize)
	if err == nil {
		field.SetFloat(floatVal)
	}
	return err
}
