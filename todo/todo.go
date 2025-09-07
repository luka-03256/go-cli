package todo



import (
	"fmt"

	"github.com/spf13/cobra"

	"io/ioutil"

	"encoding/json"

)



type Item struct {
	Text string 
	Done bool
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

func  ReadItems(filename string) ([]Item, error) {
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
