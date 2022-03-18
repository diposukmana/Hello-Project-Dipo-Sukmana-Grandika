package main
import "fmt"

func main(){
	var x, y, i, bit int

	fmt.Scan(&x)
	bit = 0
	i = 1
	for x != 0 {
		y = x % 2
		bit = (y * i) + bit
		x = x / 2
		i = i * 10
	}
	fmt.Print(bit)
}