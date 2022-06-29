package main

import (
	"fmt"
	"math"
	"strconv"
)
func main() {
	strbit := "000"
	count := ""
	for strbit != "0000000000000000000000"{
		count += AllBits(strbit)
		strbit = AddOnebit(strbit)
	}
	fmt.Print(count)
	
}
func convBit(bitByte []byte, div int, pos int, iteration int)([]byte){
	if (iteration) % div == 0 {
		if bitByte[pos] == 48 {
			bitByte[pos] = 49
		}else {
			bitByte[pos] = 48
		}
	}
	return bitByte
}
func check3bit(bitByte []byte, count int) (int){
	bitStr := string(bitByte)
	for i:=0; i < len(bitByte) - 2; i++ {
		other := i + 3
		if (bitStr[i:other] == "000") {
			count ++;
			break
		}
	}
	return count
}
func AllBits(bitStr string) (string){
	count := 0
	bitByte := []byte(bitStr)
	total := len(bitStr)

	num := int(math.Pow(2, float64(len(bitStr))) + 1)
	for i:=1;i<num;i++ {
		count =  check3bit(bitByte, count)
		temp := total
		for j := 0; j < total; j ++ {
			bitByte = convBit(bitByte, int(math.Pow(2, float64(temp - 1))),total - temp, i)
			temp -= 1
		}
		
	}
	return strconv.Itoa(int(math.Pow(2, float64(len(bitStr)))) - count) + ","
	
}

func AddOnebit(bitStr string) string {
	count := len(bitStr)
	bitStr = ""
	for i := 0; i < count+1; i++ {
		bitStr += "0"
	}
	return bitStr
}