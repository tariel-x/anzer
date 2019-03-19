# Anzer platform and language

Both the provided platform and described language implement 
convenient way for verbose and type-safe composition of cloud functions.

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

## Full documentation

Full documentation is in progress...