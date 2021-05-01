# goutils

A collection of useful utilities for golang. 

## Available methods:
#### Str methods
1. ```golang
   Str.FindInSlice(slice []string, val string) (int, bool)
2. ```golang
   Str.InSlice(a string, list []string) bool
   
#### Int methods
1. ```golang
   Int.FindInSlice(slice []int, val int) (int, bool)
2. ```golang
   Int.InSlice(a int, list []int) bool
   
#### Slice methods
1.  Check if string(val) contains any string from slice
    ```golang
    Slice.InString(slice []int, val string) (int, bool)
    ```
   
#### Http methods
1. ```golang
   Http.Request(method string, url string, headers map[string]string,
    body io.Reader) (int, []byte, error)```

## Example

```go
package main

import (   
    "encoding/json"
    "bytes"
    "github.com/DMarinuks/goutils"
)

func main() {
  arr := []string{"foo", "bar"}
  if goutils.Str.InSlice("bar", arr) {
    // string in slice
  }
  postData, _ := json.Marshal(`{"foo":"bar"}`)
  statusCode, respBody, err := goutils.Http.Request(
        "POST", 
        "http://localhost:8080",
         map[string]string{"Authorization":"Token 123"},
         bytes.NewBuffer(postData),
  )
  
  a := []string{"bar", "foo"}
  b := "something with foo"
  i, ok := goutils.Slice.InString(a, b)

}
```
