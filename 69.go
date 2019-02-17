package main
import "fmt"
var a int =10

func main(){
	//a := 10
	add1 := add()
    fmt.Println(add1())
}

func add() func () int{
	a = 20-10
	return func()int{
		return a
	}
}