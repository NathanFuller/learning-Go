package main
import (
	"fmt"
	"time"
)

func main(){
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func(){
		for j := range jobs{
			fmt.Println("recieved job ", j)
		}
		fmt.Println("received all jobs")
		done <- true
	}()

	for j := 1; j<=3; j++{
		jobs <- j
		fmt.Println("sent job ", j)
		time.Sleep(time.Second )
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}
