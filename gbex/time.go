package main

import(
	"fmt"
	"time"
)

func main(){
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(1993, 10, 19, 15, 22, 0,0, time.UTC)
	p(then)
	p(now.Year())
	p(then.Nanosecond())
	p(then.Weekday())
	p(now.Weekday())
	
	diff := now.Sub(then)
	p(diff)
	p(diff.Hours())

}
