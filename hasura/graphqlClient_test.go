package hasura

import (
	"log"
	"sync"
	"testing"
)

type abcsOvbInternal struct {
	StatusCode int
	statusDesc string
}

func TestHasura(t *testing.T) {
	wg := &sync.WaitGroup{}

	//client := graphql.NewClient("https://localhost:8003/graphql", nil)

	for i := 0; i < 2; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()

			//err := client.Query(context.Background())
		}()
	}

	wg.Wait()
	log.Println("done")
}
