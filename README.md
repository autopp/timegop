# Timegop

`timegop` is a Go library to fix/disguise current time.  
This is useful for tests involving time.

`timegop` is inspired by Ruby library `timecop`.

## Installation

```
go get github.com/autopp/timegop
```

## Usage

### Basic

```go
package main

import (
	"fmt"
	"time"
	"github.com/autopp/timegop"
)

func main() {
	fmt.Println(timegop.Now()) // => current time
	// go to unix epoc
	timegop.Freeze(time.Unix(0, 0))
	fmt.Println(timegop.Now()) // => epoc time

	// back to the past
	past, _ := time.Parse("2006/01/02 15:04", "1985/10/26 01:20")
	timegop.Travel(past)
	time.Sleep(time.Minute)
	fmt.Println(timegop.Now()) // => 1985/10/26 01:21

	// Return to real
	timegop.Return()
	fmt.Println(timegop.Now()) // => current time
}
```

### Returning to real time

After calling `Freeze`/`Travel`, you must call `Return` to returning:

```go
timegop.Freeze(t)
Work()
timegop.Return()
```

By using `defer`, you can make sure you do not forget to call:

```go
timegop.Freeze(t)
defer timegop.Return()
Work()
```

`Freeze`/`Travel` returns `Return`, so you can also write bellow:

```go
defer timegop.Freeze(t)()
Work()
```

## License

[Apache License 2.0](LICENSE)

## Author

[AuToPP](https://twitter.com/AuToPP)
