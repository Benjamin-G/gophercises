package exercises

import (
	"flag"
	"fmt"
	"os"
)

var addFlag *bool
var doFlag *int
var listFlag *bool

// Available Commands:
//   add         Add a new task to your TODO list
//   do          Mark a task on your TODO list as complete
//   list        List all of your incomplete tasks

func init() {
	addFlag = flag.Bool("add", false, "Add a new task to your TODO list")
	doFlag = flag.Int("do", -1, "Mark a task on your TODO list as complete")
	listFlag = flag.Bool("list", false, "List all of your incomplete tasks")

	flag.Usage = func() {
		fmt.Printf("Usage of %s \n", os.Args[0])
		fmt.Println("This app accepts ONLY a single flag")
		flag.PrintDefaults()
	}

	flag.Parse()
}

func TODO_APP() {
	switch {
	case flag.NFlag() == 0:
		fmt.Println("Please specify single flag")
		flag.Usage()
		return
	case flag.NFlag() > 1:
		fmt.Println("Please specify ONLY a single flag")
		return
	}

	switch {
	case *addFlag:
		fmt.Println("Run Add!")
	case flag.NFlag() > 1:
		fmt.Println("Run Do!")
	case *listFlag:
		fmt.Println("Run List!")
	}
	fmt.Println("Run Application!")
}
