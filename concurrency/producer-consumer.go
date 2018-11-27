package main

import (
	"fmt"
	"math/rand"
	"time"
)

type producer struct {
	id    int
	level int
}

type consumer struct {
	id int
}

func (p producer) produce(inputC chan<- int, waitC chan<- bool) {
	x := rand.Intn(9) * p.level
	for i := 0; i < 30; i++ {
		inputC <- x
		fmt.Println("Producer: ", p.id, " produced: ", x)
		fmt.Println("Producer: ", p.id, " produced: ", x)
		sleepTime := rand.Intn(1000) + 1000
		time.Sleep(time.Millisecond * time.Duration(sleepTime))
	}
	waitC <- true
	fmt.Println("Producer id: ", p.id, " quit.")
}

func (c consumer) consume(inputC <-chan int, waitC chan<- bool) {
	for {
		select {
		case x, ok := <-inputC:
			if ok {
				fmt.Println("Consumer: ", c.id, " consumed:", x)
			}
		default:
			fmt.Println("Consumer ", c.id, " quit")
			waitC <- true
			return
		}
		sleepTime := rand.Intn(1000) + 2000
		time.Sleep(time.Millisecond * time.Duration(sleepTime))
	}
}

func main() {
	var pLevels = [8]int{1, 10, 100, 100, 10000, 100000, 1000000, 10000000} // 定义不同的Level来区分生产者
	const producerCount int = 5                                             // 生产者数量
	const consumerCount int = 10                                            // 消费者数量
	var inputC = make(chan int, producerCount*4)                            // 生产队列，存放产品
	var waitC = make(chan bool)                                             // 等待队列，用于main方法等待所有消费者结束
	var producers [producerCount]producer                                   // 生产者数组
	var consumers [consumerCount]consumer                                   // 消费者数组

	for i := 0; i < producerCount; i++ {
		producers[i] = producer{i, pLevels[i]} // 初始化生产者id和Level
		go producers[i].produce(inputC, waitC) // 开始生产
	}

	time.Sleep(time.Second * 5) // 积累一定产品后再启动消费者

	for i := 0; i < consumerCount; i++ {
		consumers[i] = consumer{i}             // 初始化消费者
		go consumers[i].consume(inputC, waitC) // 开始消费
	}

	for i := 0; i < (producerCount + consumerCount); i++ {
		fmt.Println("waitC")
		<-waitC // 等待所有生产者和消费者结束退出
	}

	fmt.Println("closed")
	close(inputC)
	close(waitC)
}
