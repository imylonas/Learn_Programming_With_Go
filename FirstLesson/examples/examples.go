package examples

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func FirstSol(){
	fmt.Println("Hello World")

}

func SecondSol(){
	fmt.Println("Ioannis Mylonas")
}

func ThirdSol(){
	fmt.Println("My name is Ioannis Mylonas \nMy age is 36")
}

func FourSol(){

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	  }

}

func FiftSol(x,y float64) float64 {
	result := math.Pow(x,y)
	return result

}

func SixSol(){

	for i := 1; i<=1000; i++{
		fmt.Println(i,"\n")
	}

}

func SevenSol(x int) {

	fmt.Println(rand.Intn(x))
}

func EightSol(){
	fmt.Println(time.Now())
}