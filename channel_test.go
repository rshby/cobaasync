package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	// membuat channel dengan tipe data string
	chanString := make(chan string, 1)

	// menutup/close channel
	defer close(chanString)

	// untuk mengirim data bisa menggunkaan syntax : channel <- data
	myName := "reo sahobby"
	chanString <- myName

	// untuk menerima data dari channel menggunakan syntax : variabel := <-channel
	namaSaya := <-chanString
	fmt.Println(namaSaya)
}

func TestUseChannel(t *testing.T) {
	channel := make(chan string, 1) // membuat channel dengan tipe data string
	defer close(channel)            // nanti akan menjalankan perintah close() diakhir sebelum scope ini selesai dijalankan

	// membuat anonymous function yang akan dijalankan sebagai goroutine
	go func() {
		time.Sleep(2 * time.Second)        // program akan pause salama 2 detik
		channel <- "eko kurniawan khanedy" // isi value channel dengan string seperti di samping kiri
	}()

	fmt.Println("tunggu data dari goroutine...") // tampilkan informasi bahwa program main sedang menunggu channel yang ada di goroutine lain diassign dengan value
	data := <-channel                            // value dari channel akan diterima ke variabel data
	fmt.Println(data)                            // tampilkan value dari variabel data
}

func TestChannelSebagaiParameter(t *testing.T) {
	channel := make(chan string, 1) // buat variabel channel untuk menampung data string
	defer close(channel)            // close channel nanti diakhir program sebelum function selesai dijalankan

	go GiveMeResponse(channel) // jalankan function dengan goroutine

	data := <-channel                      // ambil/terima data dari channel yang diisi di goroutine lain ke dalam variabel data
	fmt.Println("isi data adalah :", data) // tampilkan value dari variabel data
}

// buat function dengan parameter channel string yang akan diisi datanya
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second) // pause program selama 2 detik
	channel <- "reo sahobby"    // isi value channel dengan data string
}

func Test2Channel(t *testing.T) {
	fmt.Println("mulai")

	wg := &sync.WaitGroup{}

	// jalan di thread lain jalan
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done() // di akhir sebelum function selesai dijalankan, kasih informasi bahwa goroutine ini selesai dijalankan
		// baru siap siap mau print
		fmt.Println("ini di dalam goroutine")
	}(wg)

	wg.Wait()
	// udah duluan selesai
	fmt.Println("selesai")
}

func Test2Channell(t *testing.T) {
	chanResult := make(chan string, 1) // buat channel untuk menampung data nasabah
	defer close(chanResult)

	go GetDataNasabahBRI(chanResult)

	//nasabah := <-chanResult // ambil data dari nasabah yang diget (blocking) -----...........
	//fmt.Println(nasabah)
	time.Sleep(2 * time.Second)
	fmt.Println("selesai")
}

func GetDataNasabahBRI(chanNasbahReturn chan string) {
	result := "success get data nasabah bri"
	// simpan hasil ke channel
	chanNasbahReturn <- result
}
