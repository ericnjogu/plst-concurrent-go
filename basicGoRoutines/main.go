package main

import (
  "time"
)

func main() {
  go func () {
    println("hello")
  }()

  go func() {
    println("go")
  }()

  dur, _ := time.ParseDuration("1s")
  time.Sleep(dur)
}
