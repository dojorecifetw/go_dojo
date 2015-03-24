package main

import "testing"

func Test1Is1(t *testing.T){
    if(fizzBuzzer(1) != "1"){
        t.Error("1 should be 1")
    }
}

func Test2Is2(t *testing.T){
	if(fizzBuzzer(2) != "2"){
        t.Error("2 should be 2")
    }
}

func Test3IsFizz(t *testing.T) {
	if(fizzBuzzer(3) != "Fizz") {
		t.Error("3 should be 'Fizz'")
	}
}

func Test5IsBuzz(t *testing.T) {
	if (fizzBuzzer(5) != "Buzz") {
		t.Error("5 should be 'Buzz'")
	}
}

func Test6IsFizz(t *testing.T) {
	if (fizzBuzzer(6) != "Fizz") {
		t.Error("6 should be 'Fizz'")
	}
}

func Test10IsBuzz(t *testing.T) {
	if (fizzBuzzer(10) != "Buzz") {
		t.Error("10 should be 'Buzz'")
	}
}

func Test15IsFizzBuzz(t *testing.T) {
	if (fizzBuzzer(15) != "FizzBuzz") {
		t.Error("15 should be 'FizzBuzz'")
	}
}

func TestProduce15DynamicFizzBuzzNumbers(t *testing.T) {
	var a []string

	a = fizzBuzz(15)

	if (len(a) != 15) {
		t.Error("Should have 15 elements.")
	}

	a = fizzBuzz(20)

	if (len(a) != 20) {
		t.Error("Should have 20 elements.")
	}
}

func TestFizzBuzzNumbersShouldContainFizzBuzzSequence(t *testing.T) {
	var a []string

	a = fizzBuzz(15)

	if (a[0] != "1") {
		t.Error("List Should contain 1.")
	}

	if (a[1] != "2") {
		t.Error("List Should contain 2.")
	}
}


