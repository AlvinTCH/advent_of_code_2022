package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

func parseStacks(scanner bufio.Scanner) map[string][]string {
    stackArr := make(map[string][]string)
    for scanner.Scan() {
        stackTxt := scanner.Text()
        stackTxtSplit := strings.Split(stackTxt, "-")
        stackArr[stackTxtSplit[0]] = strings.Split(stackTxtSplit[1], ",")
    }

    return stackArr
}

func parseMovement(movementTxt string) (int, string, string) {
    movementTxtReplace := strings.ReplaceAll(movementTxt, "move ", "")
    movementTxtSplit := strings.Split(movementTxtReplace,  " from ")

    originDstSplit := strings.Split(movementTxtSplit[1], " to ")

    movementAmt, err := strconv.Atoi(movementTxtSplit[0])
    if err != nil {
        panic(err)
    }

    return movementAmt, originDstSplit[0], originDstSplit[1]
}

func moveDataOneByOne(stackArr map[string][]string, moveNum int, origin string, dst string) {
    originArray := stackArr[origin]
    dstArray := stackArr[dst]

    slicedArray := originArray[len(originArray)-moveNum:len(originArray)]
    for i := len(slicedArray) - 1; i >= 0; i-- {
        dstArray = append(dstArray, slicedArray[i])
    }
    originArray = originArray[:len(originArray) - moveNum]

    stackArr[origin] = originArray
    stackArr[dst] = dstArray
}


func solution1(scanner bufio.Scanner, stackArr map[string][]string) {
    for scanner.Scan() {
        movementTxt := scanner.Text()
        num, origin, dst := parseMovement(movementTxt)
        
        moveDataOneByOne(stackArr, num, origin, dst)
    }

    fmt.Println(stackArr)
}

func moveBulkData(stackArr map[string][]string, moveNum int, origin string, dst string) {
    fmt.Println(stackArr, moveNum, origin, dst)
    originArray := stackArr[origin]
    dstArray := stackArr[dst]

    dstArray = append(dstArray, originArray[len(originArray)-moveNum:len(originArray)]...)
    originArray = originArray[:len(originArray) - moveNum]

    stackArr[origin] = originArray
    stackArr[dst] = dstArray
    fmt.Println(stackArr)
}

func solution2(scanner bufio.Scanner, stackArr map[string][]string) {
    for scanner.Scan() {
        movementTxt := scanner.Text()
        num, origin, dst := parseMovement(movementTxt)
        
        moveBulkData(stackArr, num, origin, dst)
    }
}

func main() {

    stack, stackErr := os.Open("./stack.txt")

    if stackErr != nil {
        log.Fatal(stackErr)
    }

    stackScanner := bufio.NewScanner(stack)
    stackArr := parseStacks(*stackScanner)

    stack.Close()

    f, err := os.Open("./input.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    // solution1(*scanner, stackArr)
    solution2(*scanner, stackArr)
}