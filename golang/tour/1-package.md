- [Source](#source)
- [Example](#example)
- [Naming convention](#naming-convention)
- [Package paths](#package-paths)
- [Bad package names](#bad-package-names)
  - [Avoid meaningless package names](#avoid-meaningless-package-names)
  - [Break up generic packages](#break-up-generic-packages)
  - [Don’t use a single package for all your APIs](#dont-use-a-single-package-for-all-your-apis)
  - [Avoid unnecessary package name collisions](#avoid-unnecessary-package-name-collisions)

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
#### Avoid meaningless package names
Packages named util, common, or misc provide clients with no sense of what the package contain.
#### Break up generic packages
- raw code
```go
// package
package util
func NewStringSet(...string) map[string]bool {...}
func SortStringSet(map[string]bool) []string {...}

// client
set := util.NewStringSet("c", "a", "b")
fmt.Println(util.SortStringSet(set))
```

- choose the right name
```go
// package
package stringset
func New(...string) map[string]bool {...}
func Sort(map[string]bool) []string {...}

// client
set := stringset.New("c", "a", "b")
fmt.Println(stringset.Sort(set))
```

- improve the package
```go
// package
package stringset
type Set map[string]bool
func New(...string) Set {...}
func (s Set) Sort() []string {...}

// client
set := stringset.New("c", "a", "b")
fmt.Println(set.Sort())
```

#### Don’t use a single package for all your APIs
package named `api`, `types`, `interfaces`, `util`, `common` will not providing any guidance to users. Instead break them up, perhaps using directories to separate public packages from implementation.

#### Avoid unnecessary package name collisions
While packages in different directories may have the same name, packages that are frequently used together should have distinct names. This reduces confusion and the need for local renaming in client code.

For the same reason, avoid using the same name as popular standard packages like `io` or `http`.

