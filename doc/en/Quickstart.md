# Quick start

## Definition of the types and functions.

Create file named `example.anz` with the contents:

```haskell
package test

type Source = {
    a :: Integer
    b :: Integer
}
type Sum = {
    sum :: Integer
}
type Result = {
    result :: Integer
}
github.com/project/example/sum[go] :: Source -> Sum
github.com/project/example/result[go] :: Sum -> Result
calc = sum . result
invoke(calc,)
```

## Generate the source code

Use command 

`anzer generate -i example.anz -f github.com/project/example/sum -o ~/go/src/github.com/project/example/sum/gen.go`


to generate common Golang types and handler function for the cloud function `github.com/project/example/sum`.

The contents of the `~/go/src/github.com/project/example/sum/gen.go` file must be like the following code:

```go

``` 