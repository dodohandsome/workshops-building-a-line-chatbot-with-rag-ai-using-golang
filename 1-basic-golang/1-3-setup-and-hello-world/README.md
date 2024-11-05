## Go Project Setup & Hello World

### 1. สร้างโฟลเดอร์สำหรับโปรเจ็กต์
1. เปิด VS และไปยัง directory ที่ต้องการสร้างโปรเจ็กต์ใหม่

### 2. เริ่มต้นโปรเจ็กต์ Go Module
1. ใช้คำสั่ง `go mod init` เพื่อสร้าง Go module สำหรับโปรเจ็กต์ (ตั้งชื่อ module เป็นชื่อที่เหมาะสม เช่น `my-go-project`)
```bash
go mod init hello-world
```

### 3. เขียนโปรแกรม Hello World
1. สร้างไฟล์ `main.go` ในโฟลเดอร์โปรเจ็กต์
```bash
touch main.go
```
2. เปิดไฟล์ `main.go` และเขียนโค้ด Go สำหรับแสดงผล "Hello, World!"
```bash
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### 4. รันโปรแกรม Hello World
1. รันคำสั่ง `go run` เพื่อรันไฟล์ `main.go`
```bash
go run main.go
```
2. หากทุกอย่างถูกต้อง คุณจะเห็นข้อความ "Hello, World!" แสดงใน Terminal
```bash
Hello, World!
```