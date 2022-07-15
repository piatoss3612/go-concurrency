# Working with Concurrency in Go (Golang)

## 1. Goroutines, go keyword and   WaitGroups

### Test WaitGroup

```cmd
$ go test -v .
=== RUN   Test_printSomething
--- PASS: Test_printSomething (0.00s)
PASS
ok      example01       (cached)
```

### Test challenge-1-solution

```cmd
$ go test -v .
=== RUN   Test_updateMessage
--- PASS: Test_updateMessage (0.00s)
=== RUN   Test_printMessage
--- PASS: Test_printMessage (0.00s) 
=== RUN   Test_main
--- PASS: Test_main (0.00s)
PASS
ok      challenge-1-solution    0.042s
```

## Reference

[Working with Concurrency in Go (Golang)](https://www.udemy.com/course/working-with-concurrency-in-go-golang/)