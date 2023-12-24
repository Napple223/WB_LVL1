Сколько существует способов задать переменную типа slice или map?

Ответ:

```
4 способа:
```

```go
var s []int
s := []int{}
s := make([]int)
s := new([]int)
```

```go
var m map[int]int
m := map[int]int{}
m := make(map[int]int)
m := new(map[int]int)
```