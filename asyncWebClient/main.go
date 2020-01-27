package main

import (
  "net/http"
  "io/ioutil"
  "fmt"
  "encoding/json"
  "os"
  "time"
)

func main() {
  start := time.Now()
  resp, respErr := http.Get("http://localhost:1920")
  if respErr != nil {
    println("request failed")
    fmt.Println("%v", respErr)
    os.Exit(2)
  }

  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)

  // https://www.socketloop.com/tutorials/golang-how-to-convert-json-string-to-map-and-slice
  // https://tour.golang.org/methods/14
  var headers = make(map[string]interface{})
  err := json.Unmarshal([]byte(body), &headers)
  if err != nil {
    println("Unmarshal failed")
    println(err)
  } else {
    fmt.Println("%v", headers)
    fmt.Println("time taken ", time.Since(start))
  }
}
