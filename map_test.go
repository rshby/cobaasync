package main

import (
	"fmt"
	"sync"
	"testing"
)

// buat function untuk menggunakan sync.Map
func TestSyncMap(t *testing.T) {
	wg := &sync.WaitGroup{} // buat variabel waitgroup untuk menunggu dan memastikan semua goroutine selesai dijalankan
	response := &sync.Map{} // buat variabel Map yang nanti bisa kita tambahkan datanya, dan pakai datanya

	mtx := &sync.Mutex{}

	for i := 0; i < 100; i++ { // buat looping 100 kali perulangan
		wg.Add(1)   // beru tahu waitgroup bahwa setiap perulangan akan bertambah 1 goroutine yang dijalankan
		go func() { // jalankan goroutine di setiap perulangan, jadi nanti bakal ada 100 goroutine yang dijalankan
			defer wg.Done() // di akhir sebelum goroutine selesai, beri tahu waitgroup bahwa goroutine ini sudah selesai dijalankan
			defer mtx.Unlock()

			mtx.Lock()
			response.Store(i, i) // isi key da value ke variabel Map. key '1', valuenya '1'. begitu seterusnya sampai 100 key
		}()
	}

	wg.Wait() // tunggu hingga waitgroup mendapat sinyal bahwa goroutine sudah selesai dijalankan semua

	response.Range(func(key, value any) bool { // looping semua data yang ada di dalam variabel Map response
		fmt.Println("key:", key, ". valuenya:", value) // tampilkan key dan value di setiap perulangan
		return true                                    // returnkan true jika ingin lanjut ke perulangan berikutnya, false jika ingin stop/keluar perulangan
	})

	key1, _ := response.Load(81)
	fmt.Println(key1)
	fmt.Println("- selesai -")
}
