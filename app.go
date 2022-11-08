package main
import (
	"fmt"
)
func main(){
	fmt.Println("Qual é o seu nome")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("oi,%s! Eu sou a linguagem Go!",name)
	}
