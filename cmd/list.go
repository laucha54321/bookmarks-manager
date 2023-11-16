/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/net/html"
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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {

	rootCmd.AddCommand(listCmd)
	fmt.Println("List of BookMarks: ")

	file, err := os.ReadFile("./bookmarks.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	doc, err := html.Parse(strings.NewReader(string(file)))
	if err != nil {
		fmt.Println(err.Error())
	}

	var links []string
	var link func(*html.Node)
	link = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					// adds a new link entry when the attribute matches
					links = append(links, a.Val)
				}
			}
		}

		// traverses the HTML of the webpage from the first child node
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			link(c)
		}
	}
	link(doc)

	// loops through the links slice
	for _, l := range links {
		fmt.Println("Link:", l)
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
