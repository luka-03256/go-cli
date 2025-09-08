/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"log"

	"sort"

	"os"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"

	"github.com/luka-03256/go-cli/todo"

	"text/tabwriter"
)

var sortBy string
var sortDesc bool
var doneOpt bool
var allOpt bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the todos",
	Long:  `List will display all todos, optionally sorted by criteria like priority.`,
	Run:   listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("dataFile"))
	if err != nil {
		log.Fatalf("Error reading todos: %v", err)
	}

	// conditionall sort if --sort=priority is passed
	/*if sortBy == "priority" {
		sort.Sort(todo.ByPriority(items))
	}*/

	// conditionall sort by all todos column names 'Priority','Done','Text'
	switch sortBy {
	case "priority":
		sort.Sort(todo.ByPriority(items))
	case "text":
		sort.Sort(todo.ByText(items))
	case "done":
		sort.Sort(todo.ByDone(items))
	case "":
		// No sorting
	default:
		fmt.Printf("Unknown sort ption: %s\n", sortBy)
		return
	}

	// if --desc is set, reverse the list
	if sortDesc {
		for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
			items[i], items[j] = items[j], items[i]
		}
	}

	// Tabwriter for formatting output
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, "No.\tStatus\tPriority\tTask")

	fmt.Fprintln(w, "No.\tStatus\tPriority\tTask")

	count := 0
	for _, item := range items {
		if allOpt || item.Done == doneOpt {
			count++
			status := " "
			if item.Done {
				status = "x"
			}

			var priorityLabel string
			switch item.Priority {
			case 1:
				priorityLabel = "High"
			case 2:
				priorityLabel = "Medium"
			case 3:
				priorityLabel = "Low"
			default:
				priorityLabel = fmt.Sprintf("%d", item.Priority)
			}

			fmt.Fprintf(w, "%d.\t[%s]\t%s\t%s\n", count, status, priorityLabel, item.Text)
		}
	}

	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")
	listCmd.Flags().StringVarP(&sortBy, "sort", "s", "", "Sort by: priority, text, or done")
	listCmd.Flags().BoolVar(&sortDesc, "desc", false, "Sort in descending order")
	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show only done todos")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show all todos")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
