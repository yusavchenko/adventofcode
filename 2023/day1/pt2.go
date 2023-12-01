package main

import (
  "bufio"
  "fmt"
  "strconv"
  "strings"
)

func Part2(scanner *bufio.Scanner) int {
  sum := 0
  for scanner.Scan() {
    text := scanner.Text();
    val := make([]string, 2);

    for i := 0; i < len(text) ; i++ {
      number, nbrErr := getNumber(text, i)

      if nbrErr != nil {
        continue;
      }

      if val[0] == "" {
        val[0] = fmt.Sprint(number);
      }
      val[1] = fmt.Sprint(number);
    }
    concatenated := strings.Join(val, "");
    res, resErr := strconv.Atoi(concatenated);
    if resErr == nil {
      sum += res;
    }
  }
  return sum;
}

func getNumber(text string, index int) (int, error) {
  numberChar, nbrCharErr := strconv.Atoi(string(text[index]));
  parsedNumberFromText, parsedNumberErr := getNumberFromLetters(text, index);

  if nbrCharErr != nil && parsedNumberErr != nil {
    return 0, fmt.Errorf("No number found")
  }
  if nbrCharErr == nil {
    return numberChar, nil
  }
  if parsedNumberErr == nil {
    return parsedNumberFromText, nil
  }
  return 0, fmt.Errorf("No number found")
}


func getNumberFromLetters(str string, index int) (int, error) {
  strMap := map[string]int{
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
    "zero": 0,
  }
  slicesStr := str[index:]
  for key, val := range strMap {
    if strings.HasPrefix(slicesStr, key) {
      return val, nil
    }
  }
  return 0, fmt.Errorf("No number found")
}

