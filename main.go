package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	// First Question

	value := 10
	fmt.Println("Memory Storage (Value I): ", &value)
	passValue(value)
	fmt.Println("Pass value (outside func): ", value) // remain unchanged

	eggInMyFridge := 20
	fmt.Println("Memory Storage (Reference I): ", &eggInMyFridge)
	passReference(&eggInMyFridge)
	fmt.Println("Pass ref (outside func), egg in my fridge:", eggInMyFridge) // changed

	// Second Question

	myPreparation := []string{"Toaster", "Blender Mixer", "Water Boiler"}
	for i := 0; i < len(myPreparation); i++ {

		tool := myPreparation[i]

		defer fmt.Println("Cleaning up my: ", tool, "and done")

		fmt.Println("Preparing up my: ", tool)
	}

	// Third Question

	var waitGroup sync.WaitGroup

	waitGroup.Add(3)
	// each meal will take random times
	go cooking("sandwich", &waitGroup)
	go cooking("coffee", &waitGroup)
	go cooking("salad", &waitGroup)
	waitGroup.Wait()

	fmt.Println("Finish cooking all meals, preparing to clean up...")
}

func passValue(num int) {
	num += 10
	fmt.Println("Memory Storage (Value II): ", &num)
	fmt.Println("Pass value (inside func): ", num)
}

func passReference(num *int) {
	*num -= 3
	fmt.Println("Memory Storage (Reference II): ", num)
	fmt.Println("Pass ref (inside func), egg in my fridge: ", *num)
}

func cooking(food string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Starting to cook: ", food)

	period := rand.Intn(10000)
	duration := time.Duration(period)
	time.Sleep(duration * time.Millisecond)

	fmt.Println("Completed cooking ", food, "after miliseconds: ", duration)
}
