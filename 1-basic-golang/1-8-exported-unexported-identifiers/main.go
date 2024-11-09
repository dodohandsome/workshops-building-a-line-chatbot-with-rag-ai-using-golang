package main

import "exported-unexported-identifiers/test"

func main() {
	// Go จะให้เราใช้ defer ในการทํางานของฟังก์ชันที่ต้องการให้ทำงานหลังสุด
	defer test.TestPrintln("Hello from main function 1!")
	defer test.TestPrintln("Hello from main function 2!")
	defer test.TestPrintln("Hello from main function 3!")
	test.TestPrintln("Hello from main function line 4!")
}
