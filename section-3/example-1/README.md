# Race Conditions, Mutexes, and Channels

## 1. Race Condition

```go
package main

import (
	"fmt"
	"sync"
)

var (
	msg string
	wg  sync.WaitGroup
)

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func main() {
	msg = "Hello, World!"

	wg.Add(2)

	go updateMessage("Hello, Universe!")
	go updateMessage("Hello, Cosmos!")

	wg.Wait()

	fmt.Println(msg)
}
```

```
$ go run .
Hello, Universe!
```

```
$ go run .
Hello, Cosmos!
```
<span>
상단의 코드를 여러 번 실행해보면 2가지 결과가 번갈아서 나오는 것을 확인할 수 있습니다

2개의 goroutine이 동시에 실행되면서 msg 변수에 동시에 접근하여 값을 변경하고자 하므로

드래곤볼에서 누가 먼저 상대방의 등 뒤로 순간이동하는지 경쟁하는 것처럼

어떤 goroutine이 먼저 종료되고 msg 변수의 값이 무엇으로 변경될지 장담할 수 없습니다
</span>

```cmd
$ go run -race main.go
==================
WARNING: DATA RACE
Write at 0x0000005a87a0 by goroutine 8:
	...
==================
Hello, Cosmos!
Found 1 data race(s)
exit status 66
```
<span>
이처럼 동시성 프로그래밍에서 동시에 동일한 자원에 접근하는 상황을 경쟁 상태(Race Condition)라고 합니다

그렇다면 어떤 자원을 하나의 goroutine이 사용하는 동안 다른 goroutine의 접근을 제한하는 방법은 무엇이 있을까?
</span>


## 2. Mutex

<span>
Go 언어에서는 sync 패키지의 Mutex 구조체를 사용하여 다른 goroutine이 사용하고 있는

자원에 대한 접근을 제한할 수 있습니다

접근을 제한하는 방법은 굉장히 간단합니다

Mutex 객체의 포인터를 실행하고자 하는 goroutine으로 넘겨줍니다

그리고 Mutex 객체의 Lock 메서드를 실행해주면 다른 goroutine들은

Mutext 객체의 Unlock 메서드가 실행될 때까지 실행되지 않습니다
</span>

```go
package main

import (
	"fmt"
	"sync"
)

var (
	msg string
	wg  sync.WaitGroup
)

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	msg = s
	m.Unlock()
}

func main() {
	msg = "Hello, World!"

	var mutex sync.Mutex

	wg.Add(2)

	go updateMessage("Hello, Universe!", &mutex)
	go updateMessage("Hello, Cosmos!", &mutex)

	wg.Wait()

	fmt.Println(msg)
}
```

<span>
이번에도 코드를 여러 번 실행해보면 두 개의 결과가 번갈아서 나오는 것을 확인할 수 있습니다

그러나 적어도 자원에 동시에 접근함으로인해 발생하는 Data Race 문제는 해결된 것을 확인할 수 있습니다

이전보다 데이터에 안전하게 접근할 수 있게 된 것입니다
</span>

```
$ go run -race main.go
Hello, Universe!
```

```cmd
$ go run -race main.go
Hello, Cosmos!
```

## 3. Test: Race Condition

```cmd
$ go test -race .
==================
WARNING: DATA RACE
Write at 0x000000b8e730 by goroutine 8:
	...
==================
--- FAIL: Test_updateMessage (0.00s)
    testing.go:1312: race detected during execution of test
FAIL
FAIL    example-1       0.369s
FAIL
```