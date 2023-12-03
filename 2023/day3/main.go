package main

import (
  "flag"
  "fmt"
  "os"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

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
  fileName = fmt.Sprintf("assets/%s", fileName);
  dat, err := os.ReadFile(fileName)
  check(err)

  sum := 0
  if part == 1 {
    sum = Part1(dat)
  } else {
    sum = Part2(dat)
  }

  fmt.Println("Sum: ", sum)

}

