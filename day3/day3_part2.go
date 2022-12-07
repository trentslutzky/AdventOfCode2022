package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

var LETTERS = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func indexOfLetter(element string) (int) {
   for k, v := range LETTERS {
       if element == v {
           return k + 1
       }
   }
   return -1
}

func letterInSlice(slice []string, element string) (bool) {
  for _, v := range slice {
    if element == v {
      return true
    }
  }
  return false
}

func main() {
  data, _ := os.Open("./day3.txt")
  scanner := bufio.NewScanner(data)
  total := 0

  for scanner.Scan(){
    line := scanner.Text()
    line_letters := strings.Split(line, "")
    fmt.Println(line)
    duplicate_value := 0
    var line_half []string
    for i, l := range line_letters {
      if(i < len(line_letters)/2) {
        line_half = append(line_half,l)
      } else {
        if(letterInSlice(line_half,l)) {
          duplicate_value = indexOfLetter(l)
        }
      }
    }
    fmt.Println(duplicate_value)
    total += duplicate_value
  }

  fmt.Println("TOTAL:", total)
}
