package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fi, err := os.Open("./input.txt");
	if err != nil {
		print("NOT FOUND");
	}
	reader := bufio.NewReader(fi);
	
	// PART 1
	// MAX_RED := 12;
	// MAX_GREEN := 13;
	// MAX_BLUE := 14;

	total := 0;

	for {
		// PART 2 Comment it for part 1
		MIN_RED := 0;
		MIN_GREEN := 0;
		MIN_BLUE := 0;

		line,err := reader.ReadString('\n');
		line = strings.TrimRight(line, "\r\n");
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("EOF");
				fmt.Println(total);
				return;
			}
		}
		modLine := line[5:];
		gameNum, rawData, _ := strings.Cut(modLine, ": ");
		gameData := strings.Split(rawData, "; "); // {'12 red, 13 blue', '14 green'}
		
		// For part 1 uncomment this
		// valid := true;

		for _, element := range gameData {
			reading := strings.Split(element, ", ");
			for _, element2 := range reading { // { '12 red' , '13 blue' }
				if strings.Contains(element2,"red"){
					numberStr := strings.Split(element2, " red");
					trueNum,_ := strconv.Atoi(numberStr[0]);
					if trueNum > MIN_RED {
						// part 1
						// valid = false;
						// part 2
						MIN_RED = trueNum;
					}
				}else if strings.Contains(element2,"green"){
					numberStr := strings.Split(element2, " green");
					trueNum,_ := strconv.Atoi(numberStr[0]);
					if trueNum > MIN_GREEN {
						// part 1
						// valid = false;
						// part 2
						MIN_GREEN = trueNum;
					}
				}else if strings.Contains(element2, "blue"){
					numberStr := strings.Split(element2, " blue");
					trueNum, _ := strconv.Atoi(numberStr[0]);
					if trueNum > MIN_BLUE {
						// part 1
						// valid = false;
						// part 2
						MIN_BLUE = trueNum;
					}
				}
				// for part 1
				//	if !valid {
				//		break;
				//	}
			}
			// for part 1
			// if !valid{
			//	break
			// }
		}

		// part 1
		//if valid{
			// trueGameNum,_ := strconv.Atoi(gameNum);
			// fmt.Println(trueGameNum);
			// total += trueGameNum;
		//}
		// part 2
		total += MIN_RED * MIN_GREEN * MIN_BLUE;
	}
}
