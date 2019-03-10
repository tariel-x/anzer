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