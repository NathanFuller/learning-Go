package main
import (
	"fmt"
	"aoc/reading"
	"strconv"
	"sync/atomic"
)

var series uint64

func main(){
	settings := make(chan []int, 200)

	Perm([]int{5,6,7,8,9}, func(array []int, channel chan []int){
		channel <- array
	},
	settings)

	outs := make([]int, 0)
	for i := 1; i <= 120; i++{
		outs = append(outs, <-signals)
	}
	max := 0
	for _, s := range outs{
		if s > max {
			max = s
		}
	}

	fmt.Printf("Ran %v series\n", series)
	fmt.Printf("Max: %v\n", max)
}

func runSeries(settings []int) int {
	atomic.AddUint64(&series, 1)
	in := 0
	for _, s := range settings{
		in = amplifier(s, in)
	}
	return in
}

func amplifier(setting, input int) int{
	into := make(chan int, 2)
	outof := make(chan int, 1)
	into <- setting
	into <- input
	runprog(into, outof)
	return <-outof
}

func getModes(instruction int) [3]bool{
	first := ((instruction - (instruction % 100)) - (instruction - (instruction % 1000))) / 100
	second := ((instruction - (instruction % 1000)) - (instruction - (instruction % 10000))) / 1000
	third := (instruction - (instruction % 10000)) / 10000
	return [3]bool{first == 1, second == 1, third == 1}
}


func runprog(input <-chan int, output chan<- int){
	var program []int

	records := reading.ReadCSV("7.in")
	//records := reading.ReadCSV("test.txt")
	for _, line := range records{
		for _, rec := range line{
			num, _ := strconv.Atoi(rec)
			program = append(program, num)
		}
	}

	/*Run the intcode program*/
	counter := 0
	for program[counter] != 99{
		switch program[counter] % 100 {
			case 1:
				var sumand0, sumand1 int
				modes := getModes(program[counter])
				if modes[0]{
					sumand0 = program[counter+1]
				} else {
					sumand0 = program[program[counter+1]]
				}
				if modes[1]{
					sumand1 = program[counter+2]
				} else {
					sumand1 = program[program[counter+2]]
				}
				program[program[counter+3]] = sumand0 + sumand1
				counter += 4

			case 2:
				var mul0, mul1 int
				modes := getModes(program[counter])
				if modes[0]{
					mul0 = program[counter+1]
				} else {
					mul0 = program[program[counter+1]]
				}
				if modes[1]{
					mul1 = program[counter+2]
				} else {
					mul1 = program[program[counter+2]]
				}
				program[program[counter+3]] = mul0 * mul1
				counter += 4

			case 3:
				num := <-input
				program[program[counter+1]] = num
				counter += 2

			case 4:
				modes := getModes(program[counter])
				if modes[0] {
					output <- program[counter+1]
				}else {
					output <- program[program[counter+1]]
				}
				counter += 2

			case 5:
				var check, value int
				modes := getModes(program[counter])
				if modes[0] {
					check = program[counter+1]
				} else {
					check = program[program[counter+1]]
				}
				if modes[1] {
					value = program[counter+2]
				} else {
					value = program[program[counter+2]]
				}
				if check != 0 {
					counter = value
				} else {
					counter += 3
				}

			case 6:
				var check, value int
				modes := getModes(program[counter])
				if modes[0] {
					check = program[counter+1]
				} else {
					check = program[program[counter+1]]
				}
				if modes[1] {
					value = program[counter+2]
				} else {
					value = program[program[counter+2]]
				}
				if check == 0 {
					counter = value
				} else {
					counter += 3
				}

			case 7:
				var a, b int
				modes := getModes(program[counter])
				if modes[0] {
					a = program[counter+1]
				} else {
					a = program[program[counter+1]]
				}
				if modes[1] {
					b = program[counter+2]
				} else {
					b = program[program[counter+2]]
				}
				if a < b {
					program[program[counter+3]] = 1
				} else {
					program[program[counter+3]] = 0
				}
				counter += 4

			case 8:
				var a, b int
				modes := getModes(program[counter])
				if modes[0] {
					a = program[counter+1]
				} else {
					a = program[program[counter+1]]
				}
				if modes[1] {
					b = program[counter+2]
				} else {
					b = program[program[counter+2]]
				}
				if a == b {
					program[program[counter+3]] = 1
				} else {
					program[program[counter+3]] = 0
				}
				counter += 4

			default:
				panic(fmt.Sprintf("Invalid opcode: %v at count %v.\nProgram state: %v" ,program[counter], counter, program))
		}
	}
}

/*From https://yourbasic.org/golang/generate-permutation-slice-string/  */
/* Perm calls f with each permutation of a. */
func Perm(a []int, f func([]int, chan int), ch chan int) {
    perm(a, f, 0, ch)
}

/* Permute the values at index i to len(a)-1. */
func perm(a []int, f func([]int, chan int), i int, ch chan int) {
    if i > len(a) {
        f(a, ch)
        return
    }
    perm(a, f, i+1, ch)
    for j := i + 1; j < len(a); j++ {
        a[i], a[j] = a[j], a[i]
        perm(a, f, i+1, ch)
        a[i], a[j] = a[j], a[i]
    }
}
