package main

func main() {
	yo := make(chan string)
	go func() {
		yo <- "Yo"
	}()
	println(<-yo)
}
