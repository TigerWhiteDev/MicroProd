package main

import (
	"fmt"
	"io/ioutil"
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
			install()
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

	var c cproject
	if args != nil && len(args) > 0 {

		for _, arg := range args {
			gendprojet(c, arg)
		}
		/*var name string
		if args[0] == "module" {
			fmt.Println("create module on project")

		}*/
	} else {

		files, _ := ioutil.ReadDir("settings/project/")
		for _, f := range files {
			name := f.Name()
			gendprojet(c, name)
		}
	}

}
func gendprojet(c cproject, name string) {

	c.getConf("settings/project/" + name + "/" + name + ".yml")
	if c.Type == "service" {
		fmt.Println("create service")
		os.Mkdir("storage/"+name, 0777)
		os.Mkdir("project/backoffice/"+name, 0777)
		os.Chdir("project/backoffice/" + name)
		os.Symlink("../../../settings/project/"+name, "settings")
		os.Symlink("../../../storage/"+name, "storage")
	} else if c.Type == "application" {
		fmt.Println("create application")
		os.Mkdir("storage/"+name, 0777)
		os.Mkdir("project/front/"+name, 0777)
		os.Chdir("project/front/" + name)
		os.Symlink("../../../settings/project/"+name, "settings")
		os.Symlink("../../../storage/"+name, "storage")
	} else {
		fmt.Println("generate type config is not valide")
	}
	genvendor(c.Lang)
	os.Chdir("../../..")
}

func genvendor(lang string) {
	switch lang {
	case "nodejs":
		if _, err := os.Stat("../../../vendor/nodejs"); os.IsNotExist(err) {
			os.Mkdir("../../../vendor/nodejs", 0777)
		}
		os.Symlink("../../../vendor/nodejs", "node_modules")
	case "php":
		if _, err := os.Stat("../../../vendor/php"); os.IsNotExist(err) {
			os.Mkdir("../../../vendor/php", 0777)
		}
		os.Symlink("../../../vendor/php", "vendor")

	default:
		fmt.Println("not language supported")
	}
}

func install() {

	fmt.Println("install")
	var folder string
	folder = "storage"
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.Mkdir(folder, 0777)
	}
	folder = "vendor"
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.Mkdir(folder, 0777)
	}

	generate(nil)

}
