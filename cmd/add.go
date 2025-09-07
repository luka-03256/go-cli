/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/luka-03256/go-cli/todo"

)


// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long: `Add will create a new todo item to the list`,
	Run: addRun,
}


// add Command function
/*
func addRun(cmd *cobra.Command, args []string) {
	for _, x := range args {
		fmt.Println(x)
	}
}
*/

// Prev version only prepends new data
/*func addRun(cmd *cobra.Command, args []string) {
	var items = []todo.Item{}
	for _, x := range args {
		items = append(items, todo.Item{Text:x})
	}
	err := todo.SaveItems("todos.json", items)
	//fmt.Println("%#v\n", items)
	if err != nil {
		fmt.Println("Error saving todos:", err)
	}
}*/

func addRun(cmd *cobra.Command, args []string) {
    if len(args) == 0 {
        fmt.Println("Please provide a todo item.")
        return
    }

    //Read existing todos
    items, err := todo.ReadItems("todos.json")
    if err != nil {
        fmt.Printf("Failed to read todos: %v\n", err)
        return
    }

    //Append new todos
    for _, text := range args {
        items = append(items, todo.Item{Text: text})
    }

    // Save the updated list
    err = todo.SaveItems("todos.json", items)
    if err != nil {
        fmt.Printf("Failed to save todos: %v\n", err)
        return
    }

    fmt.Println("Added new todo(s).")
}

/*
func SaveItems(filename string, items []Item) error {
	b := items
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	return nil
}
*/

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
