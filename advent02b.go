package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	var linesArray []string
	validPasswords := 0
	i := 0

	fmt.Println("Scanning the file")
	for fileScanner.Scan() {

		linesArray = append(linesArray, fileScanner.Text())
		i++
	}

	for j := range linesArray {
		currentLine := linesArray[j]
		positionSeparator1 := strings.IndexRune(currentLine, '-')
		positionSeparator2 := strings.IndexRune(currentLine, ' ')
		positionSeparator3 := strings.IndexRune(currentLine, ':')

		numMin := currentLine[0:positionSeparator1]
		numMax := currentLine[positionSeparator1+1 : positionSeparator2]
		symbol := currentLine[positionSeparator2+1 : positionSeparator3]
		password := currentLine[positionSeparator3+2:]

		numMinInt, _ := strconv.ParseInt(numMin, 10, 64)
		numMaxInt, _ := strconv.ParseInt(numMax, 10, 64)

		//count := strings.Count(password, symbol)

		/*fmt.Print(j, " ")
		fmt.Print("Symbol: ", symbol)
		fmt.Print(" First passwd symb: " , password[numMinInt-1:numMinInt])
		fmt.Print(" Second passwd symb: ",password[numMaxInt-1:numMaxInt])
		fmt.Print(" Password: ", password)
		fmt.Println()
*/


		if strings.Contains(password[numMinInt-1:numMinInt], symbol) && !strings.Contains(password[numMaxInt-1:numMaxInt], symbol) || !strings.Contains(password[numMinInt-1:numMinInt], symbol) && strings.Contains(password[numMaxInt-1:numMaxInt], symbol) {
			validPasswords += 1
			fmt.Println(j, "Found a valid password^^^^ ")
		}

	}
	println("Valid passwords: ", validPasswords)
}
