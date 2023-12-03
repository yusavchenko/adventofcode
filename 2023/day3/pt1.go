package main

import (
	"regexp"
	"strconv"
	"strings"
)

func Part1(file []byte) int {
  sum := 0
  rows := strings.Split(string(file), "\n")
  re := regexp.MustCompile(`\d+`)
  for y, row := range rows {
    nbrIndexes := re.FindAllIndex([]byte(row), -1)
    for _, index := range nbrIndexes {
      nbr, nbrErr := strconv.Atoi(string(row[index[0]:index[1]]));
      if nbrErr == nil {
        for x := index[0]; x < index[1]; x++ {
          if hasSymbolNeighbour(rows, x, y) {
            sum += nbr;
            break;
          }
        }
      }
    }
  }
  return sum;
}

func hasSymbolNeighbour(rows []string, x int, y int) bool {

  directions := [][]int{
    {1, 0},
    {-1, 0},
    {0, 1},
    {0, -1},
    {1, 1},
    {-1, -1},
    {1, -1},
    {-1, 1},
  }

  for _, direction := range directions {
    xIndex := x + direction[0];
    yIndex := y + direction[1];
    if xIndex < 0 || yIndex < 0 ||  yIndex >= len(rows)|| xIndex >= len(rows[yIndex])  {
      continue;
    }
    if string(rows[yIndex][xIndex]) != "." {
      _, nbrErr := strconv.Atoi(string(rows[yIndex][xIndex]));
      if nbrErr == nil {
        continue;
      }
      return true;
    }
  }

  return false;
}

