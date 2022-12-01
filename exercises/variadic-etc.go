package exercises

import "fmt"

func Variadic() {
	fmt.Println()
	fmt.Println("one")
	fmt.Println("one", "two")
	fmt.Println("one", "two", "three")

	sayHello()
	sayHello("Sammy")
	sayHello("Sammy", "Jessica", "Drew", "Jamie")

	var line string

	line = join(",", "Sammy", "Jessica", "Drew", "Jamie")
	fmt.Println(line)

	line = join(",", "Sammy", "Jessica")
	fmt.Println(line)

	line = join(",", "Sammy")
	fmt.Println(line)

	names := []string{"Sammy", "Jessica", "Drew", "Jamie"}

	// we can explode a slice by suffixing it with a set of ellipses (...) and turning it into discrete arguments that will be passed to a variadic function
	line = join(",", names...)
	fmt.Println(line)
}

func sayHello(names ...string) {
	if len(names) == 0 {
		fmt.Println("nobody to greet")
		return
	}
	for _, n := range names {
		fmt.Printf("Hello %s\n", n)
	}
}

// You can only have one variadic parameter in a function, and it must be the last parameter defined in the function. Defining parameters in a variadic function in any order other than the last parameter will result in a compilation error:

func join(del string, values ...string) string {
	var line string
	for i, v := range values {
		line = line + v
		if i != len(values)-1 {
			line = line + del
		}
	}
	return line
}
