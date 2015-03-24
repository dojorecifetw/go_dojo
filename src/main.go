package main

import ("fmt"
		"strconv")


func fizzBuzzer(n int) string {
	var t string


	if ((n % 5 == 0) &&  
		(n % 3 == 0)) {
		t = "FizzBuzz"
	} else if (n % 3 == 0) {
		t = "Fizz"
	} else if (n % 5 == 0){
		t = "Buzz"
	} else {
		t = strconv.Itoa(n)	
	}
	
	return t
}

func fizzBuzz(size int) []string {
	var a []string
	for i := 0; i < size; i++ {
		a = append(a, fizzBuzzer(i+1))
	}
	return a
}

func main(){
    fmt.Printf("It's a beginning.")
}