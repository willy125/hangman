package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Print("asdasd")
	/*
		var num1 int
		var num2 int

		fmt.Println("Introduce el primer numero")
		fmt.Scanln(&num1)

		fmt.Println("Introduce el segundo numero")
		fmt.Scanln(&num2)

		fmt.Print("number 1 is: ", num1)
		fmt.Print("number 2 is: ", num2)
	*/
	/*
		zombies := []string{"Paul", "Raul", "Saul", "Jose", "Jared"}
		toRemove := 2
		zombiesSinSaul := append(zombies[:toRemove], zombies[toRemove+1:]...)

		zombiesMutados := []string{"Nemesis", "Tyrant", "DobermanZombie"}

		zombiesCopia := append(zombies, "Lucy", "Pablo")
		todosZombies := append(zombiesCopia, zombiesMutados...)
		fmt.Println(zombies)

		fmt.Println(todosZombies)
		fmt.Println(zombiesCopia)

		fmt.Println(zombiesSinSaul)
			for i := 0; i < len(zombies); i++ {
				fmt.Println(i+1, zombies[i])
			}
	*/
	/*primes := map[int]bool{
		2:  true,
		3:  false,
		4:  false,
		5:  true,
		52: true,
		53: true,
		1:  true,
	}*/
	numeros := map[int]float64{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
	}
	numeros[6] = 6
	for key, value := range numeros {
		numeros2 := map[int]float64{}
		numeros2[key] = math.Pow(float64(value), 3)
		fmt.Println(key, "->", value, "->", numeros2)

	}

	Array := []string{"letra a", "letra b"}
	Array2 := []string{"letra c", "letra d"}
	Array3 := []string{"letra e", "letra f"}
	Array4 := []string{"letra f", "letra g", "letra h", "letra i"}
	Array = append(Array, Array2...)
	Array = append(Array, Array3...)
	Array = append(Array, Array4...)
	Array = append(Array, "letra J")

	type Student struct {
		Name string
		Age  int
	}
	var Antonio Student
	Antonio.Name = "Antonio"
	Antonio.Age = 40

	fmt.Println(Antonio)

	//fmt.Println(Array)
	//fmt.Println(Array2)
	//fmt.Println(Array3)
}
