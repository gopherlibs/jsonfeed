# JSON Feed Go Library [![Go Reference](https://pkg.go.dev/badge/github.com/gopherlibs/jsonfeed.svg)](https://pkg.go.dev/github.com/gopherlibs/jsonfeed) [![CI Status](https://circleci.com/gh/gopherlibs/jsonfeed.svg?style=shield)](https://app.circleci.com/pipelines/github/gopherlibs/jsonfeed) [![Software License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/gopherlibs/jsonfeed/master/LICENSE)

This is a library to parse and validate a JSON Feed.
JSON Feed is a modern alternative to RSS and Atom.
I'd recommend learning more about it at the [official website][jfeed] if you're unfamiliar.

**Note: This package is pre-1.0 thus the API is still changing as I prepare it for a v1.0 release.**


## Usage

As common for "Gopher Libs" Go modules, importing this package requires writing the package name twice in the import statement.

An `io.Reader` is required in order to `Parse` a JSON Feed into a `Feed` struct.
From there, you can `Validate` that struct to make sure it follows JSON Feed specs.

```go
import(
	"fmt"
	"os"

	"github.com/gopherlibs/jsonfeed/jsonfeed"
)

func main(){
	file, _ := os.Open("file-containing-a-json-feed.json")

	feed, err := Parse(file)
	if err != nil {
		println("Failed to parse feed.")
	}

	if errs := feed.Validate(); len(errs) == 0{
		fmt.Println("The feed validated perfectly.")
	}else{
		fmt.Println("There's one or more validation errors with the feed.")
	}
}
```


## Development

This library is written and tested with Go v1.15+ in mind.
`go fmt` is your friend.
Please feel free to open Issues and PRs as you see fit.
Any PR that requires a good amount of work or is a significant change, it would be best to open an Issue to discuss the change first.


## License & Credits

This module was written by Ricardo N Feliciano (FelicianoTech).
This repository is licensed under the MIT license.
This repo's license can be found [here](./LICENSE).



[jfeed]: https://www.jsonfeed.org/
