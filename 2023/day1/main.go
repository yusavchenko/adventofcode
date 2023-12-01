package main

import (
  "bufio"
  "flag"
  "fmt"
  "os"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}


// Example of how to run:
func main(){
  isDemo := false
  part := 1
  flag.BoolVar(&isDemo, "d", false, "is demo mode")
  flag.IntVar(&part, "p", 1, "Part 1 or 2")
  flag.Parse();

  fileName := "input.txt"
  if isDemo {
    fileName = "demo.txt"
  }
  fileName = fmt.Sprintf("assets/part%d/%s", part, fileName);
  dat, err := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)
  check(err)
  defer dat.Close()
  scanner := bufio.NewScanner(dat)

  sum := 0
  if part == 1 {
    sum = Part1(scanner);
  } else if part == 2 {
    sum = Part2(scanner);
  }
  fmt.Println("Sum: ", sum)

}

