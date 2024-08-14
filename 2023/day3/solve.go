package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func getAdjacency(numbers [][]int, symbols [][]int) ([]int, []int) { // {adjacentNumber} {validNumber}
	var adjacentNumber []int;
	var gearRatio []int;
	symbolAdj := make(map[int][]int);

	for _, number := range numbers {
		for _, symbol := range symbols {
			isAdjacent := false;
			
			// Check Horizontal
			if number[0] == symbol[0] {
				if math.Abs(float64(number[2] - symbol[1])) == 1 || math.Abs(float64(symbol[1] - number[3])) == 1 {
					isAdjacent = true;
				}
			}

			// Check Vertical
			if math.Abs(float64(number[0] - symbol[0])) == 1 {
				if symbol[1] >= number[2] && symbol[1] <= number[3] {
					isAdjacent = true;
				}

				// Check Diagonal
				if math.Abs(float64(symbol[1] - number[2])) == 1 || math.Abs(float64(symbol[1] - number[3])) == 1 {
					isAdjacent = true;
				}
			}

			if isAdjacent {
				adjacentNumber = append(adjacentNumber, int(number[1]));
				symbolAdj[symbol[2]] = append(symbolAdj[symbol[2]], number[1]);
			}
		}
	}

	for _, nums := range symbolAdj {
		if len(nums) == 2 {
			gearRatio = append(gearRatio, nums[0]*nums[1]);
		}
	}

	return adjacentNumber, gearRatio;
}

func main(){
	fi, _ := os.Open("./input.txt");
	reader := bufio.NewReader(fi);

	defer fi.Close();
	var total int;
	var lineArr []string;

	for {
		line, err := reader.ReadString('\n');
		line = strings.TrimRight(line,"\r\n");

		if err != nil {
			if err.Error() == "EOF" {
				break;
			}
		}
		if line != "\n"{
			lineArr = append(lineArr, line);
		}
	}

	var symArr [][]int; // []int {line, pos} starts from 0 
	var numArr [][]int; // []int {line, value, starting, ending} starts from 0
	var symId int;
	var numId int;

	for i:=0; i<len(lineArr); i++{
		lineRead := lineArr[i];
		isReading := true;
		pool := "";
		var starting int;
		var ending int;
		
		for pos, char := range lineRead {
			if (char != '.') && !(unicode.IsDigit(char)){ // SYMBOL
				temp := []int {i, pos, -1};
				if char == '*'{
					temp = []int {i, pos, symId};
					symId++;
				}
				symArr = append(symArr, temp);
			}
		}
		// Get Number
		for pos, char := range lineRead {
			if unicode.IsDigit(char) && isReading{
				pool += string(char);
				starting = pos;
				isReading = false;
			}else if isReading == false && unicode.IsDigit(char){
				pool += string(char);
			}else if isReading == false && !unicode.IsDigit(char){
				isReading = true;
				converter,_ := strconv.Atoi(pool);
				temp := []int {i, converter, starting, pos - 1, numId};
				numArr = append(numArr, temp);
				numId++;
				pool = "";
			}
		}
		
		// Absolutely neccessary if the number ends at the end of the line
		if pool != "" {
			ending = len(lineRead) - 1;
			converter, _ := strconv.Atoi(pool);
			temp := []int {i, converter, starting, ending, numId};
			numArr = append(numArr, temp);
			numId++;
		}
	}

	part1, part2 := getAdjacency(numArr, symArr);
	for _, element := range part1 {
		total += element;
	}
	fmt.Println(total);
	total = 0;

	for _, element := range part2 {
		total+= element;
	}
	fmt.Println(total);
}
