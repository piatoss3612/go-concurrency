# Goroutines, go 키워드 그리고 WaitGroups

## 1. Goroutine?

Go 언어에서 동시성 프로그래밍을 가능하게 해주는 lightweight threads입니다

매우 빠르고 Go Scheduler에 의해 관리됩니다

main 함수는 그 자체로 하나의 goroutine 입니다

## 2. go 키워드

go 키워드는 새로운 goroutine을 실행하기 위해 사용할 수 있습니다

go 키워드 뒤에 붙는 동작은 goroutine으로 실행되고 Go Scheduler에 의해 관리됩니다
<br>

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
<br>

이는 goroutine이 정상적으로 실행되기 전에 프로그램이 너무 빠르게 종료됨으로 인해 발생한 문제입니다

첫 번째 printSomething 함수를 goroutine으로 정상 실행시키기 위해 어떠한 방법을 사용해야 될까요?

### time 패키지의 Sleep 메서드

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
<br>

그런데 만약 goroutine을 사용해 데이터베이스에서 수천, 수만 개의 데이터를 가져온다고 하면

main 함수가 기다려야 하는 시간을 정확하게 계산하는 것이 가능할까요?
<br>

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

실행 결과로 출력되는 문장의 수는 천차만별입니다

여기서 main 함수의 대기 시간을 충분히 늘리면 해결할 수 있지 않을까? 하는 의문이 듭니다
<br>

그러나 이는 테스트 환경에서의 임시방편일 뿐, 

서버, 네트워크 등의 여러가지 외부 요인에 영향을 받는 프로덕션 환경에서

time.Sleep 메서드는 goroutine이 정상적으로 실행되는 것을 보장해줄 수 없습니다

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

## 3. WaitGroup

WaitGroup은 Go언어의 표준 라이브러리 sync 패키지에서 제공하는 구조체입니다

이름에서 유추할 수 있듯이 WaitGroup은 여러 개의 goroutine이 정상 실행/종료될 때까지 기다리기 위해 사용합니다 

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup // WaitGroup을 사용하기 위해 wg라는 이름의 변수로 선언합니다

	words := []string{"alpha", "beta", "delta", "pi", "zeta", "eta", "theta", "epsilon"}

	// 실행할 goroutine의 수를 WaitGroup에 더해줍니다
	// 여기선 words 슬라이스를 range 키워드로 순회한만큼 goroutine을 실행하므로
	// words 슬라이스의 길이인 len(words)를 더해줍니다
	wg.Add(len(words)) 

	for i, word := range words {
		// WaitGroup 타입의 변수 wg의 포인터를 함수의 인수로 넘겨줍니다
		// 값을 복사하여 넘겨주는 것은 예상치 못한 오류를 유발할 수 있습니다
		go printSomething(fmt.Sprintf("%d: %s", i, word), &wg)
	}

	// WaitGroup에 지정된 카운터가 0이 될 때까지
	// 즉 실행하고자 하는 goroutine이 모두 정상 실행/종료될 때까지 대기합니다
	wg.Wait()

	// 만약 WaitGroup에 카운터를 1만큼 추가하지 않는다면
	// 아래의 printSomething 함수가 실행되고
	// 함수가 종료될 때 실행된 wg.Done() 메서드가 WaitGroup의 카운터를 -1로 만들어
	// 패닉이 발생하게 됩니다
	//  panic: sync: negative WaitGroup counter
	wg.Add(1)

	printSomething("This is the second thing to be printed!", &wg)
}

// printSomething 함수는 WaitGroup의 포인터를 받습니다
func printSomething(s string, wg *sync.WaitGroup) {
	// defer 키워드는 현재 실행중인 함수가 종료될 때 다음으로 오는 동작을 실행하라는 의미입니다
	// 따라서 함수가 종료될 때 wg.Done() 메서드가 실행되고 
	// 이때 WaitGroup에 지정된 카운터가 1만큼 감소합니다 
	defer wg.Done()
	fmt.Println(s)
}
```

```cmd
$ go run .
0: alpha
3: pi
1: beta
4: zeta
5: eta
6: theta
7: epsilon
2: delta
This is the second thing to be printed!
```

WaitGroup에 지정된 카운터가 0보다 작아지게 되면 패닉(panic: sync: negative WaitGroup counter)이 발생하게 됩니다

반대로 카운터값이 실행할 goroutine의 수보다 크다면 어떤 문제가 발생할까요?

```go
wg.Add(12)
```

len(words) 값은 8이므로 총 8개의 goroutine이 실행되는데 WaitGroup 카운터값으로 12를 설정했습니다

```cmd
$ go run .
0: alpha
7: epsilon
1: beta
4: zeta
5: eta
6: theta
2: delta
3: pi
fatal error: all goroutines are asleep - deadlock!
```

출력값을 보니 8개의 goroutine이 정상 실행/종료된 뒤에 

나머지 4개의 goroutine을 WaitGroup이 한없이 기다린 결과,

deadlock이 발생하여 프로그램이 치명적으로 종료된 것을 확인할 수 있습니다

