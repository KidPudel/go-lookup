# Goroutines
> Goroutines - lightweight threads.

Unlike Kotlin's coroutines, which are kind of workers in a thread, **Goroutines are _actual threads_**, but way more lightweight.  

Unlike usual kernel threads, that are managed by operating system, **goroutines are _user-space threads managed by Go runtime_**.

```go
func sleepTrack(start int) {
  for i := 0; i < 9; i++ {
    fmt.Println("Sleeping, which started at ", start)
    time.Sleep(1000)
  }
}
```

```go
go sleepTrack(start: 19)
```
⬆️ starts a new goroutine running ⬇️
```go
sleepTrack(start: 19)
```

The evaluation of `sleepTrack` and `start` happens is current goroutine,  
While executing happens in a new goroutine.

> IMPORTANT: Goroutines run the same address space, so accessing shared data MUST be in synchronized way (to avoid deadlocks and rase conditions), for that Go has useful primitives like mutexes and atomic for protecting shared data accross Goroutines


# Channels
> Channels are typed conduct (pipe, tube), through which you can send and receive data with the help of channel operator `<-`

### send data thsrough channel
```go
func countSum(arr []int, ch chan int) {
  sum := 0
  for i, v := range arr {
    sum += v
  }
  // sending data through conduct
  ch <- sum
}
```

### receiving data from channel
```go
rightSide := <-ch
```

### Creating channel
```go
ch := make(chan int)
```

### waiting for the result
By default, sends and receives operations are block execution of their goroutine where they heppened, until the other side is ready.  

**Simple example**: calculating goroutine starts to execute algorithm, we encounter receive operation of channel that is in calculating goroutine, since algorithm is still processing (send operation is not encouintered), receive will block the further execution on its side, unitl we encounter send operation.

**Second example**: calculating goroutine already encountered the result, but doesn't see that see that channel is ready to receive (doesn't encountered receive operation), so it will be blocked until, we call to receive a value.
```go

s := []int{7, 2, 8, -9, 4, 0, 1}
ch := make(chan int)

go countSum(s[len(s)/2:], ch)
go countSum(s[:len(s)/2], ch)

// waiting for sending
x, y := <-ch, <-ch

fmt.Println(x, y, x+y) // 17 -4 13
```
Here you may ask: _"How does `x, y := <-ch, <-ch` works, that launching 2 goroutines with one channel, still asigns each variable the **correct part** of the `s` slice?"_.  

After starting both goroutines, we encounter first receive operation (block main function), meanwhile, there are 2 goroutines that are executed, and which ever executes first, will go to the first receiver, and we encounter next receiver and giving it the next _send_ data.


## Buffered channels
We can buffer a channel, by adding a number of buffers as a 2nd parameter in `make` function. Sends won't block until buffer is full, and receives won't block until buffer is empty
```go
bCh := make(chan int, 2)

bCh <- 10
// won't be blocked, bacause buffer is still have 1 space left
bCh <- 5

x := <-bCh // 10
// won't be blocked, because there is still 1 value left
y := <-bCh // 5
```

if `bCh` channel would have 1 buffer, then it would cause a **deadlock**

## Closing and Ranges
We can close a channel, to indicate that there is no values will be send.
> NOTE: Only sender can close channel, otherwise, sending to closed channel, will cause a panic

```go
package main

import "fmt"

func fibonacci(ch chan int) {
	x, y := 0, 1
	for i := 0; i < cap(ch); i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	fmt.Println("")
	ch := make(chan int, 10)

	go fibonacci(ch)

	fmt.Println("do some for me")

	for v := range ch {
		fmt.Println(v)
	}

	possibleValue, ok := <-ch

	if !ok {
		fmt.Println("Already closed")
	} else {
		fmt.Println(possibleValue)
	}

}

```

## Select
Imagine we wait (or listen) on one of the asynchronious events, asynchroniously, meaning we need channels, and which is going to trigger, the corresponding handler will be executed.  

_And for that, `select` lets goroutine to wait on mutliple operations_.  

Remember, how goroutine will be blocked when we 1. send to the channel, until we encounter a receive. 2. Receive from the channel, until we send to it some value.  

```go
func eventListener(onSendMessage, onQuit) {
	// listen...
	for {
	}
}
```
