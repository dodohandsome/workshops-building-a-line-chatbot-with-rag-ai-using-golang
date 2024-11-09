package main

import (
	"fmt"
)

// ฟังก์ชันที่รับ pointer เป็นพารามิเตอร์ เพื่อแก้ไขค่าของตัวแปรโดยตรง
func incrementByPointer(num *int) {
	*num = *num + 1 // เพิ่มค่าโดยการเข้าถึงค่าใน pointer
}

func main() {
	number := 5
	fmt.Println("Before increment:", number) // พิมพ์ค่าก่อนเพิ่ม

	incrementByPointer(&number) // ส่ง pointer ของ number ให้ฟังก์ชัน

	fmt.Println("After increment:", number) // พิมพ์ค่าหลังเพิ่ม
}
