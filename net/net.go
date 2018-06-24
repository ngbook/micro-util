package net

/**
 * Author: jsongo<jsongo@qq.com>
 */

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/emicklei/go-restful"
)

// InitRsp generates the common response.
func InitRsp(statusCode int64, rsp *restful.Response, opts ...interface{}) {
	if rsp == nil {
		return
	}
	var code int64
	if statusCode == 200 {
		// rsp.StatusCode = 200
		code = 0
	} else if statusCode > 0 {
		code = statusCode
	} else {
		code = 1500 // default error code
	}
	result := map[string]interface{}{
		"code": code,
	}
	if len(opts) > 0 { // data to add to response
		data := opts[0]
		if code == 0 {
			result["data"] = data // data to return
		} else {
			result["errMsg"] = data // error message if status is not 200
			if len(opts) > 1 {      // data to return
				result["data"] = opts[1]
			}
		}
	}
	rsp.WriteEntity(result)
}

// Fetch data from internet
func Fetch(url string) []byte {
	rsp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return body
}

// Post simply fetch data from net
func Post(url string, data []byte, contentType string) []byte {
	rsp, err := http.Post(url, contentType, bytes.NewReader(data))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return body
}
