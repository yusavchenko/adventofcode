package main

import (
  "bufio"
)


func Part1(scanner *bufio.Scanner) int {
  sum := 0

  // Input
  red := 12
  green := 13
  blue := 14

  for scanner.Scan() {
    game, err := GameFromStr(scanner.Text())
    if err != nil {
      return 0
    }
    if game.isPossible(red, green, blue) {
      sum += game.GetId()
    }
  }
  return sum;
}

