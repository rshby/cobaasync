package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestMutex(t *testing.T) {
	x := 0 // buat variabel x tipe data int dengan value awal 0

	wg := &sync.WaitGroup{} // buat variabel waitgroup untuk menunggu semua goroutine yang jalan, dan memastikan semua selesai dijalankan
	mtx := &sync.Mutex{}    // buat variabel mutex untuk lock dan unlock goroutine yang sedang berjalan, supaya tidak terjadi race condition

	for i := 0; i < 1000000; i++ { // looping sebanyak 1juta perulangan
		wg.Add(1)   // beri sinyal ke waitgroup bahwa ada 1 goroutine baru yang jalan di setiap perulangan
		go func() { // jalankan goroutine di setiap perulangan menggunakan anonymous funtion, jadi nanti bakal ada 1juta goroutine yang dijalankan
			defer func() { // jalankan defer diakhir sebelum function goroutine selesai dijalankan
				mtx.Unlock() // unlock mutex
				wg.Done()    // beri tahu ke waitgroup kalau goroutine ini sudah selesai dijalankan
			}()

			mtx.Lock() // lakukan locking pada goroutine yang sedang jalan
			x++        // lakukan increment ke value variabel x
		}()
	}

	wg.Wait()                         // tunggu waitgroup mendapat informasi bawah semua goroutine yang ada sudah selesai dijalankan
	fmt.Println("value x adalah:", x) // tampilkan hasil increment value x
}
