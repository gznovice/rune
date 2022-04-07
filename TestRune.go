package main

import "fmt"

func main(){
	eStr := "Hello World"
	cStr := "你好，世界"
	
	fmt.Println("English")
	
	for i := 0; i < len(eStr); i++{
		fmt.Printf("%c", eStr[i])		
		fmt.Printf("%T", eStr[i])	
	}
	fmt.Println("")
	fmt.Println("len(eStr):", len(eStr))
	fmt.Println("")
	fmt.Println("Chinese:")
	
	for i := 0; i < len(cStr); i++{
		fmt.Printf("%c", cStr[i])		
		fmt.Printf("%T", cStr[i])	
	}
	fmt.Println("")
	fmt.Println("len(cStr):", len(cStr))
	fmt.Println("")
	
	
	for _, v :=  range eStr{
		fmt.Printf("%c", v)		
		fmt.Printf("%T", v)
	}
	
	fmt.Println("")
	
	for _, v :=  range cStr{
		fmt.Printf("%c", v)		
		fmt.Printf("%T", v)		
	}
}