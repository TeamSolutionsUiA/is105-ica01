package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/TeamSolutionsUiA/is105-ica02/sum/sum"
)

func main() {
	v1 := os.Args
	i1, err := strconv.ParseInt(v1[1], 10, 64)
	isFloat := false
	if err != nil {
		//printe ut om out of range
		// parse string og isFloat = true om syntax feil
		log.Fatal(err)
	}
	i2, err := strconv.ParseInt(v1[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	if isFloat{
		tot := sum.SumFloat64(f1, f2)
		fmt.Println(tot)
	} else{
		tot := sum.SumInt64(i1, i2)
		fmt.Println(tot)
	}
}
