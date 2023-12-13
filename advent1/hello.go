package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "sort"
    "strconv"
)

func main() {
    f, err := os.Open("./input.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    var caloriesSum []int
    loopCalories := 0

    for scanner.Scan() {
        if scanner.Text() == "" {
            caloriesSum = append(caloriesSum, loopCalories)
            loopCalories = 0
        } else {
            i, err := strconv.Atoi(scanner.Text())
            if err != nil {
                continue
            }
            loopCalories += i
        }
    }

    sort.Slice(caloriesSum, func(i, j int) bool {
        return caloriesSum[i] > caloriesSum[j]
    })

    topThreeCalories := 0
    for i := 0; i < 3; i++ {
        topThreeCalories += caloriesSum[i]
    }

    fmt.Println(topThreeCalories)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
