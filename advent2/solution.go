package main

import (
    "bufio"
    "fmt"
    "log"
    "strings"
    "os"
)

func main() {
    
    winLoseConditionMap := map[string]string{
        "Rock": "Scissors",
        "Paper": "Rock",
        "Scissors": "Paper",
    }

    opponentTypeMap := map[string]string{
        "A": "Rock",
        "B": "Paper",
        "C": "Scissors",
    }

    selfTypeMap := map[string]string{
        "X": "Lose",
        "Y": "Draw",
        "Z": "Win",
    }

    shapeScoreMap := map[string]int{
        "Rock": 1,
        "Paper": 2,
        "Scissors": 3,
    }

    resultScoreMap := map[string]int{
        "Win": 6,
        "Draw": 3,
        "Lose": 0,
    }

    f, err := os.Open("./input.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    totalScore := 0
    for scanner.Scan() {
        resSplit := strings.Split(scanner.Text(), " ")

    	opponentType := opponentTypeMap[resSplit[0]]

        conditionalResults := selfTypeMap[resSplit[1]]

        var selfType string
        if conditionalResults == "Draw" {
            selfType = opponentType
        } else if conditionalResults == "Win" {
            selfType = winLoseConditionMap[winLoseConditionMap[opponentType]]
        } else {
            selfType = winLoseConditionMap[opponentType]
        }

        totalScore += resultScoreMap[conditionalResults] + shapeScoreMap[selfType]
    }

    fmt.Println(totalScore)
}