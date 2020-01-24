package main

import (
  "time"
)

func main() {
  go func () {
    for i := 0; i < 100; i++ {
      println("hello")
    }
  }()

  go func() {
    for i := 0; i < 100; i++ {
      println("go")
    }
  }()

  dur, _ := time.ParseDuration("1s")
  time.Sleep(dur)
}
