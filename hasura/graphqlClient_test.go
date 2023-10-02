package hasura

import (
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
		wg.Add(1)
		go func() {
			defer wg.Done()

			//err := client.Query(context.Background())
		}()
	}

}
