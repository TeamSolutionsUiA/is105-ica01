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
	v2, err := strconv.ParseInt(v1[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	v3, err := strconv.ParseInt(v1[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	tot := sum.SumInt64(v2, v3)

	fmt.Println(tot)

}
