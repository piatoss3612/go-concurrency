# Goroutines, go 키워드 그리고 WaitGroups

### 1. Goroutine?

Go 언어에서 동시성 프로그래밍을 가능하게 해주는 lightweight threads입니다
매우 빠르고 Go Scheduler에 의해 관리됩니다
main 함수는 그 자체로 하나의 goroutine 입니다

### 2. go 키워드

go 키워드는 새로운 goroutine을 실행하기 위해 사용할 수 있습니다
go 키워드 뒤에 붙는 동작은 goroutine으로 실행되고 Go Scheduler에 의해 관리됩니다

아래의 프로그램은 첫 번째 printSomething 함수의 실행을 goroutine으로 실행하여
main 함수와 동시에 실행되도록 작성되었습니다

```go
package main

import "fmt"

func main() {
	go printSomething("This is the first thing to be printed!")

	printSomething("This is the second thing to be printed!")
}

func printSomething(s string) {
	fmt.Println(s)
}
```

```cmd
$ go run .
This is the second thing to be printed!
```

그런데 이 프로그램을 실행해보면 두 번째 printSomething 함수만 실행되고
goroutine으로 선언된 첫 번째 printSomething 함수 실행되지 않은 것을 확인할 수 있습니다

이는 goroutine이 정상적으로 실행되기 전에 프로그램이 너무 빠르게 종료됨으로 인해 발생한 문제입니다

첫 번째 printSomething 함수를 goroutine으로 정상 실행시키기 위해 어떠한 방법을 사용해야 될까요?

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go printSomething("This is the first thing to be printed!")

	time.Sleep(500 * time.Millisecond)

	printSomething("This is the second thing to be printed!")
}

func printSomething(s string) {
	fmt.Println(s)
}
```

```cmd
$ go run .
This is the first thing to be printed!
This is the second thing to be printed!
```

첫 번째로 가장 간단한 방법은 main 함수 goroutine이 바로 종료되지 않도록 잠깐 멈추는 것입니다
500 밀리초 정도면 단순히 한 줄의 문자열을 출력하는 goroutine은 실행하고도 남습니다

그런데 만약 goroutine을 사용해 데이터베이스에서 수천, 수만 개의 데이터를 가져온다고 하면
main 함수가 기다려야 하는 시간을 정확하게 계산하는 것이 가능할까요?

아래의 프로그램은 8개의 문자열을 출력하는데 각각의 goroutine을 실행하고
main 함수를 10 나노초동안 멈추도록 작성되었습니다

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	words := []string{"alpha", "beta", "delta", "pi", "zeta", "eta", "theta", "epsilon"}

	for i, word := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, word))
	}

	time.Sleep(10 * time.Nanosecond)

	printSomething("This is the second thing to be printed!")
}

func printSomething(s string) {
	fmt.Println(s)
}
```

결과는 과연 어떨까요? 결과는 천차만별입니다
프로그램을 실행할 때마다 출력되는 문장의 수가 달라지는 것을 확인할 수 있습니다
물론 대기 시간을 늘리면 해결할 수 있는 것처럼 보일 수 있습니다

그러나 서버, 네트워크 등의 외부 요인에 영향을 받는 프로덕션 환경에서
time.Sleep()은 goroutine이 정상적으로 실행되는 것을 보장해줄 수 없습니다

```cmd
$ go run .
0: alpha
This is the second thing to be printed!
```
```cmd
$ go run .
0: alpha
1: beta
5: eta
4: zeta
2: delta
6: theta
7: epsilon
3: pi
This is the second thing to be printed!
```

> Info: goroutine은 실행 순서가 정해져있지 않지만, 제어할 수 있는 방법이 있습니다