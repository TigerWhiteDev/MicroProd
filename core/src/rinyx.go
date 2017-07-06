package main

import (
	"fmt"
	"os"
)

func main() {

	/*argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[1]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
	fmt.Println("hh")*/

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			fmt.Println("install option")
		case "generate":
			generate(os.Args[2:])
		case "build":
			fmt.Println("build all services")
		case "start":
			fmt.Println("start serveur dev")
		case "stop":
			fmt.Println("stop serveur dev")

		case "deploy":
			fmt.Println("deploy project")
		default:
			fmt.Println(os.Args[1] + " commande not know")
		}

	} else {
		fmt.Println("You must specify one of the command")
	}

}

func generate(args []string) {
	fmt.Println(args)
}
