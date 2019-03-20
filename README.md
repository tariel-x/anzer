# Anzer platform and language

Both the provided platform and described language implement 
convenient way for explicit and type-safe composition of cloud functions.

For example the following scheme describes two services and 
their communication.

```haskell
type Source = {
    name  :: String
    price :: Float
}
type Prepared = {
    name        :: String
    description :: String
    category    :: Integer
    price       :: Float
}
type Product = {
    id          :: Integer
    name        :: String
    description :: String
    category    :: Integer
    price       :: Float
}
github.com/project/prepare[go] :: Source -> Prepared
github.com/project/save[go] :: Prepared -> Product
create = prepare . save
invoke(create,)
```

Anzer utility currently supports only [Apache OpenWhisk](http://openwhisk.apache.org/) and [Golang](https://golang.org/).

## As quick as possible start

### Setup

1. Install docker for your system.
2. Download and install wsk utility from [github page](https://github.com/apache/incubator-openwhisk-cli/releases).
3. Find working Apache OpenWhisk instance.

### Download Anzer

```bash
go get github.com/tariel-x/anzer
anzer help
```

### Download example project

```bash
git clone https://github.com/tariel-x/anzer-example.git
cd anzer-example
anzer build -p wsk -i etl.anz
```

The output of utility would be name of functions sequence in OpenWhisk.

### Create POST API

```bash
wsk -i api create /etl post /guest/etl_sequence --response-type json
```

Make request like the following:

```bash
curl -X POST \
  http://your.whisk.com:9001/api/23bc46b1-71f6-4ed5-8c54-816aa4f8c502/etl \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "Opel",
    "model": "Astra J",
    "phone": "79095544445",
    "price": 500.50,
    "photos": [
    	"http://origin.org/path/image.jpeg"	
    ],
    "year": 2014
}'
```

The expected result is:

```json
{
    "body": null,
    "brand": "Opel",
    "generation": "J",
    "id": 38081,
    "model": "Astra",
    "phone": "+790911222233",
    "photos": [
        "http://storage.org/photo.jpeg"
    ],
    "price": 55000,
    "year": 2014
}
```

## Full documentation

Full documentation is in progress...