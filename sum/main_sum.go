package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/TeamSolutionsUiA/is105-ica01/sum/sum"
)

func main() {
	args := os.Args
	isFloat := false
	if len(args) != 3{
		fmt.Println("Bruk to argumenter")
		os.Exit(2)
	}
	i1, err := strconv.ParseInt(args[1],10, 64)
	if err != nil {
		isFloat = true
	}
	i2, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		isFloat = true
	}
	if i1 > 0 {
		if i1 > math.MaxInt64-i2 {
			isFloat = true
		}
	} else {
		if i1 < math.MinInt64-i2 {
			isFloat = true
		}
	}
	if isFloat{
		f1, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			log.Fatal(err)
	}
		f2, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		tot := sum.SumFloat64(f1, f2)
		fmt.Println(tot)
	} else{
		tot := sum.SumInt64(i1, i2)
		fmt.Println(tot)
	}
}