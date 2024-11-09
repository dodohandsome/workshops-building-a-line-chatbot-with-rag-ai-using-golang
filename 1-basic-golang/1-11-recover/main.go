package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting program")

	// ใช้ defer-recover เพื่อจัดการกับ panic ที่อาจเกิดขึ้นในฟังก์ชันนี้
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("panic handled: %v\n", e)
		}
	}()

	// เรียกใช้ฟังก์ชันที่อาจทำให้เกิด panic
	causePanic()

	fmt.Println("Program continues after handling panic")
}

// ฟังก์ชันที่จงใจทำให้เกิด panic
func causePanic() {
	fmt.Println("About to cause panic")
	panic("something went wrong!") // ทำให้เกิด panic โดยจงใจ
}
