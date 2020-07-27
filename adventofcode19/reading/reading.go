package reading
import(
	"bufio"
	"encoding/csv"
	"os"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func FileToLines(fname string) []string {
	f, err := os.Open(fname)
	check(err)
	scanner := bufio.NewScanner(f)

	rec := make([]string, 0)

	for scanner.Scan(){
		rec = append(rec, scanner.Text())
	}
	check(scanner.Err())
	return rec
}

func ReadCSV(fname string) [][]string {
	f, err := os.Open(fname)
	check(err)
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	check(err)
	return records
}
