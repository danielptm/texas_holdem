package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PokerIO struct {
	reader ReadPokerinput
}

type ReadPokerinput interface {
	getInput() (string, error)
}

type ScanImpl struct{}

func (ScanImpl) getInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	res, err := reader.ReadString('\n')
	return res, err
}

func (g PokerIO) GetChips() (int, error) {
	fmt.Print("Enter bet (max 10): ")
	// ReadString will block until the delimiter is entered
	input, err := g.reader.getInput()
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return -1, err
	}
	n, err := strconv.ParseInt(strings.Trim(input, "\n"), 10, 64)

	returnVal := int(n)

	if err != nil || returnVal > 10 {
		fmt.Println("An error occured while parsing the integer. Please pass a number less than 10", err)
		return -1, err
	}

	return returnVal, nil
}
