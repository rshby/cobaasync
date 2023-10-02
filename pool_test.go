package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// buat function untuk menggunakan pool
func TestPool(t *testing.T) {
	pl := &sync.Pool{}      // buat variabel pool, yang bisa digunakan untuk menyimpan (menimbun) data
	wg := &sync.WaitGroup{} // buat variabel waitgroup untuk menunggu semua goroutine selesai, dan memastikan semuanya selesai dijalankan

	pl.Put("reo")     // simpan value 'reo' ke pool
	pl.Put("sahobby") // simpan value 'sahobby' ke pool

	// cara mengambil pool dengan goroutine
	for i := 0; i < 5; i++ { // buat looping 2 kali perulangan (sesuai banyaknya data yang ada di pool)
		wg.Add(1)   // beri tahu waitgroup bahwa setiap perulangan akan menambah 1 goroutine yang jalan
		go func() { // jalankan goroutine dengan anonymous function di setiap perulangan, nanti bakal ada 2 kali perulangan
			defer wg.Done() // diakhir sebelum goroutine selesai dijalankan, beri tahu waitgroup bahwa goroutine ini sudah selesai dijalankan

			dataPool := pl.Get()  // ambil value yang ada di pool (urut sesuai perulangan), simpan ke dalam variabel dataPool
			fmt.Println(dataPool) // tampilkan isi value
			pl.Put(dataPool)      // masukkan lagi valuenya ke pool
		}()
	}

	wg.Wait() // tunggu sampai waitgroup mendapat informasi bahwa semua goroutine sudah selesai dijalankan
	fmt.Println("selesai")
}

// membuat function untuk override value ketika pool tidak ada datanya
func TestPoolOverrideDefault(t *testing.T) {
	wg := &sync.WaitGroup{} // buat variabel waitgroup untuk menunggu dan memastikan semua goroutine sudah selesai dijalankan
	pool := &sync.Pool{     // buat variabel pool, untuk menyimpan(menimbun) data yang nantinya bisa dipake
		New: func() any { // override property New
			return "default" // returnkan value string 'default' ketika pool akan diambil valuenya tapi kosong tidak ada datanya
		},
	}

	pool.Put("reo")     // simpan value 'reo' ke pool
	pool.Put("sahobby") // simpan value 'sahobby' ke pool

	for i := 0; i < 10; i++ { // buat looping 10 kali perulangan
		wg.Add(1)   // beri tahu waitgroup bahwa setiap perulangan akan bertambah 1 goroutine yang jalan
		go func() { // buat goroutine di setiap perulangan (menggunakan anonymous function)
			defer wg.Done() // di akhir goroutine sebelum selesai, beri tahu waitgroup bahwa goroutine ini selesai dijalankan

			dataPool := pool.Get() // ambil data dari pool, simpan ke dalam variabel dataPool
			fmt.Println(dataPool)  // tampilkan value datanya

			time.Sleep(1 * time.Second) //goroutine akan pause selama 1 second
			pool.Put(dataPool)          // simpan lagi valuenya ke pool
		}()
	}

	wg.Wait() // tunggu waitgroup mendapat informasi bahwa semua goroutine sudah selesai dijalankan
	fmt.Println("- selesai -")
}
