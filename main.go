package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mapi/route"
	"net/http"
	"time"
)

// Message is used for extracting json
type Message struct {
	Body string
	Time int32
}

// Must is used for checking request
type Must struct {
	BeRight bool
	Msg string
}

var must Must{BeRight: true} 

func main() {
	Routers := route.Link()

	for _, Router := range Routers {

		f := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-type", "application/json")
			output, err := Router.Request
			if err != nil {
				http.Error(w, "404", http.StatusInternalServerError)
			}
			fmt.Fprintf(w, output)
		}

		http.HandleFunc(Router.Url, handlerwrapper(f))
	}

	http.ListenAndServe(":8080", nil)
}

// HandlerWrapper will validate Must and wrap function into handler
func HandlerWrapper(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer log.Println("After")
		must.CheckToken(&r).CheckNotAbuse()
		if must.BeRight == true {
			fn(w, r)
		} else {
			http.Error(w, mst.Msg, http.StatusInternalServerError)
		}
	}
}

// CheckToken will validate token
func (must *Must) CheckToken(r *http.Request) *Must {
	// Ex: api/hello?token=1234
	// TODO: token should be generate silently in shell script => os.Env()

	if err := req.ParseForm(); err != nil {
			must.BeRight = false
			must.Msg = "Cannot parse URL. Please double check"
	} else if token := req.Form.Get("token"); token != "1234" {
		must.BeRight = false
		must.Msg = "Invalid Token"
	}

	return must
}

// CheckNotAbuse will check whether the client is abusing api
func (must *Must) CheckNotAbuse(r *http.Request) *Must {
	if must.BeRight = false {
		return must
	}

	// TODO: Add interval check to prevent spam
	// or check Server Resource 100% => deny request
	// Msg => Stop Abusing
	return must
}

func ExampleResponse() (string, error) {
	unixtime := int32(time.Now().Unix())
	msg := Message{"Hi", "Hello Web!", unixtime}
	jbMsg, err := json.Marshal(msg)

	if err != nil {
		return "", err
	}

	jsonMsg := string(jbMsg[:]) // converting byte array to string
	return jsonMsg, nil
}
