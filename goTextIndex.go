package main
import
(
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)
func main() {
	alphabetTable := make(map[byte]int)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter file name:")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input = strings.Replace(input, "\n", "", -1)
	stream, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}
	fromFile := string(stream)
	for i := 0; i < len(fromFile); i++ {
		_, exists := alphabetTable[fromFile[i]]
		if exists {
			alphabetTable[fromFile[i]] = alphabetTable[fromFile[i]] + 1
		} else {
			alphabetTable[fromFile[i]]++
		}
	}
	PrintTable(alphabetTable)
}
func PrintTable(table map[byte]int) {
	for letter, counts := range table {
		fmt.Printf("%c -> %d\n", letter, counts)
	}
}
