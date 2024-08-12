package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"strconv"
)

func main(){
	fi, err := os.Open("input.txt");
	
	if err != nil {
		fmt.Println("you forgot ze input!");
		return;
	}
	defer fi.Close();
	
	reader := bufio.NewReader(fi);
	
	left := -1;
	right := -1;
	total := 0;

	for {
		line, err := reader.ReadString('\n');

		if err != nil {
			fmt.Println("EOF")
			break;
		}
		dupe := line;
		for len(dupe) > 0 && left == -1 {
			switch {
				case dupe[0] >= '0' && dupe[0] <= '9':
					left = int(dupe[0] - '0');
				case strings.HasPrefix(dupe, "zero"):
					left = 0;
				case strings.HasPrefix(dupe, "one"):
					left = 1;
				case strings.HasPrefix(dupe, "two"):
					left = 2;
				case strings.HasPrefix(dupe, "three"):
					left = 3;
				case strings.HasPrefix(dupe, "four"):
					left = 4;
				case strings.HasPrefix(dupe, "five"):
					left = 5;
				case strings.HasPrefix(dupe, "six"):
					left = 6;
				case strings.HasPrefix(dupe, "seven"):
					left = 7;
				case strings.HasPrefix(dupe, "eight"):
					left = 8;
				case strings.HasPrefix(dupe, "nine"):
					left = 9;
				default:
					dupe = dupe[1:];
			}
		}

		dupe = line;
		
		for len(dupe) > 0 && right == -1 {
			switch {
				case dupe[len(dupe) - 1] >= '0' && dupe[len(dupe) - 1] <= '9':
					right = int(dupe[len(dupe) - 1] - '0');
				case strings.HasSuffix(dupe, "zero"):
					right = 0;
				case strings.HasSuffix(dupe, "one"):
					right = 1;
				case strings.HasSuffix(dupe, "two"):
					right = 2;
				case strings.HasSuffix(dupe, "three"):
					right = 3;
				case strings.HasSuffix(dupe, "four"):
					right = 4;
				case strings.HasSuffix(dupe, "five"):
					right = 5;
				case strings.HasSuffix(dupe, "six"):
					right = 6;
				case strings.HasSuffix(dupe, "seven"):
					right = 7;
				case strings.HasSuffix(dupe, "eight"):
					right = 8;
				case strings.HasSuffix(dupe, "nine"):
					right = 9;
				default:
					dupe = dupe[:(len(dupe) - 1)];
			}
		}
		
		// Because \n treated as line for some reason, so i have to resort it this way
		if left != -1 && right != -1 { 
			concat := strconv.Itoa(left) + strconv.Itoa(right);
			tempSum,_ := strconv.Atoi(concat);
			total += tempSum;
			left,right = -1, -1;
		}
	}
	fmt.Println(total);
}
