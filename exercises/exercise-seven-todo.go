package exercises

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/boltdb/bolt"
)

var addFlag *bool
var doFlag *int
var listFlag *bool

// Available Commands:
//   add         Add a new task to your TODO list
//   do          Mark a task on your TODO list as complete
//   list        List all of your incomplete tasks

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

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

func cleanTextTODO(s string) string {
	s = strings.TrimSuffix(s, "\n")
	s = strings.TrimSuffix(s, "\r")
	s = strings.TrimSpace(s)
	return s
}

func addTODO(db *bolt.DB) {
	fmt.Println("Type text to add to your TODOs, then press enter:")
	reader := bufio.NewReader(os.Stdin)
	todoText, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		return
	}

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todos"))

		id, _ := b.NextSequence()

		fmt.Printf("Key :%v", int(id))

		buf, err := json.Marshal(cleanTextTODO(todoText))
		if err != nil {
			return err
		}

		return b.Put(itob(int(id)), buf)
	})
}

func completeTask(db *bolt.DB) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todos"))

		v := b.Get(itob(*doFlag))
		if v == nil {
			log.Fatal("TODO not found")
			return nil
		}

		var data string
		json.Unmarshal(v, &data)
		fmt.Printf("Completed: %s", data)

		return b.Delete(itob(*doFlag))
	})
}

func printTODOs(db *bolt.DB) {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todos"))

		c := b.Cursor()

		if k, _ := c.First(); k == nil {
			fmt.Println("You have completed all TODO.")
			return nil
		}

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var data string
			json.Unmarshal(v, &data)
			fmt.Printf("key:%v, list-item: %s\n", binary.BigEndian.Uint64(k), data)
		}

		return nil
	})
}

// go get github.com/boltdb/bolt/...

func TODO_APP() {
	var app func(db *bolt.DB)

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

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("todos"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	app(db)
}
