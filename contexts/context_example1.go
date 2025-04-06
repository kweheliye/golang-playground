package contexts

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Response struct {
	val int
	err error
}

func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 20)
	return 666, nil
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	responseChannel := make(chan Response)

	go func() {
		val, err := fetchThirdPartyStuffWhichCanBeSlow()
		responseChannel <- Response{
			val: val,
			err: err,
		}
	}()

	for {
		fmt.Println("running")
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fetching data from third party took long ")
		case resp := <-responseChannel:
			return resp.val, resp.err
		}
	}

}

func RunContextExample1() {

	start := time.Now()
	ctx := context.Background()
	userID := 10
	val, err := fetchUserData(ctx, userID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result:", val)
	fmt.Println("took:", time.Since(start))

}
