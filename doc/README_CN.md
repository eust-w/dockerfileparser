[🇺🇸](../README.md)[English](../README.md)

# dockerfileparser
dockerfileparser can be used to parse dockerfile

# 如何使用
`go get github.com/eust-w/dockerfileparser`


# 例子
## 通过文件解析
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
## 通过io.Reader解析
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