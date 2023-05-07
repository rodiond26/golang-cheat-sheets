### Простые типы данных

---

#### Целые числа

* Знаковые:
    * int, int8, int16, int32, int64.
* Беззнаковые:
    * uint, uint8, int16, uint32, uint64.
* Символы unicode:
    * type rune = int32.
* Байты:
    * type byte = uint8.

#### Числа с плавающей точкой

* float32 (точность ~6 знаков), float64 (точность ~15 знаков).
* complex64, complex128.

#### Логический тип:

* bool:
    * true, false.

#### Строки

* Строка - это неизменяемая последовательность байт.
* Примеры строк:

```go
var hello = "\thello\t" // строка с двойными кавычками   = "        hello"
var title = `\thello\t` // строка с одинарными кавычками = `\thello\t`
word := "hello"
w[0] = "H" // будет ошибка, т. к. строки неизменяемы
```

* Строки из коробки поддерживают UTF-8.
* Если записать символ в одинарных кавычках, то вернется его байтовое представление.

```go
str := "12345"
fmt.Println(str[0]) // обращение к первому байту, а не символу в строке
len(str)            // возвращает количество байтов, а не символов в строке
str[2:4] // возвращает байты в строке со 2 до 4
```

* Строки можно сравнивать, они сравниваются лексикографически (т. е. как в словаре).
* При конкатенации возвращается новая строка, при этом старые строки не изменяются.
* Один символ в строке может содержать 8 (byte = uint8) или 32 (rune = int32) байта.
* Получение подстроки:

```go
str[from:to] // получение подстроки
str[from:] // получение подстроки, где to = len(str)
str[:to]   // получение подстроки, где from = 0
```