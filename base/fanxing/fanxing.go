package main

import "log"

func main() {

}

func sum[T any](a T) T {

	return a
}

func DemoCompare[A comparable, B any](a A, b B) {

	log.Printf("a=%+v, b=%+v\b", a, b)
}
