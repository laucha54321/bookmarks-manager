package fnbookmark

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/net/html"
)

type bookmark struct {
	URL   string
	title string
}

type bookmarkList struct {
	title string
	list  []bookmark
}

// parsing the html file
func ListFunction() {

	var list []bookmark

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
			if token.Data == "h3" {
				aux := bookmark{URL: "ajsdklf"}
				tokenType = doc.Next()
				aux.title = doc.Token().Data
				list = append(list, aux)
			} else if token.Data == "a" {
				aux := bookmark{URL: token.Attr[0].Val}
				tokenType = doc.Next()
				aux.title = doc.Token().Data
				list = append(list, aux)
			}

		}

	}
	List(list)
}

// List the bookmarks to the console(after parsing).
func List(aux []bookmark) {
	const bm = `
	██████╗ ██████╗ ███╗   ███╗
	██╔══██╗██╔══██╗████╗ ████║
	██████╔╝██████╔╝██╔████╔██║
	██╔══██╗██╔══██╗██║╚██╔╝██║
	██████╔╝██████╔╝██║ ╚═╝ ██║
	╚═════╝ ╚═════╝ ╚═╝     ╚═╝
							   
	`

	color.Blue(bm)

	for i := 0; i < len(aux); i++ {
		println()
		color.Blue("  " + aux[i].title)
		color.White("  " + aux[i].URL)
	}
}
