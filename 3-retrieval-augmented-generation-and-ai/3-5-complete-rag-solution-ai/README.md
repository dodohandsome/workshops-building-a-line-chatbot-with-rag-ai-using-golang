# Project Setup Guide

## ขั้นตอนการตั้งค่าโปรเจกต์


### 1. เริ่มต้นโปรเจ็กต์ Go Module
ใช้คำสั่ง `go mod init` เพื่อสร้าง Go module สำหรับโปรเจ็กต์ (ตั้งชื่อ module เป็นชื่อที่เหมาะสม เช่น `my-go-project`)
```bash
go mod init my-go-project
```

### ขั้นที่ 2: ติดตั้ง Dependencies
รันคำสั่งต่อไปนี้ใน terminal เพื่อดาวน์โหลด dependencies ที่จำเป็น:

   ```bash
   go mod tidy
   ```

### ขั้นที่ 3: ตั้งค่าไฟล์ .env

สร้างไฟล์ `.env` ใน root directory ของโปรเจกต์ จากนั้นใส่ค่าต่อไปนี้:

```plaintext
OPENAI_API_KEY=YOUR_OPENAI_API_KEY
CHANNEL_ID=YOUR_CHANNEL_ID
CHANNEL_SECRET=YOUR_CHANNEL_SECRET
PORT=YOUR_PORT
INDEX_NAME=YOUR_INDEX_NAME
PINECONE_API_KEY=YOUR_PINECONE_API_KEY
```
หมายเหตุ: กรุณาแทนที่ `YOUR_OPENAI_API_KEY`, `YOUR_CHANNEL_ID`, `YOUR_CHANNEL_SECRET`, `YOUR_PORT`, `YOUR_INDEX_NAME`, และ `YOUR_PINECONE_API_KEY` ด้วยข้อมูลที่ได้รับจริง

### ขั้นที่ 4: เริ่มรันโปรเจกต์
หลังจากตั้งค่าทุกอย่างแล้ว คุณสามารถเริ่มรันโปรเจกต์ได้ด้วยคำสั่งต่อไปนี้:

  ```bash
   go run .
   ```
