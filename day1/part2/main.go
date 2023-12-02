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

		l := scanner.Text()

		lr := replaceLine(l)

		fmt.Printf("actual: %s, replace: %s \n", l, lr)

		chars := strings.Split(lr, "")

		fv := getFirstNumber(chars)
		lv := getLastNumber(chars)

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

func replaceLine(line string) string {
	//one3
	replacers := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	for i := range replacers {
		line = strings.ReplaceAll(line, i, replacers[i])
	}
	return line
}
