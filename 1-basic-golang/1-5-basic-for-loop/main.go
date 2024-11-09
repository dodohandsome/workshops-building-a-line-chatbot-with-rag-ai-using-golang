package main

import (
	"fmt"
)

func main() {
	// 1. Basic for loop (วนลูปแบบทั่วไป)
	fmt.Println("Basic for loop:")
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// 2. for loop แบบเงื่อนไข (เหมือน while loop)
	fmt.Println("\nFor loop with condition:")
	count := 1
	for count <= 3 {
		fmt.Println("Count:", count)
		count++
	}

	// 3. Infinite loop (ลูปที่ไม่มีที่สิ้นสุด - ระวังใช้!)
	fmt.Println("\nInfinite loop (breaking after 3 iterations):")
	counter := 1
	for {
		fmt.Println("Counter:", counter)
		counter++
		if counter > 3 {
			break // ใช้ break เพื่อออกจากลูป
		}
	}

	// 4. Using for-each to loop through a slice (วนลูปผ่าน slice ด้วย range)
	fmt.Println("\nFor-each loop with range:")
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index %d, Value %d\n", index, value)
	}

	// 5. Loop with continue statement (ข้ามบางเงื่อนไขด้วย continue)
	fmt.Println("\nLoop with continue statement (skip even numbers):")
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			continue // ข้ามการแสดงผลหาก i เป็นเลขคู่
		}
		fmt.Println("Odd number:", i)
	}
}
