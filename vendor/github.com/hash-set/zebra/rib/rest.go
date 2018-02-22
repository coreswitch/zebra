// Copyright 2016 Zebra Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rib

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/coreswitch/component"
)

type Resource interface {
	Get(url string, queries url.Values) (Status, interface{})
	Post(url string, queries url.Values) (Status, interface{})
	Put(url string, queries url.Values) (Status, interface{})
	Delete(url string, queries url.Values) (Status, interface{})
}

type apiheader struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
type apienvelope struct {
	Header   apiheader   `json:"header"`
	Response interface{} `json:"response"`
}

type Status struct {
	success bool
	code    int
	message string
}

type ResourceBase struct{}

func (ResourceBase) Get(url string, queries url.Values) (Status, interface{}) {
	return FailCode(http.StatusMethodNotAllowed), nil
}

func (ResourceBase) Post(url string, queries url.Values) (Status, interface{}) {
	return FailCode(http.StatusMethodNotAllowed), nil
}

func (ResourceBase) Put(url string, queries url.Values) (Status, interface{}) {
	return FailCode(http.StatusMethodNotAllowed), nil
}

func (ResourceBase) Delete(url string, queries url.Values) (Status, interface{}) {
	return FailCode(http.StatusMethodNotAllowed), nil
}

func ResourceHandler(resource Resource) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b := bytes.NewBuffer(make([]byte, 0))
		io.TeeReader(r.Body, b)

		r.Body = ioutil.NopCloser(b)
		defer r.Body.Close()

		r.ParseForm()

		var status Status
		var data interface{}

		switch r.Method {
		case "GET":
			status, data = resource.Get(r.URL.Path, r.Form)
		case "POST":
			status, data = resource.Post(r.URL.Path, r.Form)
		case "PUT":
			status, data = resource.Put(r.URL.Path, r.Form)
		case "DELETE":
			status, data = resource.Delete(r.URL.Path, r.Form)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var content []byte
		var e error

		if !status.success {
			content, e = json.Marshal(apienvelope{
				Header: apiheader{Status: "fail", Message: status.message},
			})
		} else {
			// content, e = json.Marshal(apienvelope{
			// 	Header:   apiheader{Status: "success"},
			// 	Response: data,
			// })
			content, e = json.Marshal(data)
		}
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status.code)
		w.Write(content)
	}
}

type InterfacesStateHandler struct {
	ResourceBase
}

// (ietf-)interfaces/interfaces-state/[name]/statistics
func (InterfacesStateHandler) Get(url string, queries url.Values) (Status, interface{}) {
	paths := strings.Split(url, "/")

	if len(paths) < 4 {
		return FailCode(http.StatusMethodNotAllowed), nil
	}

	//fmt.Printf("Interface name: %s\n", paths[3])
	v := VrfDefault()

	ifp := v.IfLookupByName(paths[3])
	if ifp == nil {
		return FailCode(http.StatusNotFound), nil
	}

	if len(paths) >= 5 && paths[4] == "statistics" {
		IfStatsUpdate()
		return Status{success: true, code: http.StatusOK}, &ifp.Stats
	}

	return FailCode(http.StatusNotFound), nil
}

func FailCode(code int) Status {
	return Status{success: false, code: code, message: strconv.Itoa(code) + " " + http.StatusText(code)}
}

type RestComponent struct {
}

func NewRestComponent() *RestComponent {
	return &RestComponent{}
}

func (r *RestComponent) Start() component.Component {
	go func() {
		http.HandleFunc("/", ResourceHandler(ResourceBase{}))
		http.HandleFunc("/interfaces/interfaces-state/", ResourceHandler(InterfacesStateHandler{}))
		http.ListenAndServe(":3000", nil)
	}()
	return r
}

func (r *RestComponent) Stop() component.Component {
	return r
}
