# Quick start with the Apache OpenWhisk instance

Setup the OpenWhisk instance with IBM CloudFunctions:

1. Register at [cloud.ibm.com](https://cloud.ibm.com).
2. Follow the instruction at [cloud.ibm.com/openwhisk/learn/cli](https://cloud.ibm.com/openwhisk/learn/cli).
3. Check the settings with `wsk action list`. Also `~/.wskprops` must contain the content like the following.
```
APIHOST=eu-gb.functions.cloud.ibm.com
NAMESPACE=_
AUTH=short_token
APIGW_ACCESS_TOKEN=big_token
```

You can also use `ibmcloud fn` utility instead of the `wsk` one.

Any [self-hosted OpenWhisk instance](https://openwhisk.apache.org/documentation.html#openwhisk_deployment) is also suitable.

Then install Anzer cli as [described here](./Installation.md).

## Definition of the types and functions.

Create the file named `example.anz` with the contents:

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

Replace `github.com/project/example` with the actual name of the repository you are using.

## Generate the source code

### Generate `github.com/project/example/sum`

Use commands 

```bash
mkdir sum
anzer generate -i example.anz -f github.com/project/example/sum -o ~/go/src/github.com/project/example/sum/gen.go
```

to generate common Golang types and handler function for the cloud function `github.com/project/example/sum`.

The contents of the `~/go/src/github.com/project/example/sum/gen.go` file must be like the following code:

```go
// Thank robots for this file that was generated for you at 2020-07-05 18:07:57.065434549 +0300 MSK m=+0.037309363
package sum

type TypeIn struct {
	A int `json:"a"`
	B int `json:"b"`
}

type TypeOut struct {
	Sum int `json:"sum"`
}

func Handle(input TypeIn) TypeOut {
	var out TypeOut
	return out
}
``` 

Edit file in the following way:

```go
// Thank robots for this file that was generated for you at 2020-07-05 18:07:57.065434549 +0300 MSK m=+0.037309363
package sum

type TypeIn struct {
	A int `json:"a"`
	B int `json:"b"`
}

type TypeOut struct {
	Sum int `json:"sum"`
}

func Handle(input TypeIn) TypeOut {
	var out TypeOut
	out.Sum = input.A + input.B
	return out
}
```

### Generate `github.com/project/example/result`

Generate the second function:

```bash
mkdir result
anzer generate -i example.anz -f github.com/project/example/result -o ~/go/src/github.com/project/example/result/gen.go
```

Edit generated file in the following way:

```go
// Thank robots for this file that was generated for you at 2020-07-05 18:16:56.626712397 +0300 MSK m=+0.009667573
package result

import (
	"log"
	"os"
	"strconv"
)

type TypeIn struct {
	Sum int `json:"sum"`
}

type TypeOut struct {
	Result int `json:"result"`
}

func Handle(input TypeIn) TypeOut {
	var out TypeOut
	rawMultiplicator := os.Getenv("__ANZ_MULTIPLY")
	multiplicator, err := strconv.Atoi(rawMultiplicator)
	if err != nil {
		log.Printf("error parsing __ANZ_MULTIPLY into int: %q", err)
		return out
	}
	out.Result = input.Sum * multiplicator
	return out
}
```

## Configure

Create file `envs.yaml` with environment variables:

```yaml
github.com/project/example/result:
  MULTIPLY: "5"
```

Variable `MULTIPLY` would be available in the `result` function with the name `__ANZ_MULTIPLY`.

## Deploy

Commit all files and push the changes to the repository `github.com/project/example`. Deploy functions composition 
to an OpenWhisk instance:

```bash
anzer build -p wsk -i example.anz --envs envs.yaml
``` 

The result of the command execution would be like `Your function name is: /ibm@cloud.account/test/calc_sequence`.

Also the generated `anzer.sum` file would store git commit hashes of the currently using functions.

```csv
github.com/project/example/sum,06fd5d4fca1b022d6686ab4aed07e3f9c9b6d1fe,
github.com/project/example/result,06fd5d4fca1b022d6686ab4aed07e3f9c9b6d1fe,
```

The function can be updated with the command `anzer sum update -f github.com/project/example/result`.

## Create OpenWhisk API

To create OpenWhisk API endpoint use the command `wsk -i api create /test /calc post /ibm@cloud.account/test/calc_sequence --response-type json`.
Or use [IBM Cloud web interface](https://cloud.ibm.com/functions/apimanagement) to create it. The command or the web interface
would return the sequence URL, e.g. `https://xxxxxxxx.xx-xx.apiconnect.appdomain.cloud/test/sum`.

Make POST query to test the composition.

```bash
curl --location --request POST 'https://xxxxxxxx.xx-xx.apiconnect.appdomain.cloud/test/sum' \
--header 'Content-Type: application/json' \
--data-raw '{
    "a": 1,
    "b": 2
}'
```

The expected result would be

```json
{
  "result": 15
}
```