package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func checkDupe(signalTxt *string) bool {
    hasDupe := false
    for j, char := range *signalTxt {
        currentTxt := string(char)
        for k := j + 1; k < len(*signalTxt); k++ {
            nextTxt := string((*signalTxt)[k])

            if currentTxt == nextTxt {
                hasDupe = true
                break
            }
        }
    }

    return hasDupe
}

func checkDupeSet(signalTxtStr *string, setLen int) int {
    finalPos := 0

    for i, _ := range *signalTxtStr {
        nextX := (*signalTxtStr)[i:i+setLen]

        if (i + setLen) >= len(*signalTxtStr) {
            break
        }

        fmt.Println(nextX)
        hasDupe := checkDupe(&nextX)

        if !hasDupe {
            finalPos = i + setLen
            break
        }
    }

    return finalPos
}

func solution1(scanner bufio.Scanner) {
    for scanner.Scan() {
        signalTxt := scanner.Text()

        var signalTxtStr string = string(signalTxt)

        finalPos := checkDupeSet(&signalTxtStr, 4)
        fmt.Println("Final position: ", finalPos)
    }
}

func solution2(scanner bufio.Scanner) {
    for scanner.Scan() {
        signalTxt := scanner.Text()

        var signalTxtStr string = string(signalTxt)

        finalPos := checkDupeSet(&signalTxtStr, 14)
        fmt.Println("Final position: ", finalPos)
    }
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