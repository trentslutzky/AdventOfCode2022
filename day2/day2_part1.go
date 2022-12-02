package main

import (
  "fmt"
  "strings"
  "bufio"
  "os"
)

func main() {
  data, _ := os.Open("./day2.txt")
  scanner := bufio.NewScanner(data)

  for scanner.Scan(){
    line := scanner.Text()
    line_split := strings.Split(line, " ")
    opponent_move := line_split[0] 
    my_move := line_split[1]
    fmt.Println(opponent_move, my_move)
  }
}
