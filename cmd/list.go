/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
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

	doc := html.NewTokenizer(strings.NewReader(string(file)))
	if err != nil {
		fmt.Println(err.Error())
	}

	for {
		//get the next token type
		tokenType := doc.Next()

		//if it's an error token, we either reached
		//the end of the file, or the HTML was malformed
		if tokenType == html.ErrorToken {
			err := doc.Err()
			if err == io.EOF {
				//end of the file, break out of the loop
				break
			}
			//otherwise, there was an error tokenizing,
			//which likely means the HTML was malformed.
			//since this is a simple command-line utility,
			//we can just use log.Fatalf() to report the error
			//and exit the process with a non-zero status code
			log.Fatalf("error tokenizing HTML: %v", doc.Err())
		}

		//process the token according to the token type...
		if tokenType == html.StartTagToken {
			token := doc.Token()
			if token.Data == "a" {
				fmt.Println(token.Attr[0].Val)
				tokenType = doc.Next()
				fmt.Println(doc.Token())
			}
		}

	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
