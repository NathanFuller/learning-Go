package main
import "fmt"

func main(){
	queue := make(chan string)
	queue <- "one"
	queue <- "two"

	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}

