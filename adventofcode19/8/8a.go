package main
import(
	"fmt"
	//"aoc/reading"
	"io/ioutil"
)


func main(){
	const width = 25
	const height = 6

	//pic := getPicture("test.txt")
	pic := getPicture("eight.in")
	numLayers := len(pic) / (width * height)
	layers := make([][]uint8, numLayers)

	index := 0
	for i := 0; i < numLayers; i++{
		layers[i] = make([]uint8, width * height)
		for j := 0; j < (width * height); j++{
			layers[i][j] = pic[index] - 48
			index++
		}
	}


	leastZ := 1000
	leastI := -1
	for i, arr := range layers{
		numZ := countInArray(arr, 0)
		if numZ < leastZ{
			leastZ = numZ
			leastI = i
		}
	}

	fmt.Printf("Layer %v has %v zeroes.\n", leastI, leastZ)
	fmt.Printf("#1s * #2s = %v\n", countInArray(layers[leastI], 1) * countInArray(layers[leastI], 2))
}

func countInArray(array []uint8, val uint8) int {
	count := 0
	for _, v := range array{
		if v == val{
			count++
		}
	}
	return count
}

func getPicture(fname string) []uint8{
	dat, err := ioutil.ReadFile(fname)
	if err != nil{
		panic(err)
	}
	return dat
}
