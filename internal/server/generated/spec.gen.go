// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8xYbW/T1hf/Ktb58+K/yWnSJ+j8BhUSWCREWVr2pmTgOJfGkNjGvkFUkaU2YQMJUDc2",
	"CWnSmDa+QFo1Iysk/QrnfqPp3JsHJ3ZLtnUbbyL7+vrc3+93Hp0GWG7Ncx3m8ACMBnimb9YYZ768y2fp",
	"t8wCy7c9brsOGIBvsIddPMI2vhPPsSee4e8advEQO9gTTeyKx+qxaGJf7Gh4jH258y328UAud/Cd2AMd",
	"bLLnmbwCOjhmjYEBdhl08NmDuu2zMhjcrzMdAqvCaiZBOeezu2DA/9Jj1Gn1NEjfvJnPQhiG9H7guU7A",
	"JIec77s+XViuw5nD6dL0vKptmcQofS8gWg1gj8yaV2VqZ5mwXF7L5m7nCoW1AuhQY0FgbtEyvsZj7Ipd",
	"bEshOhr2xVPs4j4eYRfCWeEqXBLvlMCvR+baGh6IZ0pB7JHA2BV7Gr7FNh6LHewTCjpyvW5ZLAhmoDkb",
	"uAILvKHNJIivxgA00RK7eIwd8RR72I8hhnAkyaQ/pkz+gB3RHAYVtolmF3viMQURvse2+EZx7+N+VPEO",
	"6OD5rsd8biuHm5z7dqnOWcIh34qm2KE3RUs0dU3saniEfRWp4hm+18QuHog9KXCP5B8f1AYd+LZHERBw",
	"33a2SHcVKbFjfsQ+Hk7EhXHL0T7V7lxazd4u5L64mVvfuKOlNPwJOxLAjtihK9Ec5lPEwxoeEayO2MED",
	"+hUtZSt/fSNXuL567fZ6rvBlrqAiVVp9iT3iJ232sCf2pG6RqIqaw/YtB3RgTr0GxiZEEIIOiWdAUR9n",
	"y9QLMYVGeRMT6dQ0ilkKo2VhUwlfHO1yS/eYxem8AntQZwG/7DOTsw0zuJ9w8q9UgShUJ/JIw0OqSyRO",
	"H9/ioQI2TLdDbIsnEthktE2YbsT5c5uronI6H7UtmdA4FeNUXopd8Z0mvpZZQgL2YwjdBAnc+xDxoaqy",
	"g5NLrltlphND6N4/ER7V2tkEp1Q7UIn+z8htl2drEzM7Rjak071D9K8yTtyDmcmrrtgSTyaruuRNrZGz",
	"WvAhKlLtcITK9H1zm+6HXvi4pZPW/u54EQ1juHC+XF5mpZWUtbC4nFqaL5VSK9Z8JrWwcIGVMxdKlrW0",
	"TK3cfHSNOVu8AsbieR1qthO99UzOmU9QvtpsFC9uZlKfmam7q6krxcZKmPp/9H4pTH3SWAyjS/ML4WZY",
	"vHguoYCRxM5dN0559UZ+mAy4L17gkXiBPYoETRanLkUP/jaKk4217Nq1/PqGLNsDV8BwUVu9kQcdHjI/",
	"UMYzc5m5eZLb9ZhjejYYsDiXmVsEybQiIyNdYWaV+Ddgi8nJgeJGzg35MhhwlfHP1Y6p2WohkzkpZkb7",
	"0uM5QoegXquZ/jb5+WcqWIMWFO9Jmmy/dIX7gwZNQ0ZfVot9GjTErmiK5ySBuRVQuOEv1E7kePlexk0H",
	"inRmmg8T8yRyKnOTuZ3ZODVRJZJmqu+lh4l9l5o8FQpN5sGxeCJa5MKlWeQezJU6LP+J3VOuGdWmLmk8",
	"XZFHgr+KLBdDHTw3SBA40hlUfWABv+SWt89Q3emWH06WIupvYcy982fu3mkEH6uDL486roQw8q5onezd",
	"USKlG3Y5VEWsytSQPenvrFwf+Dv6PbmZjHW8JU3tpRhz1FJSl5j45hAtyeadrJXdSUb/Yt68GaKYyprT",
	"dNXBqyckzRXbsYPKGYo4/0ERJd4j+R30X4r414qPNML8h0ONJplm2UNNPQUd6n4VDKhw7gVGOl11LbNa",
	"cQNurGRWMlK5wRGN4X8SCY0l1MdPo0DCYvhHAAAA///oeJeHUhEAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}