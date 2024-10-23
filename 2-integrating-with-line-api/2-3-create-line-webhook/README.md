# LINE Messaging API Bot Setup

โปรเจ็กต์นี้เป็นตัวอย่างการตั้งค่าและใช้งาน LINE Messaging API Bot ใน Go

## การติดตั้งและการตั้งค่าเบื้องต้น


1. **เริ่มต้นโปรเจ็กต์ Go Module**
   ใช้คำสั่ง `go mod init` เพื่อสร้าง Go module สำหรับโปรเจ็กต์ (ตั้งชื่อ module เป็นชื่อที่เหมาะสม เช่น `my-go-project`)
   ```bash
   go mod init my-go-project
   ```

2. **รับคำสั่งเพื่อ Download Libery ต่างๆลงมาใช้งาน**
   ```bash
   go get github.com/go-resty/resty/v2
   go get github.com/gofiber/fiber/v2
   go get github.com/gofiber/fiber/v2/middleware/logger
   go get github.com/joho/godotenv
   ```

3. **สร้างไฟล์ `.env`**
   
   สร้างไฟล์ `.env` ในไดเรกทอรีหลักของโปรเจ็กต์ และเพิ่มข้อมูลที่จำเป็นดังนี้
   ```bash
   CHANNEL_ID=YOUR_CHANNEL_ID
   CHANNEL_SECRET=YOUR_CHANNEL_SECRET
   PORT=3000  # สามารถระบุพอร์ตอื่นได้
   ```
   - นำ `CHANNEL_ID` และ `CHANNEL_SECRET` ที่ได้จาก LINE Messaging API มาใส่ใน `.env`
   - เลือก PORT สำหรับรันเซิร์ฟเวอร์ (เช่น 3000)

4. **รันโปรเจ็กต์**
   
   เริ่มต้นการรันเซิร์ฟเวอร์โดยใช้คำสั่ง
   ```bash
   go run main.go
   ```

5. **รัน Ngrok**

   ใช้ Ngrok เพื่อลิงก์เซิร์ฟเวอร์ของคุณกับ URL ภายนอก
   ```bash
   ngrok http 3000
   ```
6. **นำ Endpoint ไปใช้ใน LINE Messaging API**

    - เมื่อ Ngrok ทำงานแล้ว จะได้ Endpoint ที่มีรูปแบบ https://xxxxxxxx.ngrok.io
    - นำ Endpoint ที่ได้ไปตั้งค่าใน LINE Developers Console โดยไปที่ Messaging API > Webhook URL และใส่ Endpoint นี้