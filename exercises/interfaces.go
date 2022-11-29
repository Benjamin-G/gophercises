package exercises

import "fmt"

// Note - The above code snippet is an example of favouring composition over inheritance and is typically the approach you want to take for all components within your Go systems!

type Employee interface {
	GetName() string
	PrintDetails()
}

type Engineer struct {
	Name string
}

func (e *Engineer) GetName() string {
	return "Engineer Name: " + e.Name
}

func (e Engineer) PrintDetails() {
	fmt.Println(e.GetName())
}

type Manager struct {
	Name string
}

func (m *Manager) GetName() string {
	return "Manager Name: " + m.Name
}

func (e Manager) PrintDetails() {
	fmt.Println(e.GetName())
}

func PrintDetails(e Employee) {
	fmt.Println(e.GetName())
}

type Counter interface {
	addWithPointer()
	addWithValue() Click
}

type Click struct {
	clicked    bool
	numClicked int
}

func (c *Click) addWithPointer() {
	c.clicked = true
	c.numClicked++
}

func (c Click) addWithValue() Click {
	return Click{
		clicked:    true,
		numClicked: c.numClicked + 1,
	}
}

func Interfaces() {
	engineer := Engineer{Name: "Elliot"}
	manager := Manager{Name: "Donna"}
	PrintDetails(&engineer)
	PrintDetails(&manager)
	engineer.PrintDetails()
	manager.PrintDetails()

	newCounter := Click{clicked: false, numClicked: 0}
	fmt.Println(newCounter)
	newCounter.addWithPointer()
	fmt.Println(newCounter)
	newCounter = newCounter.addWithValue()
	fmt.Println(newCounter)
}

// When working with concurrent code, there are a few different options for safe operation. We’ve gone over two of them:

// Synchronization primitives for sharing memory (e.g., sync.Mutex)

// Synchronization via communicating (e.g., channels)

// However, there are a couple of other options that are implicitly safe within multiple concurrent processes:

// Immutable data

// Data protected by confinement
