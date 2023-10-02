package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestSimulasiRaceCondition(t *testing.T) { // function untuk simulasi bagaimana race condition terjadi
	x := 0 // buat variabel x tipe data integer dengan value awal 0

	wg := &sync.WaitGroup{} // buat variabel waitgroup untuk menunggu semua goroutine selesai dikerjakan

	for i := 0; i < 1000000; i++ { // buat looping 1juta kali
		wg.Add(1)   // tambahkan penanda bahwa ada 1 goroutine yang dijalankan di setiap perulangannya
		go func() { // jalankan goroutine menggunakan anonymous function di setiap perulangan, jadi nanti ada 1juta goroutine yang bakal jalan
			defer wg.Done() // diakhir sebelum goroutine selesai, beri tahu bahwa ke waitgroup bahwa goroutine ini sudah selesai dijalankan
			x++             // tambahkan increment 1 variabel x
		}()
	}

	wg.Wait()      // tunggu hingga waitgroup mendapat informasi bahwa semua goroutine sudah selesai
	fmt.Println(x) // tampilkan valu hasil x, berapa increment yang dihasilkan
}
