package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// membuat dan menggunakan channel hanya sebagai parameter In atau Out
func TestChannelInOut(t *testing.T) {
	channel := make(chan string, 1) // buat variabel channel untuk menampung tipe data string
	defer close(channel)            // close channel diakhir sebelum function test selesai dijalankan

	wg := &sync.WaitGroup{} // buat variabel waitgroup untuk menunggu goroutine selesai dikerjakan

	wg.Add(2)               // tambahkan waitgroup 2, sebanyak goroutine yang akan dijalankan
	go OnlyIn(wg, channel)  // jalankan goroutine, di sini channel akan diisi valuenya dengan data string
	go OnlyOut(wg, channel) // jalankan goroutine, di sini channel akan diambil valuenya dan diprint ke console
	wg.Wait()               // tunggu pastikan semua goroutine sudah selesai sebelum program melanjutkan execute syntax di bawahnya

	fmt.Println("program selesai dikerjakan") // tampilkan ke console kalau function test ini sudah selesai dijalankan
}

// buat sebuah function yang menggunakan parameter channel untuk in (diisi dengan value)
func OnlyIn(wg *sync.WaitGroup, channel chan<- string) {
	defer wg.Done() // diakhir sebelum scope function ini selesai, beri tahu bahwa goroutine ini selesai dijalankan

	time.Sleep(1 * time.Second) // pause program selama 1 detik
	channel <- "reo sahobby"    // isi value ke parameter variabel channel
}

// buat sebuah function yang parameter channel digunakan untuk mengoper value (diambil value dari channelnya)
func OnlyOut(wg *sync.WaitGroup, channel <-chan string) {
	defer wg.Done() // diakhir sebelum function ini selesai dijalankan, beri tahu bawah goroutine ini selesai dijalankan

	data := <-channel // ambil value dari channel, simpan ke variabel data
	fmt.Println(data) // tampilkan value dari variabel data
}
