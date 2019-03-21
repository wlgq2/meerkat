package meerkat

import "net/http"
import "encoding/json"
import "encoding/xml"
import "errors"
import "strings"
import "reflect"
import "strconv"

type Binder interface {
	GetBindHanler(*http.Request) BinderFunc
	
} 
type BinderFunc func(*Context, interface{}) error
type DefaultBinder struct{

}

func  (binder *DefaultBinder) JSONBind(context *Context, obj interface{}) error {
	req := context.request;
	if req == nil || req.Body == nil {
		errors.New("req is null.")
	}
	return json.NewDecoder(req.Body).Decode(obj);
}

func   (*DefaultBinder) XMLBind(context *Context, obj interface{}) error {
	req := context.request;
	decoder := xml.NewDecoder(req.Body)
	return decoder.Decode(obj)
}

func   (binder *DefaultBinder) FormBind(context *Context, obj interface{}) error {
	param,err := context.FormParams();
	if err!=nil{
		return err
	}
	return binder.bindData(obj,param,"form")
}

func   (binder *DefaultBinder) QueryBind(context *Context, obj interface{}) error {
	return binder.bindData(obj,context.QueryParams(),"query")
}


func (binder *DefaultBinder) GetBindHanler(req *http.Request) BinderFunc {
	if req.Method == http.MethodGet  || req.Method == http.MethodDelete{
		return binder.QueryBind
	}
	contentType := req.Header.Get("Content-Type")
	switch  {
	case strings.HasPrefix(contentType,ContentTypeJSON):
		return binder.JSONBind
	case strings.HasPrefix(contentType,ContentTypeXML), strings.HasPrefix(contentType,ContentTypeXML2):
		return binder.XMLBind
	case strings.HasPrefix(contentType,ContentTypePOSTForm),strings.HasPrefix(contentType,ContentTypeMultipartPOSTForm):
		return binder.FormBind
	default:
		return nil
	}
}

func   (binder *DefaultBinder) bindData(obj interface{} ,data map[string][]string, tag string) error {
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
			err := binder.setbindValue(valueField,inputValue[0])
			if err !=nil {
				return err
			}
		}
	}
	return nil
}

func (binder *DefaultBinder) setbindValue( value reflect.Value,valueStr string) error{
	valueKind := value.Kind()
	switch valueKind {
	case reflect.Int:
		return binder.setIntField(valueStr, 0, value)
	case reflect.Int8:
		return binder.setIntField(valueStr, 8, value)
	case reflect.Int16:
		return binder.setIntField(valueStr, 16, value)
	case reflect.Int32:
		return binder.setIntField(valueStr, 32, value)
	case reflect.Int64:
		return binder.setIntField(valueStr, 64, value)
	case reflect.Uint:
		return binder.setUintField(valueStr, 0, value)
	case reflect.Uint8:
		return binder.setUintField(valueStr, 8, value)
	case reflect.Uint16:
		return binder.setUintField(valueStr, 16, value)
	case reflect.Uint32:
		return binder.setUintField(valueStr, 32, value)
	case reflect.Uint64:
		return binder.setUintField(valueStr, 64, value)
	case reflect.Bool:
		return binder.setBoolField(valueStr, value)
	case reflect.Float32:
		return binder.setFloatField(valueStr, 32, value)
	case reflect.Float64:
		return binder.setFloatField(valueStr, 64, value)
	case reflect.String:
		value.SetString(valueStr)
	default:
		return errors.New("unknown type")
	}
	return nil
}


func (binder *DefaultBinder) setIntField(value string, bitSize int, field reflect.Value) error {
	if value == "" {
		value = "0"
	}
	intVal, err := strconv.ParseInt(value, 10, bitSize)
	if err == nil {
		field.SetInt(intVal)
	}
	return err
}

func(binder *DefaultBinder)  setUintField(value string, bitSize int, field reflect.Value) error {
	if value == "" {
		value = "0"
	}
	uintVal, err := strconv.ParseUint(value, 10, bitSize)
	if err == nil {
		field.SetUint(uintVal)
	}
	return err
}

func (binder *DefaultBinder) setBoolField(value string, field reflect.Value) error {
	if value == "" {
		value = "false"
	}
	boolVal, err := strconv.ParseBool(value)
	if err == nil {
		field.SetBool(boolVal)
	}
	return err
}

func (binder *DefaultBinder) setFloatField(value string, bitSize int, field reflect.Value) error {
	if value == "" {
		value = "0.0"
	}
	floatVal, err := strconv.ParseFloat(value, bitSize)
	if err == nil {
		field.SetFloat(floatVal)
	}
	return err
}
