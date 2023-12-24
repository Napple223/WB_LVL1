Что выведет данная программа и почему?

```go

func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}

```

Ответ:

```
Вывод: [b b a][a a]

В анонимной функции, при выполнении append, был создан новый слайс,
с которым происходят дальнейшие манипуляции.
Они не затрагивают изначальный слайс [a, a].
```