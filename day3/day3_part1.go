package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  data, _ := os.Open("./day3.txt")
  scanner := bufio.NewScanner(data)

  for scanner.Scan(){
    line := scanner.Text()
    fmt.Println(len(line))
  }
}
