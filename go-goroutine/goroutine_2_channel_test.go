package gogoroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func helloWorld(channel chan string) {
	channel <- "Hello World"

	fmt.Println("Finish Send Data")
}

func helloWorldIn(channel chan<- string) {
	channel <- "Hello World"

	fmt.Println("Finish Send Data")
}

func helloWorldOut(channel <-chan string) {
	data := <-channel

	fmt.Println(data)

	fmt.Println("Finish Receive Data")
}

func TestChannel(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	go func() {
		channel <- "Hello World"

		fmt.Println("Finish Send Data")
	}()

	data := <-channel

	fmt.Println(data)

	fmt.Println("Finish Receive Data")

	time.Sleep(3 * time.Second)
}

func TestChannelAsParamater(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	go helloWorld(channel)

	result := <-channel

	fmt.Println(result)

	fmt.Println("Finish Receive Data")

	time.Sleep(3 * time.Second)
}

func TestChannelInOut(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	go helloWorldIn(channel)

	go helloWorldOut(channel)

	time.Sleep(3 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)

	defer close(channel)

	fmt.Println(cap(channel))

	for i := 1; i <= 3; i++ {
		channel <- "Hello World - " + strconv.Itoa(i)
	}

	for i := 1; i <= 3; i++ {
		fmt.Println(len(channel))

		data := <-channel

		fmt.Println(data)
	}

	time.Sleep(3 * time.Second)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 1; i <= 10; i++ {
			channel <- "Hello World - " + strconv.Itoa(i)
		}

		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	time.Sleep(3 * time.Second)
}

func TestSelectChannel(t *testing.T) {
	channelOne := make(chan string)
	channelTwo := make(chan string)
	channelThree := make(chan string)

	defer close(channelOne)
	defer close(channelTwo)
	defer close(channelThree)

	go func() {
		for i := 1; i <= 10; i++ {
			channelOne <- "Hello World - " + strconv.Itoa(i)
			channelTwo <- "Hello World - " + strconv.Itoa(i)
			channelThree <- "Hello World - " + strconv.Itoa(i)
		}

		time.Sleep(1 * time.Second)
	}()

	for i := 1; i <= 30; {
		select {
		case data := <-channelOne:
			fmt.Println("Channel One -", data)

			i++
		case data := <-channelTwo:
			fmt.Println("Channel Two -", data)

			i++
		case data := <-channelThree:
			fmt.Println("Channel Three -", data)

			i++
		default:
			fmt.Println("Waiting For Data")
		}
	}

	time.Sleep(3 * time.Second)
}
