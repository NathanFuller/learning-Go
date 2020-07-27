package main
import(
	"fmt"
	"regexp"
	"aoc/reading"
)


func main(){
	/* Get the map and regex ready */
	orbits := make(map[string]string)
	re, _ := regexp.Compile(`([^)]+)\)(.+)`)

	/*Read in from file and map bodies to the body they orbit */
	lines := reading.FileToLines("test.txt")
	//lines := reading.FileToLines("six.in")

	for _, line :=range lines{
		mat := re.FindStringSubmatch(line)
		orbits[mat[2]] = mat[1]
	}

	meSet := make(map[string]int)
	santaSet := make(map[string]int)

	transfers := 0

	var comAns string

	p1 := orbits["YOU"]
	p2 := orbits["SAN"]
	for {
		meSet[p1] = transfers
		santaSet[p2] = transfers
		if _, mine := meSet[p2]; mine{
			comAns = p2
			break
		}
		if _, his := santaSet[p1]; his {
			comAns = p1
			break
		}
		p1 = orbits[p1]
		p2 = orbits[p2]
		transfers++
	}

	fmt.Println(comAns, "is a common ancestor")
	fmt.Println("My set: ", meSet)
	fmt.Println("Santa's set: ", santaSet)
	fmt.Println("Distance: ", meSet[comAns]+santaSet[comAns])

}
