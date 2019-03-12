package meerkat

import "net/http"
import "encoding/xml"

type Response struct{
	writer http.ResponseWriter
	committed   bool
}

func (resp *Response) Reset(writer http.ResponseWriter){
	resp.writer = writer
	resp.committed = false
}

func (resp *Response) setHeaderAndCode(code int, contentType string) {
	if resp.committed {
		LogInstance().Warnln("response already committed...")
		return
	}
	resp.writer.Header().Set("Content-Type", contentType)
	resp.writer.WriteHeader(code)
	resp.committed = true
}

func (resp *Response) Write(code int, contentType string, data []byte)  error{
	resp.setHeaderAndCode(code,contentType)
	_, err := resp.writer.Write(data)
	return err
}

func (resp *Response) HTML(code int, data []byte) error{
	return resp.Write(code, "text/html;charset=UTF-8", data)
}

func (resp *Response) JSON(code int, data []byte) error{
	return resp.Write(code, "application/json;charset=UTF-8", data)
}

func (resp *Response) JavaScript(code int, data []byte) error{
	return resp.Write(code, "application/javascript;charset=UTF-8", data)
}

func (resp *Response) XML(code int, contentType string,data []byte) error{
	resp.setHeaderAndCode(code,contentType)
	if _, err := resp.writer.Write([]byte(xml.Header)); err != nil {
		return err
	}
	_, err := resp.writer.Write(data)
	return err
}
