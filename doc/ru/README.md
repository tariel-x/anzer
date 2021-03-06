# Anzer: платформа и язык

Anzer - это средство для композиции функций в serverless, состоящее из языка и
клиента к FaaS-платформам.

Композиция функций - это объединение их в последовательность, в которой результаты 
работы каждой передаются следующей по цепочке. Ближайшие аналоги Anzer: 
AWS Step Functions, Azure App Logic, Fn Workflows и другие.

Возможности языка:

1. Описание последовательностей выполнения функций.
2. Определение типов сигнатур облачных функций.
3. Определение порадка обработки ошибок.

Возможности платформы:

1. Проверка согласованности типов функций в схеме композиции.
2. Проверка корретной реализации типов аргументов и результатов функций.
3. Сборка и деплой функций в FaaS-платформе.

## Установка

### Исполняемый файл из Github

### `go get`

```
go get github.com/tariel-x/anzer
anzer help
```

### Сборка из исходного кода

## Настройка

Утилита Anzer использует тот же конфигурационный файл, что и OpenWhisk. Если у Вас 
уже установлена и настроена утилита `wsk`, никакой дополнительной настройки не требуется.
Иначе, выполните команду:

```
anzer init -p wsk
```

После этого будет создан конфигурационный файл клиента OpenWhisk.