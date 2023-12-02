package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	f, err := os.Open("index")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var total int
	for scanner.Scan() {

		r := scanner.Text()

		chars := strings.Split(r, "")
		fv := getFirstNumber(chars)
		lv := getLastNumber(chars)

		fmt.Printf("%s = first: %d, last: %d \n", r, fv, lv)

		realOne := fmt.Sprintf("%d%d", fv, lv)

		n, _ := strconv.Atoi(realOne)

		total += n
	}

	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getFirstNumber(c []string) int {
	for i := 0; i < len(c); i++ {
		n, err := strconv.Atoi(c[i])
		if err != nil {
			continue
		}
		return n
	}
	return 0

}

func getLastNumber(c []string) int {

	for i := len(c) - 1; i >= 0; i-- {
		n, err := strconv.Atoi(c[i])
		if err != nil {
			continue
		}
		return n
	}
	return 0
}
