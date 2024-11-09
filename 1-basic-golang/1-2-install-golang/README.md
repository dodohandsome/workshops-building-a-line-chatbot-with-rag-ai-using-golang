## ติดตั้ง Golang บน Mac และ Windows

### For Mac

### วิธีที่ 1: ติดตั้งผ่านแพ็กเกจ Installer จากเว็บไซต์ทางการ
1. ดาวน์โหลด Go:
   - ไปที่ เว็บไซต์ดาวน์โหลด Go https://golang.org/dl/
   - ดาวน์โหลดไฟล์ติดตั้งสำหรับ macOS ที่มีนามสกุล .pkg
2. ติดตั้ง Go:
   - ดับเบิลคลิกไฟล์ .pkg ที่ดาวน์โหลดมา
   - ทำตามคำแนะนำบนหน้าจอ โดยคลิก “Continue” และ “Install”
   - ใส่รหัสผ่านผู้ดูแลระบบ (ถ้าจำเป็น) เพื่อยืนยันการติดตั้ง
3. ตรวจสอบการติดตั้ง:
   - เปิด Terminal
   - พิมพ์คำสั่ง:
        ```bash
        go version
        ```
   - ควรเห็นข้อความแสดงเวอร์ชันของ Go ที่ติดตั้ง

### วิธีที่ 2: ติดตั้งผ่าน Homebrew
1. ติดตั้ง Homebrew (ถ้ายังไม่มี):
   - เปิด Terminal
   - พิมพ์คำสั่ง:
        ```bash
        /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
        ```
2. ติดตั้ง Go ด้วย Homebrew:
   - brew update
   - brew install go
3. ตรวจสอบการติดตั้ง:
   - พิมพ์:
        ```bash
        go version
        ```
   - ควรเห็นเวอร์ชันของ Go ที่ติดตั้ง

### การตั้งค่า GOPATH และ PATH
1. เปิดไฟล์การตั้งค่า Shell:
   - สำหรับ bash:
        ```bash
        nano ~/.bash_profile
        ```
   - สำหรับ zsh (ค่าเริ่มต้นใน macOS ใหม่):
        ```bash
        nano ~/.zshrc
        ```
2. เพิ่มบรรทัดต่อไปนี้:
    ```bash
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOPATH/bin
    ```
3. บันทึกและออกจาก nano:
   - กด Ctrl + O เพื่อบันทึก
   - กด Ctrl + X เพื่อออก
4. โหลดการตั้งค่าใหม่:
    ```bash
    source ~/.bash_profile
    # หรือ
    source ~/.zshrc
    ```

#

### For Window
### ขั้นตอนการติดตั้ง
1. ดาวน์โหลด Go:
   - ไปที่ เว็บไซต์ดาวน์โหลด Go https://golang.org/dl/
   - ดาวน์โหลดไฟล์ติดตั้งสำหรับ macOS ที่มีนามสกุล .msi
2. ติดตั้ง Go:
   - ดับเบิลคลิกไฟล์ .msi ที่ดาวน์โหลดมา
   - ทำตามคำแนะนำบนหน้าจอ โดยคลิก “Next” และยอมรับเงื่อนไขการใช้งาน
   - เลือกตำแหน่งที่ต้องการติดตั้ง (ค่าเริ่มต้นคือ C:\Go)
   - คลิก “Install” และรอจนกว่าการติดตั้งจะเสร็จสิ้น
   - คลิก “Finish” เพื่อปิดตัวติดตั้ง
3. ตรวจสอบการติดตั้ง:
   - เปิด Command Prompt หรือ PowerShell
   - พิมพ์คำสั่ง:
        ```bash
        go version
        ```
   - ควรเห็นข้อความแสดงเวอร์ชันของ Go ที่ติดตั้ง
### การตั้งค่า Environment Variables
1. เปิด System Properties:
   - กดปุ่ม Windows + R เพื่อเปิดหน้าต่าง Run
   - พิมพ์ sysdm.cpl แล้วกด Enter
   - หรือ เปิด **Control Panel > System and Security > System > Advanced system settings**
2. แก้ไขตัวแปรระบบ:
   - ในแท็บ “Advanced” คลิกที่ “Environment Variables”
   - ในส่วน “System variables” เลื่อนหาและเลือก Path แล้วคลิก “Edit”
3. เพิ่ม Go ใน Path:
   - คลิก “New” แล้วเพิ่ม:
        ```bash
        C:\Go\bin
        ```
   - คลิก “OK” เพื่อบันทึกและปิดหน้าต่างทั้งหมด
4. รีสตาร์ท Command Prompt/PowerShell เพื่อให้การตั้งค่ามีผล



#
### Extensions

ติดตั้ง Extentions Go ให้เรียบร้อย

<p align="center" width="100%">
    <img src="../../assets/extensions-go.png"> 
</p>


