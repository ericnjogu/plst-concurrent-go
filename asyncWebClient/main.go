package main

import (
  "net/http"
  "io/ioutil"
  "fmt"
  "encoding/json"
  "time"
  "sync"
  "runtime"
)

func makeRequest(url string, wg *sync.WaitGroup) (map[string]interface{}, error) {
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
  wg.Done()
  return headers, err
}

func main() {
  runtime.GOMAXPROCS(4)
  var wg sync.WaitGroup
  start := time.Now()
  limit := 500
  wg.Add(limit)
  for i := 0; i < limit; i++ {
    go makeRequest("http://localhost:1920", &wg)
  }
  wg.Wait()
  fmt.Println("\ntime taken ", time.Since(start))
}
