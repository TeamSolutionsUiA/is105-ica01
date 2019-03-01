package main

import (
"fmt"	
"strconv"
"os";
"github.com/TeamSolutionsUiA/ica01/is105-ica02/sum/sum";	
)


func main() {
v1 :=os.Args;
v2, err :=strconv.ParseInt(v1[1],10,64);
v3, err :=strconv.ParseInt(v1[2],10,64);


tot :=sum.SumInt64(v2, v3)
if err==nil {
fmt.Println(tot);

}else {
fmt.Println(err);	
}



}

