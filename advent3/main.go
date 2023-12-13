package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func solution1(CapAsciiNum *int, LowerAsciiNum *int, scanner bufio.Scanner) {
    totalScore := 0
    for scanner.Scan() {
        rucksackContains := scanner.Text()
        itemsLen := len(rucksackContains)

        splitStop := itemsLen / 2

        firstHalf := rucksackContains[0:splitStop]
        secondHalf := rucksackContains[splitStop:itemsLen]

        charMap := make(map[string]int)
        for _, char := range firstHalf {
            if !strings.Contains(secondHalf, string(char)) {
                continue
            }
            charAsciiNum := int(char)
            if _, ok := charMap[string(char)]; ok {
                continue
            }
            if (charAsciiNum < *LowerAsciiNum) {
                totalScore += charAsciiNum - *CapAsciiNum + 1 + 26
                charMap[string(char)] = charAsciiNum
                continue
            }
            totalScore += charAsciiNum - *LowerAsciiNum + 1
            charMap[string(char)] = charAsciiNum
        }
    }

    fmt.Println(totalScore)
}

func solution2(CapAsciiNum *int, LowerAsciiNum *int, scanner bufio.Scanner) {
    totalScore := 0
	
	var ruckSackMap []string

    for scanner.Scan() {
        rucksackContains := scanner.Text()
        if len(ruckSackMap) < 2 {
            ruckSackMap = append(ruckSackMap, rucksackContains)
            continue
        }
        
        fmt.Println(rucksackContains, ruckSackMap)
        for _, char := range rucksackContains {
            if !strings.Contains(ruckSackMap[0], string(char)) {
                continue
            }
            if !strings.Contains(ruckSackMap[1], string(char)) {
                continue
            }

            fmt.Println(string(char))

            charAsciiNum := int(char)
            if (charAsciiNum < *LowerAsciiNum) {
                totalScore += charAsciiNum - *CapAsciiNum + 1 + 26
                continue
            }
            totalScore += charAsciiNum - *LowerAsciiNum + 1
            break
        }

        ruckSackMap = nil
    }

    fmt.Println(totalScore)
}

func main() {

    f, err := os.Open("./input.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    startingCapASCIINumber := 65
    startingLowerASCIINumber := 97

    // solution1(&startingCapASCIINumber, &startingLowerASCIINumber, *scanner)
    solution2(&startingCapASCIINumber, &startingLowerASCIINumber, *scanner)
}