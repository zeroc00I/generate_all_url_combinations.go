package main
// make with love by zeroc00i
import (
	"fmt"
	"net/url"
	"strings"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument")
		return
	}
	arg := os.Args[1]
	urlString := arg
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}

	pathParts := strings.Split(parsedUrl.Path, "/")

	for i := range pathParts {
		newPath := strings.Join(pathParts[0:i+1], "/")
		parsedUrl.Path = newPath
		fmt.Println(parsedUrl.String())
	}
}
