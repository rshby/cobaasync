package main

import (
	"fmt"
	"testing"
	"time"
)

// function untuk menjalankan looping channel dengan for range
func TestForRangeChannel(t *testing.T) {
	chanResult := make(chan int) // buat variabel channel untuk menampung value int

	go KirimChannel(chanResult) // jalankan goroutine. di dalan goroutine ini channel akan diassing dengan value

	for x := range chanResult { // looping selama channel chanResult belum diclose
		time.Sleep(1 * time.Second) // program akan pause selama 1 second
		fmt.Println("hasil :", x)   // tampilkan value yang diterima dari channel
	}
}

// buat function untuk mengirim data ke channel
func KirimChannel(chanResult chan int) {
	defer close(chanResult) // diakhri sebelum function selesai dijalankan, close channel supaya for range yang ada di goroutine lain bisa stop

	for i := 0; i < 10; i++ { // buat looping dari nilai 0 sampai 9, 10 kali perulangan
		chanResult <- i + 1 // assign channel dengan value nilai i (bakalan diassign di setiap perulangan)
	}
}
