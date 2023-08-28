package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"reflect"
	"strings"

	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcenginebody"
)

var bypassBuildHandler = request.NamedHandler{Name: "BypassBuildHandler", Fn: bypassBuild}

func bypassBuild(r *request.Request) {
	body := url.Values{}

	params := r.Params
	if reflect.TypeOf(r.Params) == reflect.TypeOf(&map[string]interface{}{}) {
		if v, ok := (*r.Params.(*map[string]interface{}))[BypassInfoUrlParam]; ok {
			for k1, v1 := range v.(map[string]string) {
				body.Add(k1, v1)
			}
		}
		if v, ok := (*r.Params.(*map[string]interface{}))[BypassInfoInput]; ok {
			params = v
		}
	}

	r.Params = params

	r.HTTPRequest.Host = r.HTTPRequest.URL.Host
	if r.Config.ExtraUserAgent != nil && *r.Config.ExtraUserAgent != "" {
		if strings.HasPrefix(*r.Config.ExtraUserAgent, "/") {
			request.AddToUserAgent(r, *r.Config.ExtraUserAgent)
		} else {
			request.AddToUserAgent(r, "/"+*r.Config.ExtraUserAgent)
		}
	}
	contentType := r.HTTPRequest.Header.Get("Content-Type")
	if strings.ToUpper(r.HTTPRequest.Method) == "PUT" ||
		strings.ToUpper(r.HTTPRequest.Method) == "POST" ||
		strings.ToUpper(r.HTTPRequest.Method) == "DELETE" ||
		strings.ToUpper(r.HTTPRequest.Method) == "PATCH" {
		r.HTTPRequest.Header.Set("Content-Type", contentType)
		if strings.Contains(strings.ToLower(contentType), "application/json") {
			if r.HTTPRequest.Header.Get("Content-Length") == "" {
				volcenginebody.BodyJson(&body, r)
			}
		}
	} else {
		if len(contentType) > 0 && !strings.Contains(strings.ToLower(contentType), "x-www-form-urlencoded") {
			r.HTTPRequest.Header.Del("Content-Type")
		}
		volcenginebody.BodyParam(&body, r)
		if len(contentType) > 0 {
			r.HTTPRequest.Header.Set("Content-Type", contentType)
		}
	}
}

var bypassUnmarshalHandler = request.NamedHandler{Name: "BypassUnmarshalHandler", Fn: bypassUnmarshal}

func bypassUnmarshal(r *request.Request) {
	defer r.HTTPResponse.Body.Close()
	if r.DataFilled() {
		body, err := ioutil.ReadAll(r.HTTPResponse.Body)
		if err != nil {
			fmt.Printf("read volcenginebody err, %v\n", err)
			r.Error = err
			return
		}

		if reflect.TypeOf(r.Data) == reflect.TypeOf(&map[string]interface{}{}) {
			(*r.Data.(*map[string]interface{}))[BypassHeader] = r.HTTPResponse.Header
			temp := make(map[string]interface{})
			if len(body) == 0 {
				(*r.Data.(*map[string]interface{}))[BypassResponse] = temp
				return
			}

			if strings.Contains(strings.ToLower(r.HTTPResponse.Header.Get("Accept")), "application/json") ||
				strings.Contains(strings.ToLower(r.HTTPResponse.Header.Get("Content-Type")), "application/json") {
				if err = json.Unmarshal(body, &temp); err != nil {
					fmt.Printf("Unmarshal err, %v\n", err)
					r.Error = err
					return
				}
				(*r.Data.(*map[string]interface{}))[BypassResponse] = temp
			} else {
				(*r.Data.(*map[string]interface{}))[BypassResponse] = temp
				//(*r.Data.(*map[string]interface{}))[TosPlainResponse] = string(body)
			}

			(*r.Data.(*map[string]interface{}))[BypassResponseData] = string(body)

		}

	}
}
