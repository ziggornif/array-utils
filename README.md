# Go array utils

## Description

This package provides useful array operation functions such as can be found in the javascript and java languages.

The array is encapsulated in a `Array` struct for chaining methods and doing "functional" programming.

## Installation and usage

To install it, run:

    go get github.com/drouian-m/array-utils

## Example

```go
myarr := Array[int]{[]int{1, 2, 3, 4, 5}}
result := myarr.Map(func(val int) int {
  return val * 3
}).Filter(func(val int) bool {
  return val % 2 == 0
})
fmt.println(result.Arr) // => [6 12]
```

## Available operations

- Find : Find a value in the current array
- FindIndex : Find value index in the current array
- Filter : Filter array values (chainable)
- Map : Create a new array populated with the result of calling the input func on each array element (chainable)
