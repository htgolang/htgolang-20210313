package main 
import (
	"fmt"
)
func main(){
	for i:=1;i<10;i++{
		for j:=1;j< i+1;j++{
			fmt.Printf("%v * %v = %v \t",j,i,i*j)			
		}
		fmt.Println()	
	}
}
