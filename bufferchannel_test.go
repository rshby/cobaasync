package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// function untuk membuat dan menggunakan buffer channel
func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 2) // membuat channel untuk menampung data string, dengan buffer capacity 2 value
	defer close(channel)            // diakhir program sebelum function selesai, close channel supapa tidak goroutine leak (optimasi memory)
	fmt.Println("isi data pertama ke channel")
	channel <- "reo"                                   // assign value ke channel, akan memakan buffer capacity 1 (masih sisa 1 value yang bisa masuk)
	fmt.Println("len channel menjadi :", len(channel)) // tampilkan informasi panjang (berapa banyak value yang ada di dalam channel)

	fmt.Println("value pertama channel :", <-channel)  // tampilkan isi value pertama yang ada di channel (slot buffer pertama akan menjadi kosong)
	fmt.Println("len channel menjadi :", len(channel)) // tampilkan panjang channel (berapa banyak value yang ada di channel)

	channel <- "sahobby" // asssign value ke channel, memakan buffer capacity 1 (sudah tidak ada tempat untuk value masuk lagi)

	fmt.Println("panjang channel adalah :", len(channel))  // tampilkan informasi panjang channel (berapa banyak value yang ada di channel)
	fmt.Println("capacity channel adalah :", cap(channel)) // tampilkan informasi capacity channel (berapa maksimal value yang bisa masuk ke channel)
}

// function untuk membuat buffer channel dan wait group
func TestBufferChannelWait(t *testing.T) {
	chanResult := make(chan string, 2) // buat variabel channel untuk menampung value string, dengan capacity 2 value
	defer close(chanResult)            // diakhir program sebelum function selesai, close channel (supaya tidak goroutine leaks)

	wg := &sync.WaitGroup{} // buat variabel waitgroup untuk menunggu

	wg.Add(2)                                // informasikan ada 2 goroutine yang siap ditunggu
	go IsiChannel(wg, "reo", chanResult)     // jalankan goroutine (akan mengisi channel chanResult dengan sebuah value string)
	go IsiChannel(wg, "sahobby", chanResult) // jalankan goruutine(akan mengisi channel chanResult dengan sebuah value string)
	wg.Wait()                                // tunggu semua goroutine yang jalan di thread lain selesai semua

	fmt.Println(<-chanResult)                             // tampilkan value dari channel (len akan berkurang satu, karena value sudah diambil/diconsume)
	fmt.Println("len channel menjadi :", len(chanResult)) // tampilkan panjang channel

	fmt.Println(<-chanResult)                             // tampilkan value dari channel (len akan berkurang satu)
	fmt.Println("len channel menjadi :", len(chanResult)) // tampilkan panjang channel (sehrusnya 0 karena sudah diconsume semua valuenya)
}

// function yang digunakan untuk mengisi channel dengan value
func IsiChannel(wg *sync.WaitGroup, inputName string, chanResult chan<- string) {
	defer wg.Done() // diakhir sebelum function selesai, beri tahu kalau goroutine ini sudah selesai dijalankan

	time.Sleep(2 * time.Second) // goroutine ini akan pause selama 2 detik
	chanResult <- inputName     // isi value dari inputName ke dalam variabel channel
}

func TestBufferChannelLagi(t *testing.T) {
	channel := make(chan int, 1)

	go func() {
		channel <- 10 // pas di sini harus ada gorouine lain yang merima
	}()

	//fmt.Println(nyValue)
	fmt.Println("selesai")
}
