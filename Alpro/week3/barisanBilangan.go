package main

import "fmt"
func cetak(n,m int){
	for i := n; i <= m; i++{
		if n < m {
			fmt.Print(i)
			fmt.Print(" ")
		} else {
			return
		}
	} 
}

func main(){
	var a,b int
	fmt.Print("Coba masukan 2 bilangan:")
	fmt.Scan(&a, &b)
	cetak(a, b)
}
