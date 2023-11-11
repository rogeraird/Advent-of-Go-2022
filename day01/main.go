package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type inventory struct {
	calories []int
}

func newInventory() inventory {
	return inventory{
		calories: make([]int, 5),
	}
}

func (i inventory) calculateTotal() int {
	total := 0

	for _, c := range i.calories {
		total += c
	}

	return total
}

func main() {
	total_inventory := make([]inventory, 10)

	file, err := os.Open("../inputs/1.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	current_inventory := newInventory()

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			val, err := strconv.Atoi(line)

			if err != nil {
				panic(err)
			}

			current_inventory.calories = append(current_inventory.calories, val)
		} else {
			total_inventory = append(total_inventory, current_inventory)
			current_inventory = newInventory()
		}
	}

	highestCalories := 0

	for _, inv := range total_inventory {
		if inv.calculateTotal() > highestCalories {
			highestCalories = inv.calculateTotal()
		}
	}

	fmt.Println("Answer to part 1", highestCalories)

	sort.Slice(total_inventory, func(i, j int) bool {
		return total_inventory[i].calculateTotal() > total_inventory[j].calculateTotal()
	})

	top_3_sum := 0

	for i := 0; i < 3; i++ {
		top_3_sum += total_inventory[i].calculateTotal()
	}

	fmt.Println("Answer to part 2", top_3_sum)

}
