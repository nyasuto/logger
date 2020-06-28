# what
A simple log util with automated level

# Use
```go
package main
import (
	"errors"
	"github.com/nyasuto/logger/log"
)

func main() {

	log.Logger("%v", "Hello world")
	// 2019/10/19 17:58:20 [INFO] Hello world

	log.Logger("%v", errors.New("I am error"))
	// 2019/10/19 17:58:20 [ERROR] I am error

}
```

