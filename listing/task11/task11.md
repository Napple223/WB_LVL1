Что выведет данная программа и почему?

```go

func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 5; i++ {
     wg.Add(1)
     go func(wg sync.WaitGroup, i int) {
        fmt.Println(i)
        wg.Done()
     }(wg, i)
  }
  wg.Wait()
  fmt.Println("exit")
}

```

Ответ:

```
Сначал в рандомном порядке выведет числа от 0 до 4,
а потом дедлок, так как WaitGroup
необходимо передавать по указателю.
```