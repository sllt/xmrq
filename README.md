# xmrq



```go
func RunTasker(q *xmrq.Queue) {
  for i := 0; i < q.Length()+1; i++ {
    f, _ := q.Peek().(func())
    f()
    q.Remove()
  }
}

func hello() {
  fmt.Println("Hello")
}

func world() {
  fmt.Println("world")
}

func main() {

  q := xmrq.NewQueue()

  q.Add(hello)
  q.Add(world)

  RunTasker(q)
}

```