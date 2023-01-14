package main
import "fmt"
func main() {
	i:=1
	for i<=3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("New loop")
	for n:=0; n<=10; n++ {
		if n % 2 == 0 {
			continue
		}
		fmt.Println(n)
	}
} //End of program
