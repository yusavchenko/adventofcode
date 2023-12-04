package main

import "strings"

func parseText(text string) ([]string, []string) {
  content := strings.Split(text, ": ")[1];
  splitContent := strings.Split(content, " | ");
  winItems := strings.Split(splitContent[0], " ");
  items := strings.Split(splitContent[1], " ");

  return winItems, items;
}

func getWinMap(winItems []string) map[string]bool {
    winMap := make(map[string]bool);

    for _, item := range winItems {
      key := strings.Trim(item, " ");
      if key != "" {
        winMap[item] = true;
      }
    }
    return winMap;
}

func getWinCount(winMap map[string]bool, items []string) int {
    res := 0;
    for _, item := range items {
      if winMap[item] == true {
        res += 1;
      }
    }
    return res;
}

