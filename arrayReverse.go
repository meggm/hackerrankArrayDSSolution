package main

import "fmt"

func main(){
  a:=[]int{5,3,1,6,9,2}
  revArr:=reverseArray(a)
  fmt.Println(revArr)
}

func reverseArray(a []int) []int {
 //creating the empty slice to add the array elements in reverse order
 revarr:=[]int{}
 
 for _,v:=range a{
     //Appending the elemnt to the beginning of the slice with each iteration
     revarr=append([]int{v},revarr...)
 }
    
    return revarr

}
