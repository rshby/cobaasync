package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestOnce(t *testing.T) {
	onc := &sync.Once{}     // buat variabel once untuk menjalankan function di dalam goroutin, namun hanya satu kali dijalankan
	wg := &sync.WaitGroup{} // buat variabel waitgroup untuk menunggu semua goroutine yang sedang dijalankan, dan memastikan semuanya selesai dijalankan

	x := 0 // buat variabel x integer dengan value awal 0

	for i := 0; i < 10; i++ { // buat looping 10 kali perulangan
		wg.Add(1)   // beri tahu waitgroup setiap perulangan akan menambah 1 goroutine yang akan dijalankan
		go func() { // buat goroutine dengan anonymous function di setiap perulangan, jadi bakal ada 10 goroutine
			defer wg.Done() // diakhir sebelum goroutine selesai, beri tahu waitgroup bahwa function ini sudah selesai dijalankan

			onc.Do(func() { // jalankan function tapi hanya satu kali dijalankan, karena menggunakan onc.Do()
				x++ // isi function yang dijalankan adalah menambah increment x
			})
		}()
	}

	wg.Wait()      // tunggi waitgroup mendapat informasi bahwa semua goroutine sudah Done(), atau selesai dijalankan semua
	fmt.Println(x) // tampilkan informasi hasil increment x yang ditambah menggunakan goroutine
}
