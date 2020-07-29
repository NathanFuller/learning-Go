package main
import(
	"fmt"
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

	view := make([]uint8, width * height)
	for i := 0; i < width * height; i++{
		for j := 0; j < numLayers; j++ {
			if layers[j][i] != 2 {
				view[i] = layers[j][i]
				break
			}
		}
	}

	idx := 0
	for y:=0; y < height; y++{
		for x := 0; x < width; x++{
			if view[idx] == 0{
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
			idx++
		}
		fmt.Print("\n")
	}
}

func getPicture(fname string) []uint8{
	dat, err := ioutil.ReadFile(fname)
	if err != nil{
		panic(err)
	}
	return dat
}
