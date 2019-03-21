package meerkat

import "net/http"
import "encoding/xml"
import "encoding/json"
import "io"
import "os"
import "errors"

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

func (resp *Response) Bytes(code int, data []byte) error{
	return resp.Write(code, ContentTypePlain+";"+CharsetUTF8, data)
}

func (resp *Response) HTML(code int, data []byte) error{
	return resp.Write(code, ContentTypeXML2+";"+CharsetUTF8, data)
}

func (resp *Response) JSON(code int, data []byte) error{
	return resp.Write(code, ContentTypeJSON+";"+CharsetUTF8, data)
}

func (resp *Response) JSONEncode(code int, data interface{}, indent string) error{
	enc := json.NewEncoder(resp.writer)
	if(indent != ""){
		enc.SetIndent("",indent)
	}
	resp.setHeaderAndCode(code,ContentTypeJSON+";"+CharsetUTF8)
	return enc.Encode(data)
}

func (resp *Response) JSONP(code int, callback string,data []byte) (err error){
	resp.setHeaderAndCode(code,ContentTypeJSONP+";"+CharsetUTF8)
	if _, err = resp.writer.Write([]byte(callback + "(")); err != nil {
		return err
	}
	if _, err = resp.writer.Write(data); err != nil {
		return err
	}
	_, err = resp.writer.Write([]byte(")"))
	return err
}

func (resp *Response) JSONPEncode(code int, callback string,data interface{}) (err error){
	resp.setHeaderAndCode(code,ContentTypeJSONP+";"+CharsetUTF8)
	if _, err = resp.writer.Write([]byte(callback + "(")); err != nil {
		return err
	}
	enc := json.NewEncoder(resp.writer)
	if err = enc.Encode(data); err != nil {
		return
	}
	_, err = resp.writer.Write([]byte(")"))
	return err
}

func (resp *Response) XML(code int, data []byte) error{
	resp.setHeaderAndCode(code,ContentTypeXML+";"+CharsetUTF8)
	if _, err := resp.writer.Write([]byte(xml.Header)); err != nil {
		return err
	}
	_, err := resp.writer.Write(data)
	return err
}

func (resp *Response) XMLEncode(code int, data interface{}, indent string) error{
	enc := xml.NewEncoder(resp.writer)
	if(indent != ""){
		enc.Indent("",indent)
	}
	resp.setHeaderAndCode(code,ContentTypeXML+";"+CharsetUTF8)
	return enc.Encode(data)
}

func (resp *Response) Stream(code int, contentType string,reader io.Reader) error{
	resp.setHeaderAndCode(code,contentType)
	_, err := io.Copy(resp.writer, reader)
	return err;
}

func (resp *Response) Redirect(code int, url string) error {
	resp.writer.Header().Set("Location", url)
	resp.writer.WriteHeader(code)
	return nil
}

func (resp *Response) NoContent(code int) error {
	resp.writer.WriteHeader(code)
	return nil
}

func (resp *Response) File(fileName string,req *http.Request) error{
	file, err := os.Open(fileName)
	if err != nil {
		return errors.New("not find file or directory : "+fileName)
	}
	defer file.Close()

	info, _ := file.Stat()
	if info.IsDir() {
		return errors.New(fileName+" is a directory.")
	}

	http.ServeContent(resp.writer, req, info.Name(), info.ModTime(), file)
	return nil
}
