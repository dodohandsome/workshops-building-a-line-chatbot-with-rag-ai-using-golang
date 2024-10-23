package main

import (
	"errors"
	"fmt"
	"time"
)

type Profile struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

// ใช้ interface แทนการกำหนดความสามารถที่ struct นั้นๆ ควรมี (คล้ายกับการกำหนด method ในคลาส)
type IProfileService interface {
	Validate(c *Profile) bool
	CreateProfile(c *Profile) error
	GetProfile(name string) (*Profile, error)
}

// ใช้ struct แทนคลาสเพื่อเก็บข้อมูล
type profileService struct {
	defaultAge int                 // เก็บค่า defaultAge ที่กำหนด
	storage    map[string]*Profile // ใช้เก็บข้อมูล Profile ใน map ที่ทำหน้าที่เหมือนฐานข้อมูลเล็ก ๆ
}

// เป็นการประกอบ IProfileService และ profileService เข้าด้วยกัน เรียกว่า Composition
func NewProfileService(defaultAge int) IProfileService {
	return &profileService{
		defaultAge: defaultAge,
		storage:    make(map[string]*Profile),
	}
}

func main() {
	// Go จะให้เราใช้ defer ในการทํางานของฟังก์ชันที่ต้องการให้ทำงานหลังสุด
	// defer test.TestPrintln("Hello from main function line 40!")
	// defer test.TestPrintln("Hello from main function line 41!")
	// defer test.TestPrintln("Hello from main function line 42!")

	// Go จะให้เราใช้ var ในการประกาศตัวแปร โดยกำหนดชื่อและประเภทข้อมูลของตัวแปรนั้น
	// var name string = "Dodo" // ประกาศตัวแปร name ที่เป็นประเภท string
	// var age int = 25         // ประกาศตัวแปร age ที่เป็นประเภท int (จำนวนเต็ม)
	// message := "สวัสดี " + name
	// message2 := fmt.Sprintf("สวัสดี %s", name)
	// fmt.Println(message)
	// fmt.Println(message2)
	// เราสามารถประกาศตัวแปรแบบ shorthand ได้ด้วย :=
	// ซึ่งสามารถใช้ได้เฉพาะการประกาศตัวแปรครั้งแรกในฟังก์ชันนั้นๆ
	// country := "Thailand"

	// fmt.Println("Name:", name)
	// fmt.Println("Age:", age)
	// fmt.Println("Country:", country)

	// Go เป็นภาษาที่ใช้ static typing ซึ่งหมายความว่าแต่ละตัวแปรต้องระบุประเภทข้อมูลอย่างชัดเจน
	// var isStudent bool = true        // ตัวแปรประเภท boolean
	// var height float64 = 5.9         // ตัวแปรประเภท float64
	// var grades []int = []int{90, 85} // ตัวแปรประเภท slice ของ int (คล้าย list)

	// fmt.Println("Is Student:", isStudent)
	// fmt.Println("Height:", height)
	// fmt.Println("Grades:", grades)

	// แต่เราสามารถประกาศตัวแปรแบบ interface ได้ด้วย map
	// profileAny := map[string]interface{}{}
	// profileAny["firstname"] = "thamrong"
	// profileAny["lastname"] = "chaiwong"
	// fmt.Println("profileAny", profileAny)
	// js, _ := json.Marshal(profileAny)
	// fmt.Println("JSON : ", string(js))

	// Go ประกาศตัวแปรแบบ pointer
	// profileFirst := &Profile{
	// 	Firstname: "thamrong",
	// 	Lastname:  "chaiwong",
	// }

	// fmt.Println("profileFirst:", profileFirst)
	//เปลี่ยนค่าแบบ pointer
	// showcasePointer(profileFirst)
	// fmt.Println("showcasePointer:", profileFirst)

	// Go ประกาศตัวแปรแบบไม่ pointer
	// profileSec := Profile{
	// 	Firstname: "thamrong",
	// 	Lastname:  "chaiwong",
	// }
	// fmt.Println("profileSec:", profileSec)
	// เปลี่ยนค่าแบบ Copy by Value
	// profileSecResp := showcaseNoPointer(profileSec)
	// fmt.Println("showcaseNoPointer:", profileSec)
	// fmt.Println("profileSecResp:", profileSecResp)

	// Go มีแค่ for loop เท่านั้น ไม่มี while loop
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Looping:", i)
	// }

	// ลูปผ่าน array หรือ slice ก็ทำได้ง่ายๆ
	// names := []string{"item1", "item2", "item3"}
	// for index, name := range names {
	// 	fmt.Printf("Index: %d, Name: %s\n", index, name)
	// }

	// names = append(names, "item4")
	// fmt.Println("names:", names)
	// index := 2
	// names = append(names[:index], names[index+1:]...)
	// fmt.Println("remove names index 2:", names)
	// fmt.Println("len names:", len(names))

	// เป็นการกลายเป็น if-else และ switch case

	// currentHour := time.Now().Hour()

	// if currentHour < 12 {
	// 	fmt.Println("Good morning!")
	// } else if currentHour < 18 {
	// 	fmt.Println("Good afternoon!")
	// } else {
	// 	fmt.Println("Good evening!")
	// }

	// switch {
	// case currentHour < 12:
	// 	fmt.Println("Good morning!")
	// case currentHour < 18:
	// 	fmt.Println("Good afternoon!")
	// default:
	// 	fmt.Println("Good evening!")
	// }

	// day := "Saturday"

	// if day == "Monday" {
	// 	fmt.Println("Today is Monday.")
	// } else if day == "Tuesday" {
	// 	fmt.Println("Today is Tuesday.")
	// } else if day == "Wednesday" {
	// 	fmt.Println("Today is Wednesday.")
	// } else if day == "Thursday" {
	// 	fmt.Println("Today is Thursday.")
	// } else if day == "Friday" {
	// 	fmt.Println("Today is Friday.")
	// } else {
	// 	fmt.Println("It's the weekend!")
	// }

	// switch day {
	// case "Monday":
	// 	fmt.Println("Today is Monday.")
	// case "Tuesday":
	// 	fmt.Println("Today is Tuesday.")
	// case "Wednesday":
	// 	fmt.Println("Today is Wednesday.")
	// case "Thursday":
	// 	fmt.Println("Today is Thursday.")
	// case "Friday":
	// 	fmt.Println("Today is Friday.")
	// default:
	// 	fmt.Println("It's the weekend!")
	// }

	// if day == "Saturday" || day == "Sunday" {
	// 	fmt.Println("It's the weekend!")
	// } else {
	// 	fmt.Println("It's a weekday.")
	// }

	// switch day {
	// case "Saturday", "Sunday":
	// 	fmt.Println("It's the weekend!")
	// default:
	// 	fmt.Println("It's a weekday.")
	// }

	// Go routine เป็นวิธีง่ายๆ ในการทำงานแบบขนาน (concurrent) ใน Go
	// ใช้คำสั่ง go ก่อนฟังก์ชันที่ต้องการรันแบบขนาน
	// เรียกใช้ printNumbers แบบ goroutine
	// go printNumbers()

	// เรียกใช้ printLetters แบบ goroutine
	// go printLetters()

	// ใช้ time.Sleep เพื่อรอให้ Go routine ทำงานเสร็จก่อนโปรแกรมจะจบ
	// time.Sleep(4 * time.Second) // รอ 4 วินาที

	// dataChannel := make(chan string)

	// เรียกใช้ goroutines สำหรับดึงข้อมูลจากหลายแหล่งพร้อมกัน
	// go fetchData("Source 1", dataChannel)
	// go fetchData("Source 2", dataChannel)
	// go fetchData("Source 3", dataChannel)

	// รวบรวมผลลัพธ์จากทุก goroutine
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(<-dataChannel)
	// }

	// fmt.Println("All data fetched!")

	// ถึงแม้ Go จะไม่ได้เป็นภาษา OOP โดยตรง (เพราะไม่มีคลาสหรือการสืบทอดแบบ OOP ทั่วไป) แต่เรายังสามารถใช้ struct และ interface เพื่อออกแบบโค้ดให้ใกล้เคียงกับ OOP ได้
	// เรียกใช้ NewProfileService และกำหนด defaultAge เป็น 18
	// service := NewProfileService(18)

	// สร้างโปรไฟล์ใหม่
	// profile := &Profile{Firstname: "Thamrong", Age: 20}
	// if service.Validate(profile) {
	// 	err := service.CreateProfile(profile)
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 	}
	// }

	// ดึงข้อมูลโปรไฟล์ที่สร้างไว้
	// profileData, err := service.GetProfile("Thamrong")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Retrieved Profile:", profileData)
	// }

	// fmt.Println("Main function finished!")
	// test.TestPrintln("Hello from main function line 228!")
}

// ฟังก์ชันที่จะถูกเรียกใช้ใน Go routine
// ฟังก์ชันที่พิมพ์ตัวเลขตั้งแต่ 1-5 โดยมีดีเลย์ระหว่างการพิมพ์แต่ละตัวเลข
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(500 * time.Millisecond) // รอครึ่งวินาที
	}
}

// ฟังก์ชันที่พิมพ์ตัวอักษร A ถึง E โดยมีดีเลย์ระหว่างการพิมพ์แต่ละตัวอักษร
func printLetters() {
	for _, letter := range []string{"A", "B", "C", "D", "E"} {
		fmt.Println("Letter:", letter)
		time.Sleep(700 * time.Millisecond) // รอเจ็ดในสิบวินาที
	}
}

func showcasePointer(profile *Profile) {
	profile.Firstname = "thamrong2"
}

func showcaseNoPointer(profile Profile) Profile {
	profile.Firstname = "thamrong2"
	return profile
}

func fetchData(source string, ch chan string) {
	time.Sleep(1 * time.Second) // จำลองเวลาในการดึงข้อมูล
	ch <- fmt.Sprintf("Data from %s", source)
}

func (p *profileService) Validate(c *Profile) bool {
	return c.Age >= p.defaultAge
}

// Implement ฟังก์ชัน CreateProfile เพื่อเพิ่มโปรไฟล์ลงใน storage
func (p *profileService) CreateProfile(c *Profile) error {
	if _, exists := p.storage[c.Firstname]; exists {
		return errors.New("profile already exists")
	}
	p.storage[c.Firstname] = c
	fmt.Println("Profile created:", c.Firstname)
	return nil
}

// Implement ฟังก์ชัน GetProfile เพื่อดึงข้อมูลโปรไฟล์จาก storage
func (p *profileService) GetProfile(name string) (*Profile, error) {
	profile, exists := p.storage[name]
	if !exists {
		return nil, errors.New("profile not found")
	}
	return profile, nil
}
