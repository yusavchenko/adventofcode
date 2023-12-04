
package main

import (
  "bufio"
)

func Part1(scanner *bufio.Scanner) int {
  sum := 0
  for scanner.Scan() {
    res := 0;
    winItems, items := parseText(scanner.Text());
    winMap := getWinMap(winItems)

    for _, item := range items {
      if winMap[item] == true {
        if res == 0 {
          res = 1;
        } else {
          res = res * 2;
        }
      }
    }
    sum += res;
  }
  return sum;
}

