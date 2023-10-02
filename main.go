package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// membuat channel

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) // bikin context timeout 2 detik
	defer cancel()

	chanDone := make(chan bool, 1) // buat varibale channel untuk menampung sinyal done apakah proses sudah selesai
	defer close(chanDone)

	go TimeOutTest(ctx, chanDone) // jalanin goroutinenya jalan async

	for { // looping terus menerus
		select { // pilih salah satu
		case <-ctx.Done(): // in case jika kena timeout
			fmt.Println("kena timeout karena masih diproses lebih dari 2 second")
			return
		case <-chanDone: // in case jika selesai sebelum timeout
			fmt.Println("selesai dikerjakan sebelum batas timeout")
			return
		}
	}
}

func TimeOutTest(ctx context.Context, chanDone chan bool) {
	defer ctx.Done()

	time.Sleep(3 * time.Second)

	// get data
	// inserr, dll
	fmt.Println("sukses dikerjakan")
	chanDone <- true
}
