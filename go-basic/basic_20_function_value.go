package main

func sayGoodBye() {
	println("Good Bye")
}

func main() {
	goodBye := sayGoodBye

	goodBye()
}
