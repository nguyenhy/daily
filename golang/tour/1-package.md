- [Source](#source)
- [Example](#example)
- [Naming convention](#naming-convention)
- [Package paths](#package-paths)
- [Bad package names](#bad-package-names)

### Source
- [Package names](https://go.dev/blog/package-names)

### Example
```go
package main

// By convention, the package name is the same as the last element of the import path
import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
```

### Naming convention
1. Lower case, short and clear
2. Use simple noun
3. If abbreviating a package name makes it ambiguous or unclear, don’t do it.
4. Avoid repetition
	- If the `HTTP` server provided by the `http` package is called `Server`, not `HTTPServer`
	- Client code will refer to this type as `http.Server`
5. Simplify function name
	- When a function in package pkg returns a value of type `pkg.Pkg` (or `*pkg.Pkg`), the function name can often omit the type name without confusion:

	```go
	start := time.Now()                                  // start is a time.Time
	t, err := time.Parse(time.Kitchen, "6:06PM")         // t is a time.Time
	ctx = context.WithTimeout(ctx, 10*time.Millisecond)  // ctx is a context.Context
	ip, ok := userip.FromContext(ctx)                    // ip is a net.IP
	```

	- A function named New in package pkg returns a value of type `pkg.Pkg`

	```go
	 q := list.New()  // q is a *list.List
	```

	- When a function returns a value of type `pkg.T`, where `T` is not `Pkg`, the function name may include `T` to make client code easier to understand

	```go
	d, err := time.ParseDuration("10s")  // d is a time.Duration
	elapsed := time.Since(start)         // elapsed is a time.Duration
	ticker := time.NewTicker(d)          // ticker is a *time.Ticker
	timer := time.NewTimer(d)            // timer is a *time.Timer
	```

### Package paths
### Bad package names
