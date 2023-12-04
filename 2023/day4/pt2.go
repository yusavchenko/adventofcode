package main

import (
  "bufio"
)

func Part2(scanner *bufio.Scanner) int {
  sum := 0
  index := 0;
  resMap := make(map[int]int);
  for scanner.Scan() {
    resMap[index] += 1;
    winItems, items := parseText(scanner.Text());
    winCardsCount := getWinCount(getWinMap(winItems), items);

    for i := index + 1; i < winCardsCount + index + 1; i++ {
      resMap[i] += resMap[index]
    }
    sum += resMap[index];
    index += 1;
  }
  return sum ;
}

