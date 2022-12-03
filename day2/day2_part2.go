package main

import (
  "fmt"
  "strings"
  "bufio"
  "os"
)

//    A  B  C
// X  
// Y  
// Z

var score_matrix = [3][3]int{
  {3, 1, 2} ,// lose/scissors    lose/rock      lose/paper
  {4, 5, 6} ,// draw/rock        draw/paper     draw/scissors
  {8, 9, 7} ,// win/paper        win/scissors   win/rock
}

func move_to_index(move string) int {
  if (move == "A" || move == "X") { return 0 }
  if (move == "B" || move == "Y") { return 1 }
  if (move == "C" || move == "Z") { return 2 }
  return -1
}

func main() {
  data, _ := os.Open("./day2.txt")
  scanner := bufio.NewScanner(data)
  total := 0

  for scanner.Scan(){
    line := scanner.Text()
    line_split := strings.Split(line, " ")
    opponent_move := line_split[0] 
    my_move := line_split[1]
    round_score := score_matrix[move_to_index(my_move)][move_to_index(opponent_move)]
    total += round_score
    fmt.Println(round_score, total)
  }

  fmt.Println("TOTAL:", total)
}
