//This is the first part solved (02a). Part B is in another file.

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

		count := strings.Count(password, symbol)

		if int64(count) >= numMinInt && int64(count) <= numMaxInt {
			fmt.Println(j, "Password is valid!: ", currentLine)
			validPasswords += 1
		} else {
			fmt.Println(j, "Invalid password...")
		}

	}
	println("Valid passwords: ", validPasswords)
}
