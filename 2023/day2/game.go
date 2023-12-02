package main

import (
  "fmt"
  "strconv"
  "strings"
)

type Round struct {
  blue int
  red int
  green int
}

type Game struct {
  id int
  rounds []Round
}

func (g *Game) GetId() int {
  return g.id;
}

// Idk how to name this
func (g *Game) LessPossibleNumbersMultiply() int {
  red := g.rounds[0].red
  blue := g.rounds[0].blue
  green := g.rounds[0].green
  for _, round := range g.rounds {
    if round.red > red {
      red = round.red
    }
    if round.blue > blue {
      blue = round.blue
    }
    if round.green > green {
      green = round.green
    }
  }
  return red * blue * green;
}

func (g *Game) Print() {
  fmt.Println("Game: ", g.id);
  for _, round := range g.rounds {
    fmt.Printf("Round: %d red, %d blue, %d green\n", round.red, round.blue, round.green);
  }
}

func (g *Game) isPossible(red int, green int, blue int) bool {
  for _, round := range g.rounds {
    if round.red > red || round.blue > blue || round.green > green {
      return false;
    }
  }
  return true;
}

func GameFromStr(str string) (*Game, error) {
  str1 := strings.Split(str, ":");
  gamestr := str1[0];
  gameID := strings.Split(gamestr, " ")[1];
  id, startErr := strconv.Atoi(string(gameID));

  if startErr != nil {
    return nil, startErr
  }

  setsStr := strings.Split(str1[1], ";");
  game := &Game{
    id: id,
    rounds: make([]Round, 0),
  }

  for _, setStr := range setsStr {
    cubes := strings.Split(setStr, ",");
    for _, cube := range cubes {
      splitCube := strings.Split(strings.TrimLeft(cube, " "), " ")
      color := splitCube[1];
      count, err := strconv.Atoi(splitCube[0]);
      round := Round{}
      if err != nil {
        return nil, err
      }

      switch color {
      case "red":
        round.red += count;
      case "blue":
        round.blue += count;
      case "green":
        round.green += count;
      }
      game.rounds = append(game.rounds, round);
    }
  }

  return game, nil
}
