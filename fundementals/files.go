package fundementals

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func FilesMain() {
	files, err := os.ReadDir("data")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mode\t\tSize\t\tName\t\tDate")
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("->", file.Name())
			continue
		}
		info, err := file.Info()
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("[File] %s, modified %v, %v(bytes) \n", info.Name(), info.ModTime().Format(time.RFC822), info.Size())
		fmt.Printf("%s\t%d\t\t%s\t\t%s\n", info.Mode(), info.Size(), info.Name(), info.ModTime().Format(time.RFC822))
	}

	info, err := os.Stat("data/a.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mode\t\tSize\tName")
	fmt.Printf("%s\t%d\t%s\n", info.Mode(), info.Size(), info.Name())

	// Walk root directory
	err = filepath.WalkDir("fundementals/testdata", func(path string, d fs.DirEntry, err error) error {
		// if there was an error, return it
		// if there is an error, it is most likely
		// because an error was encountered trying
		// to read the top level directory
		if err != nil {
			return err
		}
		// if the file is a directory
		// return nil to tell walk to continue
		// walking the directory,
		// but to no longer continue
		// operating on the directory itself
		if d.IsDir() {
			return nil
		}
		// get the file info for the file
		info, err := d.Info()
		if err != nil {
			return err
		}
		// if the file is not a directory
		// then print its mode, size, and path
		fmt.Printf("%s\t%d\t%s\n", info.Mode(), info.Size(), path)
		// return nil to tell walk to continue
		return nil

	})
	if err != nil {
		log.Fatal(err)
	}
}

func Walk() ([]string, error) {
	var entries []string
	err := filepath.WalkDir("testdata", func(path string, d fs.DirEntry, err error) error {
		// if there was an error, return it
		// if there is an error, it is most likely
		// because an error was encountered trying
		// to read the top level directory
		if err != nil {
			return err
		}
		// if the entry is a directory, handle it
		if d.IsDir() {
			// name of the file or directory
			name := d.Name()
			// if the directory is a dot return nil
			// this may be the root directory
			if name == "." || name == ".." {
				return nil
			}
			// if the directory name is "testdata"
			// or it starts with "."
			// or it starts with "_"
			// then return filepath.SkipDir
			if strings.HasPrefix(name, ".") || strings.HasPrefix(name, "_") {
				return fs.SkipDir
			}
			return nil
		}
		// append the entry to the list
		entries = append(entries, path)
		// return nil to tell walk to continue
		return nil
	})
	return entries, err
}

func WalkFS(cab fs.FS) ([]string, error) {
	var entries []string
	err := fs.WalkDir(cab, ".", func(path string, d fs.DirEntry, err error) error {
		// if there was an error, return it
		// if there is an error, it is most likely
		// because an error was encountered trying
		// to read the top level directory
		if err != nil {
			return err
		}
		// if the entry is a directory, handle it
		if d.IsDir() {
			// name of the file or directory
			name := d.Name()
			// if the directory is a dot return nil
			// this may be the root directory
			if name == "." || name == ".." {
				return nil
			}
			// if the directory name is "testdata"
			// or it starts with "."
			// or it starts with "_"
			// then return filepath.SkipDir
			if strings.HasPrefix(name, ".") || strings.HasPrefix(name, "_") {
				return fs.SkipDir
			}
			return nil
		}
		// append the entry to the list
		entries = append(entries, path)
		// return nil to tell walk to continue
		return nil
	})
	return entries, err
}

func Create(name string, body []byte) error {
	// create a new file, this will
	// truncate the file if it already exists
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	// write the body to the file
	_, err = f.Write(body)
	return err
}

func Append(name string, body []byte) error {
	// if the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	// write the body to the file
	_, err = f.Write(body)
	return err
}

func Read(fp string, w io.Writer) error {
	f, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(w, f)
	return err
}
