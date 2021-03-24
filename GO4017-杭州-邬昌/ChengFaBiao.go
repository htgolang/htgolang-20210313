package main

import "fmt"
func main(){
   
   //使用for循环打印乘法口诀
   for i:=1;i<10;i++{
      for j:=1;j<=i;j++{


         fmt.Print(i,"*",j,"=",i*j," ")

      }
     fmt.Println(" ")
  }

}
