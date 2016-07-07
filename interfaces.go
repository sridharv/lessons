package main

import (
	"fmt"
)

// Dog implements Runner
type Dog struct {
	Name string
	Age  int
}

func (d Dog) Run(distance int) int {
	return distance * d.Age
}

func (d Dog) String() string { return fmt.Sprintf("Doggy %s (aged %d)", d.Name, d.Age) }

// RaceDogs takes a distance and a slice of Dogs and returns the index of the winner
func RaceDogs(distance int, dogs []Dog) int {
	min, winner := dogs[0].Run(distance), 0

	for i, d := range dogs[1:] {
		timeTaken := d.Run(distance)
		if timeTaken < min {
			min, winner = timeTaken, i+1
		}
	}
	return winner
}

// RaceHumans takes a distance and a slice of Humans and returns the index of the winner
func RaceHumans(distance int, humans []Human) int {
	min, winner := humans[0].Run(distance), 0
	for i, h := range humans[1:] {
		timeTaken := h.Run(distance)
		if timeTaken < min {
			min, winner = timeTaken, i+1
		}
	}
	return winner
}

type Human struct {
	Name string
	Fit  bool
}

func (m Human) String() string { return fmt.Sprintf("Human %s (fit: %v)", m.Name, m.Fit) }

func (m Human) Run(distance int) int {
	if m.Fit {
		return 50 * distance
	}
	return 1000 * distance
}

// Runner is the interface holding the Run method.
type Runner interface {
	Run(distance int) int
}

// Race takes a distance and a slice of Runners and returns the index of the winner.
func Race(distance int, contestants []Runner) int {
	min, winner := contestants[0].Run(distance), 0
	for i, c := range contestants[1:] {
		timeTaken := c.Run(distance)
		if timeTaken < min {
			min, winner = timeTaken, i+1
		}
	}
	return winner
}

func main() {
	// contestants1 is a slice of Dog
	contestants1 := []Dog{
		Dog{"Jeff", 5}, Dog{"Mapillai", 10}, Dog{"Mama", 11},
	}
	winner := contestants1[RaceDogs(100, contestants1)]
	fmt.Println(winner)

	// contestants2 is a slice of Human
	contestants2 := []Human{
		Human{"Hajmola", false}, Human{"Rudramurthi", true},
	}
	winner2 := contestants2[RaceHumans(100, contestants2)]
	fmt.Println(winner2)

	// contestants3 is a slice of Runner
	contestants3 := []Runner{
		Dog{"Jeff", 5}, Dog{"Mapillai", 10}, Human{"Hajmola", false}, Human{"Rudramurthi", true},
	}
	winner3 := contestants3[Race(100, contestants3)]
	fmt.Println(winner3)
}
