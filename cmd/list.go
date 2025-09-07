/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/luka-03256/go-cli/todo"

	"log"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: listRun,
}


func listRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems("todos.json")
	if err != nil {
		log.Printf("%v", err)
	}

	fmt.Println("Todos list:")
	for i, item := range items {
		status := " "
		if item.Done {
			status = "x"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, item.Text)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
