package todo


import (
	"os"

	"fmt"

	"sort"

	"github.com/spf13/cobra"

	"io/ioutil"

	"encoding/json"
)


type Item struct {
	Text string 
	Priority int
	Done bool
	position int
}

// TodoCmd represents the `todo` command
var TodoCmd = &cobra.Command{
	Use: "todo",
	Short: "Manage your todos",
	Run: runTodo,
}

// runTodo handles the execution logic
func runTodo(cmd *cobra.Command, args[] string) {
	items := []Item{}
	for _, x := range args {
		items = append(items, Item{Text: x})
	}
	fmt.Println("Todo List: ")
	for _, item := range items {
		fmt.Println("-", item.Text)
	}
}


// SaveItems handles the execution logic for saving content on storage disk
func SaveItems(filename string, items []Item) error {

	b, err := json.Marshal(items)

	if err != nil {
		return err
	}

	/*fmt.Println(string(b))

	return nil*/

	return ioutil.WriteFile(filename, b, 0644)
}

// ReadItems reads data from storage disk
/*
func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Item{}, nil
	}
	var items []Item

	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}
	return items, nil
}
*/

func ReadItems(filename string) ([]Item, error) {
    // Check if file exists
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        // Create empty file with empty array JSON content
        err = ioutil.WriteFile(filename, []byte("[]"), 0644)
        if err != nil {
            return nil, err
        }
        return []Item{}, nil
    }

    b, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    var items []Item
    if err := json.Unmarshal(b, &items); err != nil {
        return nil, err
    }
    return items, nil
}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

func (i *Item) PrettyDone() string {
	if i.Done {
		return "X"
	}
	return ""
}

// Sort todos alphabetically by text
type ByText []Item
func (a ByText) Len() int           { return len(a) }
func (a ByText) Swap(i,j int)       { a[i], a[j] = a[j], a[i] }
func (a ByText) Less(i,j int) bool  { return a[i].Text < a[j].Text }

// ByDone sorts todos with completed items last
type ByDone []Item 
func (a ByDone) Len() int           { return len(a) }
func (a ByDone) Swap(i,j int)       { a[i], a[j] = a[j], a[i] }
func (a ByDone) Less(i,j int) bool  {
    return !a[i].Done && a[j].Done
}

// Sort todos ByPriority
type ByPriority []Item
func (a ByPriority) Len() int       { return len(a) }
func (a ByPriority) Swap(i,j int)   { a[i], a[j] = a[j], a[i] }
func (a ByPriority) Less(i,j int) bool {
	if a[i].Priority == a[j].Priority {
		return a[i].Text < a[j].Text // Fallback to alphabetical if equal
	}
	return a[i].Priority < a[j].Priority
}


// PrintTodos prints a nicely formatted and sorted todo list
func PrintTodos(items []Item) {
	// Sort items by priority
	sort.Sort(ByPriority(items))

	fmt.Println("Todos list:")
	for i, item := range items {
		status := " "
		if item.Done {
			status = "x"
		}

		fmt.Printf("%d. [%s] %s (Priority %d)\n", i+1, status, item.Text, item.Priority)

	}
}

