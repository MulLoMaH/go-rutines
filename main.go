package main

import (
	"1/post"
	"1/post/feature1"
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	resultCh := post.PostLogic(ctx, 6)
	wg.Add(1)

	go func() {
		defer wg.Done()
		for msg := range resultCh {
			fmt.Println("Получено:", msg)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Second)
		cancel()
	}()

	feature1.Feature1()
	wg.Wait()

}
