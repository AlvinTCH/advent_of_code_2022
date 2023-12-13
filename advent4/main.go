package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

func convertToInt(pair []string) []int {
    var convertedPair []int
    for i, _ := range pair {
        convertedNum, err := strconv.Atoi(pair[i])
        if err != nil {
            panic(err)
        }

        convertedPair = append(convertedPair, convertedNum)
    }

    return convertedPair
}

func parseRow(setTxt *string) ([]int, []int) {
    elfSplit := strings.Split(*setTxt, ",")

    firstElfSection := convertToInt(strings.Split(elfSplit[0], "-"))
    secondElfSection := convertToInt(strings.Split(elfSplit[1], "-"))

    return firstElfSection, secondElfSection
}

func checkContain(pair1, pair2 *[]int) bool {
    return ((*pair1)[0] >= (*pair2)[0]) && ((*pair1)[1] <= (*pair2)[1])
}

func solution1(scanner bufio.Scanner) {
    encompassCount := 0
    for scanner.Scan() {
        setTxt := scanner.Text()

        firstElfSection, secondElfSection := parseRow(&setTxt)

        if (checkContain(&firstElfSection, &secondElfSection)) || (checkContain(&secondElfSection, &firstElfSection)) {
            encompassCount++
        }
    }

    fmt.Println(encompassCount)
}


func checkOverlap(pair1, pair2 *[]int) bool {
    if ((*pair1)[0] >= (*pair2)[0]) && ((*pair1)[0] <= (*pair2)[1]) {
        return true
    }

    if ((*pair1)[1] >= (*pair2)[0]) && ((*pair1)[1] <= (*pair2)[1]) {
        return true
    }

    return checkContain(pair1, pair2) || checkContain(pair2, pair1)
}

func solution2(scanner bufio.Scanner) {
    overlapCount := 0
    for scanner.Scan() {
        setTxt := scanner.Text()

        firstElfSection, secondElfSection := parseRow(&setTxt)
        if (checkOverlap(&firstElfSection, &secondElfSection)) {
            overlapCount++
        }
    }

    fmt.Println(overlapCount)
}

func main() {

    f, err := os.Open("./input.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    // solution1(*scanner)
    solution2(*scanner)
}