
package main

import (
  "bufio"
  "fmt"
  "strconv"
  "strings"
)

func Part1(scanner *bufio.Scanner) int {
  sum := 0
  for scanner.Scan() {
    text := scanner.Text();
    runes := []rune(text);
    rLength := len(runes)
    val := make([]string, 2);

    for i := 0; i < rLength ; i++ {
      endIndex := rLength - 1 - i;
      startNumber, startErr := strconv.Atoi(string(runes[i]));
      endNumber, endErr := strconv.Atoi(string(runes[endIndex]));

      if startErr == nil && val[0] == "" {
        val[0] = fmt.Sprint(startNumber)
      }
      if endErr == nil && val[1] == "" {
        val[1] = fmt.Sprint(endNumber)
      }

      if val[0] != "" && val[1] != "" {
        concatenated := strings.Join(val, "");
        res, resErr := strconv.Atoi(concatenated);
        if resErr == nil {
          sum += res;
        }
        break;
      }
    }
  }
  return sum;
}

