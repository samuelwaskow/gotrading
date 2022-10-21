package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type State struct {
	open1  float32
	close1 float32
	low1   float32
	high1  float32
	ppen2  float32
	close2 float32
	low2   float32
	high2  float32
	ma20   float32
}

/**
 * Main
 */
func main() {
	records := readCSVFile("WIN.csv")

	total := len(records)
	for row := 1; row < total; row++ {
		transformState(records[row])
	}

}

/**
 * Transform the state prices to a more readable format
 */
func transformState(rawstate []string) {

	open1, _ := strconv.ParseFloat(rawstate[1], 8)
	close1, _ := strconv.ParseFloat(rawstate[2], 8)
	low1, _ := strconv.ParseFloat(rawstate[3], 8)
	high1, _ := strconv.ParseFloat(rawstate[4], 8)

	open2, _ := strconv.ParseFloat(rawstate[5], 8)
	close2, _ := strconv.ParseFloat(rawstate[6], 8)
	low2, _ := strconv.ParseFloat(rawstate[7], 8)
	high2, _ := strconv.ParseFloat(rawstate[8], 8)

	ma20, _ := strconv.ParseFloat(rawstate[9], 8)

	hilo1 := high1 - low1
	hilo2 := high2 - low2

	newOpen1 := (open1 - low1) / hilo1
	newClose1 := (close1 - low1) / hilo1

	newOpen2 := (open2 - low2) / hilo2
	newClose2 := (close2 - low2) / hilo2

	newMA := (0.5 * ma20) / close2

	fmt.Println(newOpen1, newClose1, newOpen2, newClose2, newMA)

}

/**
 *
 */
func readCSVFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

/**
 * Function for handling messages
 **/
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

/**
 * Function for handling errors
 **/
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
