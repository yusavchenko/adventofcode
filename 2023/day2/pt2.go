package main

import (
  "bufio"
)


func Part2(scanner *bufio.Scanner) int {
  sum := 0
  for scanner.Scan() {
    game, err := GameFromStr(scanner.Text())
    if err != nil {
      return 0
    }
    sum += game.LessPossibleNumbersMultiply()
  }
  return sum;
}

