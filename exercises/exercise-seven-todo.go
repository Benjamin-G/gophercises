package exercises

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
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

func addTODO() {
	fmt.Println("Run Add!")
}

func completeTask() {
	fmt.Println("Run Do! ", *doFlag)
}

func printTODOs() {
	fmt.Println("Run List!")
}

// go get github.com/boltdb/bolt/...

func TODO_APP() {
	var app func()

	switch {
	case flag.NFlag() > 1:
		fmt.Println("Please specify ONLY a single flag")
		return
	case *addFlag:
		app = addTODO
	case *listFlag:
		app = printTODOs
	case *doFlag != -1:
		app = completeTask
	default:
		fmt.Println("Please specify single flag")
		flag.Usage()
		return
	}

	db, err := bolt.Open("todo.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app()
}
