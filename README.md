[π¨π³](doc/README_CN.md)[δΈ­ζ](doc/README_CN.md)

# dockerfileparser
parsing dockerfile

# Install
`go get github.com/eust-w/dockerfileparser`


# Usage
## parsing by file
```
import (
	"fmt"
	"github.com/eust-w/dockerfileparser"
)

func main() {
	d, err := dockerfileparser.ParseFile("./dockerfile")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(d.All), d.All)
	fmt.Println(len(d.Run), d.Run)
}
```
## parsing by io.Reader
```
import (
	"fmt"
	"github.com/eust-w/dockerfileparser"
	"os"
)

func main() {
	f, err := os.Open("./dockerfile")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	d, err := dockerfileparser.Parse(f) // handleErr
	if err != nil {
		panic(err)
	}
	fmt.Println(len(d.All), d.All)
	fmt.Println(len(d.Run), d.Run)
}
```