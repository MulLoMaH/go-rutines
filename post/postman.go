package post

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func postman(
	stop int,
	ctx context.Context,
	wg *sync.WaitGroup,
	ch chan<- string,
	number int,
	post string) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		fmt.Println("Контекст отменен на почтальоне №", number)
		return
	case <-time.After(time.Duration(stop) * time.Second):
		fmt.Println("Прошло время и я взял письмо. Почтальон №", number)
	}

	select {
	case <-ctx.Done():
		fmt.Println("Контекст отменен на почтальоне №", number)
		return
	case ch <- post:
		fmt.Printf("Сообщение #%d доставлено\n", number)
	}
}

func PostLogic(
	ctx context.Context,
	number int) <-chan string {
	wg := &sync.WaitGroup{}
	ch := make(chan string)

	wg.Add(number)
	for i := 1; i <= number; i++ {
		go postman(1, ctx, wg, ch, i, post())
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func post() string {
	return "random"
}
adsfgsdfgsdf

ds
fg
Secondfg
Secondfggs
dfg
