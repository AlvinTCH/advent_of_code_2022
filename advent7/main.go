package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func getFileDirSizeMap(scanner bufio.Scanner) map[string]int {
    fileDirSize := make(map[string]int)

    var parentDir []string

    for scanner.Scan() {
        cmdTxt := scanner.Text()

        isGoPath := strings.HasPrefix(cmdTxt, "$ cd ")
        isLs := strings.HasPrefix(cmdTxt, "$ ls")
        isDir := strings.HasPrefix(cmdTxt, "dir ")

        if isLs {
            continue
        }

        if isDir {
            continue
        }

        if isGoPath {
            // go back to home, clear parent dir
            if cmdTxt == "$ cd /" {
                parentDir = nil
                parentDir = append(parentDir, "/")
                continue
            }

            if cmdTxt == "$ cd .." {
                parentDir = parentDir[:len(parentDir)-1]
                continue
            }
            
            replacemenStr := strings.ReplaceAll(cmdTxt, "$ cd ", "")
            if (len(parentDir) == 1) {
                parentDir = append(parentDir, replacemenStr)
                continue
            }
            
            parentDir = append(parentDir, "/" + strings.ReplaceAll(cmdTxt, "$ cd ", ""))
            continue
        }

        cmdTxtSplit := strings.Split(cmdTxt, " ")
        intVar, err := strconv.Atoi(cmdTxtSplit[0])

        if err != nil {
            log.Fatal(err)
        }

        pathBuild := ""
        for _, fileDir := range parentDir {
            fileDirSize[pathBuild + fileDir] += intVar
            pathBuild += fileDir
        }
    }

    return fileDirSize
}

func solution1(scanner bufio.Scanner) {
    fileDirSize := getFileDirSizeMap(scanner)

    totalSize := 0
    for _, element := range fileDirSize {
        if element <= 100000 {
            totalSize += element
        }
    }

    fmt.Println("totalSize: ", fileDirSize)
}

func solution2(scanner bufio.Scanner) {
    fileDirSize := getFileDirSizeMap(scanner)
    
    fileSizeToFree := 30000000 - (70000000 - fileDirSize["/"])
    
    leastFileSize := 0
    for _, element := range fileDirSize {
        if element < fileSizeToFree {
            continue
        }

        if leastFileSize == 0 {
            leastFileSize = element
            continue
        }

        if element < leastFileSize {
            leastFileSize = element
        }
    }

    fmt.Println("least file size: ", leastFileSize)
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