package main

import (
	"fmt"
)

type Profile struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	// การประกาศตัวแปรแบบกำหนดชนิดข้อมูล
	var name string = "Go Language"
	var age int = 25

	// การประกาศตัวแปรแบบไม่กำหนดค่าเริ่มต้น
	var score int // ค่าเริ่มต้นจะเป็น 0
	fmt.Println("Score:", score)

	// การกำหนดตัวแปรแบบสั้น (Short Variable Declaration)
	country := "Thailand"
	isStudent := true
	gpa := 3.5

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Country:", country)
	fmt.Println("Is Student:", isStudent)
	fmt.Println("GPA:", gpa)

	// การใช้งานค่าคงที่ (Constants)
	const pi = 3.14159
	const language = "Go"
	fmt.Println("Pi:", pi)
	fmt.Println("Programming Language:", language)

	// การแปลงชนิดข้อมูล (Type Conversion)
	var days int = 7
	var daysInFloat float64 = float64(days) // แปลงจาก int เป็น float64
	fmt.Println("Days in int:", days)
	fmt.Println("Days in float64:", daysInFloat)

	// แสดงค่าเริ่มต้นของชนิดข้อมูลต่าง ๆ
	var defaultInt int
	var defaultFloat float64
	var defaultBool bool
	var defaultString string

	fmt.Println("Default int:", defaultInt)
	fmt.Println("Default float64:", defaultFloat)
	fmt.Println("Default bool:", defaultBool)
	fmt.Println("Default string:", defaultString)

	// การประกาศตัวแปรแบบมีโครงสร้าง
	var profile Profile
	profile.Firstname = "thamrong"
	profile.Lastname = "chaiwong"
	profile.Age = 14

	profile2 := Profile{
		Firstname: "thamrong",
		Lastname:  "chaiwong",
		Age:       14,
	}

	fmt.Println("Profile:", profile)
	fmt.Println("Profile2:", profile2)
}
