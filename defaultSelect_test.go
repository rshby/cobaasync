package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDefaultSelect(t *testing.T) {
	channel1 := make(chan string, 1) // untuk membuat channel yang menampung value string, capacity 1 value
	channel2 := make(chan string, 2) // untuk membuat channel yang menampung value string, capacity 1 value

	defer func() { // defer akan menjalankan ini sebelum fuction test selesai dijalankan
		close(channel2) // close channel supaya tidak terjadi goroutine leaks
		close(channel1) // close channel supaya tidak terjadi goroutine leaks
	}()

	go GiveMeResponse(channel1) // jalankan goroutine, channel1 akan diassign dengan sebuah value string
	go GiveMeResponse(channel2) // jalankan goroutine, channel2 akan diassign dengan sebuah value string

	for { // looping terus menerus
		select { // select untuk memilih salah satu case yang memenuhi
		case data1 := <-channel1: // case jika value channel1 sudah lebih dulu diassign dengan string, simpan ke variabel data1
			fmt.Println("channel1 sudah selesai") // tampilkan informasi bawha channel 1 sudah selesai
			fmt.Println("datanya adalah:", data1) // tampilkan value yang ada di variabel data1
			return                                // stop perulangan kalau sudah ketemu datanya, supaya tidak looping terus meneru
		case data2 := <-channel2: // in case jika channel2 sudah lebih dulu diassign dengan value string, simpan ke dalam variabel data2
			fmt.Println("channel2 sudah selesai") // tampilkan informasi bahwa channel2 sudah selesai
			fmt.Println("datanya adalah:", data2) // tampilkan value dari variabel data2
			return                                // stop perulangan kalau sudah ketemu datanya, supaya tidak looping terus meneru
		default: // jika semua kondisi case tidak ada yang terpenuhi
			fmt.Println("nungguin ya..")       // tampilkan informasi bahwa program masih menunggu case 1 atau case 2 dipilih
			time.Sleep(500 * time.Millisecond) // program akan pause selala setengah detik
		}
	}
}
