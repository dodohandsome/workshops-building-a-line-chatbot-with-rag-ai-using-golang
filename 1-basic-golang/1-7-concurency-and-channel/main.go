package main

import (
	"fmt"
	"time"
)

func main() {
	// เรียกใช้ printNumbers แบบ goroutine
	go printNumbers()

	// เรียกใช้ printLetters แบบ goroutine
	go printLetters()

	time.Sleep(4 * time.Second) // รอ 4 วินาที
	fmt.Println("Finished!")

	dataChannel := make(chan string)

	// เรียกใช้ goroutines สำหรับดึงข้อมูลจากหลายแหล่งพร้อมกัน
	go fetchData("Source 1", dataChannel)
	go fetchData("Source 2", dataChannel)
	go fetchData("Source 3", dataChannel)

	// รวบรวมผลลัพธ์จากทุก goroutine
	for i := 0; i < 3; i++ {
		fmt.Println(<-dataChannel)
	}

	fmt.Println("All data fetched!")
}

// ฟังก์ชันที่จะถูกเรียกใช้ใน Go routine
// ฟังก์ชันที่พิมพ์ตัวเลขตั้งแต่ 1-5 โดยมีดีเลย์ระหว่างการพิมพ์แต่ละตัวเลข
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(500 * time.Millisecond)
	}
}

// ฟังก์ชันที่พิมพ์ตัวอักษร A ถึง E โดยมีดีเลย์ระหว่างการพิมพ์แต่ละตัวอักษร
func printLetters() {
	for _, letter := range []string{"A", "B", "C", "D", "E"} {
		fmt.Println("Letter:", letter)
		time.Sleep(700 * time.Millisecond)
	}
}

// ฟังก์ชันจำลองดึงข้อมูลจากหลายแหล่ง
func fetchData(source string, ch chan string) {
	time.Sleep(1 * time.Second) // จำลองเวลาในการดึงข้อมูล
	ch <- fmt.Sprintf("Data from %s", source)
}
