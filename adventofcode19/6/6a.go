package main
import(
	"fmt"
	"regexp"
	"aoc/reading"
)


func main(){
	/* Get the map and regex ready */
	myMap := make(map[string]string)
	re, _ := regexp.Compile(`([^)]+)\)(.+)`)

	/*Read in from file and map bodies to the body they orbit */
	//lines := reading.FileToLines("test.txt")
	lines := reading.FileToLines("six.in")

	for _, line :=range lines{
		mat := re.FindStringSubmatch(line)
		myMap[mat[2]] = mat[1]
	}

	/*Get workers counting the orbital relationships */
	const numWorkers = 1
	bodies := make(chan string)
	orbits := make(chan int)

	for w := 0; w < numWorkers; w++{
		go worker(myMap, bodies, orbits)
	}

	go func(){
		for job := range myMap{
			bodies <- job
		}
		close(bodies)
	}()

	/*Get the answers back and add them up */
	var numOrbits int = 0
	for i := 1; i <= len(myMap); i++{
		num := <-orbits
		numOrbits += num
	}

	/*Report total*/
	fmt.Println("Total: ", numOrbits, "orbits.")

}

func worker(myMap map[string]string, bodies <-chan string, results chan<-int){
	for body := range bodies{
		count := 0
		for curr := body; curr != "COM"; curr = myMap[curr]{
			count++
		}
		results <- count
	}

}
