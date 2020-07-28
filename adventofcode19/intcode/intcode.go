package main
import (
	"fmt"
	"aoc/reading"
	"strconv"
	"sync"
	"sync/atomic"
)

var amp uint64
var wg sync.WaitGroup

func main(){

	settings := make(chan []int, 120)

	Perm([]int{5,6,7,8,9}, func(array []int, channel chan []int){
		fmt.Println("in Perm: ", array)
		channel <- array
	},
	settings)
	close(settings)

	max := 0
	for st := range settings{
		num := runSeries(st)
		if num > max {
			max = num
		}
	}
	fmt.Println("Max thrust:", max)

}

func runSeries(stArr []int) int{
	fmt.Println(stArr)

	atob := make(chan int, 1)
	btoc := make(chan int, 1)
	ctod := make(chan int, 1)
	dtoe := make(chan int, 1)
	etoa := make(chan int, 2)

	etoa <-stArr[0]
	atob <-stArr[1]
	btoc <-stArr[2]
	ctod <-stArr[3]
	dtoe <-stArr[4]

	etoa <- 0

	wg.Add(1)
	go amplifier(etoa, atob) //amp A
	wg.Add(1)
	go amplifier(atob, btoc) //amp B
	wg.Add(1)
	go amplifier(btoc, ctod) //amp C
	wg.Add(1)
	go amplifier(ctod, dtoe) //amp D
	wg.Add(1)
	go amplifier(dtoe, etoa) //amp E

	wg.Wait()
	return <-etoa
}

func amplifier(input chan int, output chan int) {
	defer wg.Done()
	atomic.AddUint64(&amp, 1)
	runprog(input, output)
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

func getModes(instruction int) [3]bool{
	first := ((instruction - (instruction % 100)) - (instruction - (instruction % 1000))) / 100
	second := ((instruction - (instruction % 1000)) - (instruction - (instruction % 10000))) / 1000
	third := (instruction - (instruction % 10000)) / 10000
	return [3]bool{first == 1, second == 1, third == 1}
}


/*From https://yourbasic.org/golang/generate-permutation-slice-string/  */
/* Perm calls f with each permutation of a. */
func Perm(a []int, f func([]int, chan []int), ch chan []int) {
    perm(a, f, 0, ch)
}

/* Permute the values at index i to len(a)-1. */
func perm(a []int, f func([]int, chan []int), i int, ch chan []int) {
    if i > len(a) {
	    s := make([]int, len(a))
	    copy(s, a)
        f(s, ch)
        return
    }
    perm(a, f, i+1, ch)
    for j := i + 1; j < len(a); j++ {
        a[i], a[j] = a[j], a[i]
        perm(a, f, i+1, ch)
        a[i], a[j] = a[j], a[i]
    }
}
