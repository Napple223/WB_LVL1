Что выведет данная программа и почему?

```go

func main() {
  n := 0
  if true {
     n := 1
     n++
  }
  fmt.Println(n)
}

```

Ответ:

```
Вывод: 0.

В зоне if своя зона видимости, в которой инициализируется
и инкрементируется локальная переменная n.
```