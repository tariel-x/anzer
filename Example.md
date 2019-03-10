# Part of the example scheme

```haskell
type Req = {
    brand  :: String
    model  :: String
    body   :: Optional String
    year   :: Integer
    price  :: Float
    phone  :: String
    photos :: List String
}

type Car = {
    brand      :: String
    model      :: String
    generation :: String
    body       :: Optional String
    year       :: Integer
    price      :: Float
    phone      :: String
    rawImages  :: List String
}

tr github.com/tariel-x/anzer-examples/transform[go] :: Req -> Car
```

## Create POST api:

1. Build `wsk -i action update github.com_tariel-x_anzer-example_transform --web true`.
2. Create api endpoint `wsk -i api create /transform /post post github.com_tariel-x_anzer-example_transform --response-type json` and get endpoint url.
3. Make the following request

```bash
curl -X POST \
  http://wsk.tariel.space:9001/api/23bc46b1-71f6-4ed5-8c54-816aa4f8c502/transform/post \
  -H 'Content-Type: application/json' \
  -d '{
    "brand": "Opel",
    "model": "Astra J",
    "phone": "79095544445",
    "price": 500.50,
    "photos": [
    	"http://origin.org/path/image.jpeg"	
    ],
    "year": 2014
}'
```

The expected response is:
```json
{
    "value": {
        "body": null,
        "brand": "Opel",
        "generation": "J",
        "model": "Astra",
        "phone": "79095544445",
        "price": 500,
        "rawImages": [
            "http://origin.org/path/image.jpeg"
        ],
        "year": 2014
    }
}
```