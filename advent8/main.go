package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func parseMatrix(scanner bufio.Scanner) [][]int {
    treesMatrix := make([][]int, 0)
    for scanner.Scan() {
        treeRow := scanner.Text()
        treeRowSplit := strings.Split(treeRow, "")
        convertedTreeRow := make([]int, 0)
        for _, v := range treeRowSplit {
            intVar, err := strconv.Atoi(v)

            if err != nil {
                log.Fatal(err)
            }

            convertedTreeRow = append(convertedTreeRow, intVar)
        }
        treesMatrix = append(treesMatrix, convertedTreeRow)
    }

    return treesMatrix
}

func solution1(scanner bufio.Scanner) {
    treesMatrix := parseMatrix(scanner)

    var treesCount int

    for treeRowIndex, treeRow := range treesMatrix {
        if (treeRowIndex == 0) || (treeRowIndex == len(treesMatrix) - 1) {
            treesCount += len(treeRow)
            continue
        }

        for treeColIndex, treeCol := range treeRow {
            if (treeColIndex == 0) || (treeColIndex == len(treeRow) - 1) {
                treesCount++
                continue
            }

            // visible from left
            leftColSlice := treeRow[0:treeColIndex]
            isLeftVisible := true
            for i := 0; i < len(leftColSlice); i++ {
                if (leftColSlice[i] >=  treeCol) {
                    isLeftVisible = false
                    break
                }
            }
            if (isLeftVisible) {
                treesCount++
                continue
            }

            // visible from right
            rightColSlice := treeRow[treeColIndex+1:len(treeRow)]
            isRightVisible := true
            for i := 0; i < len(rightColSlice); i++ {
                if (rightColSlice[i] >=  treeCol) {
                    isRightVisible = false
                    break
                }
            }
            if (isRightVisible) {
                treesCount++
                continue
            }

            // visible from top
            isTopVisible := true
            for i := 0; i < treeRowIndex; i++ {
                if (treesMatrix[i][treeColIndex] >=  treeCol) {
                    isTopVisible = false
                    break
                }
            }
            if (isTopVisible) {
                treesCount++
                continue
            }

            // visible from bottom
            isBottomVisible := true
            for i := treeRowIndex + 1; i < len(treeRow); i++ {
                if (treesMatrix[i][treeColIndex] >=  treeCol) {
                    isBottomVisible = false
                    break
                }
            }
            if (isBottomVisible) {
                treesCount++
                continue
            }
        }
    }

    fmt.Println(treesCount)
}

func solution2(scanner bufio.Scanner) {
    treesMatrix := parseMatrix(scanner)

    biggestTreeScore := 0
    for treeRowIndex, treeRow := range treesMatrix {
        if (treeRowIndex == 0) || (treeRowIndex == len(treesMatrix) - 1) {
            continue
        }

        for treeColIndex, treeCol := range treeRow {
            if (treeColIndex == 0) || (treeColIndex == len(treeRow) - 1) {
                continue
            }

            // visible from left
            leftColSlice := treeRow[0:treeColIndex]
            leftScore := 0
            for i := len(leftColSlice) - 1; i >= 0; i-- {
                leftScore++
                if (leftColSlice[i] >=  treeCol) {
                    break
                }
            }

            // visible from right
            rightColSlice := treeRow[treeColIndex+1:len(treeRow)]
            rightScore := 0
            for i := 0; i < len(rightColSlice); i++ {
                rightScore++
                if (rightColSlice[i] >= treeCol) {
                    break
                }
            }

            // visible from top
            topScore := 0
            for i := treeRowIndex - 1; i >= 0; i-- {
                topScore++
                fmt.Println("top: ", i)
                if (treesMatrix[i][treeColIndex] >=  treeCol) {
                    break
                }
            }

            // visible from bottom
            bottomScore := 0
            for i := treeRowIndex + 1; i < len(treeRow); i++ {
                bottomScore++
                if (treesMatrix[i][treeColIndex] >=  treeCol) {
                    break
                }
            }

            treeScore := leftScore * rightScore * topScore * bottomScore
            if (treeScore > biggestTreeScore) {
                biggestTreeScore = treeScore
            }
        }
    }

    fmt.Println(biggestTreeScore)
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