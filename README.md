# Go array utils

## Description

This package provides useful array operations functions such as can be found in the javascript and java languages.

The native array is encapsulated in a `Array` struct for chaining methods and doing "functional" programming.

## Requirements

Go >= 1.18 version.

## Installation

To install it, run:

    go get github.com/drouian-m/array-utils

## Usage

### Create a new Array

```go
package main
import (
	"github.com/drouian-m/array-utils"
)

func main() {
	arr := array.NewArray([]int{1, 2, 3})
}
```

### Available operations

- Find : Find a value in the current array
- FindIndex : Find value index in the current array
- Filter : Filter array values (chainable)
- Map : Create a new array populated with the result of calling the input func on each array element (chainable)

### Retrieve native Go array

```go
package main
import (
	"github.com/drouian-m/array-utils"
)

func main() {
	arr := array.NewArray([]int{1, 2, 3})
	arr.ToNative() // => return []int array
}
```

### Chaining example

```go
package main
import (
	"fmt"
	"github.com/drouian-m/array-utils"
)

func main() {
	myarr := array.NewArray([]int{1, 2, 3, 4, 5})
	result := myarr.Map(func(val int) int {
		return val * 3
	}).Filter(func(val int) bool {
		return val % 2 == 0
	})
	fmt.Println(result.ToNative()) // => [6 12]
}
```

