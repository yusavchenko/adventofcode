package main

import (
  "regexp"
  "strconv"
  "strings"
  "fmt"

)

func Part2(file []byte) int {
  sum := 0
  rows := strings.Split(string(file), "\n")
  re := regexp.MustCompile(`\d+`)
  pairMap := make(map[string][]int)

  for y, row := range rows {
    nbrIndexes := re.FindAllIndex([]byte(row), -1)
    for _, index := range nbrIndexes {
      nbr, nbrErr := strconv.Atoi(string(row[index[0]:index[1]]));
      if nbrErr == nil {
        for x := index[0]; x < index[1]; x++ {
          index, err := getStarIndexNeighbour(rows, x, y);
          if err == nil {
            starKey := fmt.Sprintf("%d_%d", index[0], index[1]);

            pairMap[starKey] = append(pairMap[starKey], nbr);
            if len(pairMap[starKey]) == 2 {
              sum += pairMap[starKey][0] * pairMap[starKey][1];
            }
            break;
          }
        }
      }
    }
  }
  return sum;
}

func getStarIndexNeighbour(rows []string, x int, y int) ([]int, error) {

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
    if string(rows[yIndex][xIndex]) == "*" {
      return []int{xIndex, yIndex}, nil;
    }
  }

  return nil, fmt.Errorf("No star found");

}

