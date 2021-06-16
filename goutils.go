package goutils

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type strInterface struct{}
type sliceInterface struct{}
type intInterface struct{}
type httpInterface struct{}

// FindInSlice takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func (strInterface) FindInSlice(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// InSlice takes a string and a slice. It will return bool.
func (strInterface) InSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// HasString takes slice and a string and checks if any of the strings in slice
// matching given string
func (sliceInterface) HasString(a []string, b string) (int, bool) {
	for i, v := range a {
		if v == b {
			return i, true
		}
	}
	return -1, false
}

// Contains takes slice and a string and checks if given string contains any
// of the strings from slice
func (sliceInterface) Contains(a []string, b string) (int, bool) {
	for i, v := range a {
		if strings.Contains(b, v) {
			return i, true
		}
	}
	return -1, false
}

// TrimSpace for all strings in slice
func (sliceInterface) TrimSpace(arr []string) {
	for i, a := range arr {
		arr[i] = strings.TrimSpace(a)
	}
}

// IndexOf takes a slice and looks for an element in it. If found it will
// return it's index, otherwise it will return -1 and a bool of false.
func (intInterface) IndexOf(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// InSlice takes an int and a slice. It will return bool.
func (intInterface) InSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Request - takes method as string, url as string, headers as map[string]string and boyd as io.Reader.
// It returns status code as int, response body as []byte and error
func (httpInterface) Request(method string, url string, headers map[string]string, body io.Reader) (int, []byte, error) {
	req, err := http.NewRequest(strings.ToUpper(method), url, body)
	if err != nil {
		return -1, nil, err
	}
	for key, val := range headers {
		req.Header.Set(key, val)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return -1, nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, nil, err
	}
	return resp.StatusCode, respBody, nil
}
