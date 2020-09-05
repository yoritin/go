package main

func main() {
	msg := "Hello"
	func() {
		println(msg)
	}()
}