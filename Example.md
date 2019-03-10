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

Example curl to run transform:

```bash
curl -X GET \
  'http://wsk.tariel.space:9001/api/23bc46b1-71f6-4ed5-8c54-816aa4f8c502/create?brand=Opel&model=Astra%20J&price=700&year=2014&body=%D1%85%D1%8D%D1%82%D1%87%D0%B1%D1%8D%D0%BA%203%20%D0%B4%D0%B2&phone=795400005500' \

```