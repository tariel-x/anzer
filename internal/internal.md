# композиция

Композиция - передача `a . b` - передача результата a в функцию b.

```haskell
a :: String -> String
b :: String -> String
c :: String -> String


compose = a . b . c
```

```haskell
a :: String -> Either Error String
b :: String -> String
c :: String -> Either Error String

>>= :: Right String -> f -> String
>>= s f = f(s)

compose = a . >>= . b . c
compose = a >>= b c
```

Применение - применение параметров к функции. >>= применяет параметр к функции, но это не композиция.