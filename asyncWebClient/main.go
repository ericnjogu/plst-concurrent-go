package main

import (
  "net/http"
  "io/ioutil"
  "fmt"
  "encoding/json"
  "time"
)

func makeRequest(url string) (map[string]interface{}, error) {
  var headers = make(map[string]interface{})
  resp, respErr := http.Get(url)
  if respErr != nil {
    return headers, respErr
  }

  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)

  // https://www.socketloop.com/tutorials/golang-how-to-convert-json-string-to-map-and-slice
  // https://tour.golang.org/methods/14
  err := json.Unmarshal([]byte(body), &headers)
  return headers, err
}

func main() {
  start := time.Now()
  for i := 0; i <= 99; i++ {
    _, err := makeRequest("http://localhost:1920")
    if err != nil {
      fmt.Println("%v", err)
    } else {
      fmt.Printf("\rmade requests: %v", i)
    }
  }
  fmt.Println("\ntime taken ", time.Since(start))
}
