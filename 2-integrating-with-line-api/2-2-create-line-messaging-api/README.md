# LINE Official Account & Messaging API Setup Guide

คู่มือการตั้งค่า LINE Official Account (LINE OA) และการเชื่อมต่อ LINE Messaging API เพื่อเริ่มพัฒนาแชทบอท

## ขั้นตอนที่ 1: สร้าง LINE Official Account

1. **เข้าสู่ LINE Official Account Manager**: 
   - ไปที่ [LINE Official Account Manager](https://manager.line.biz/) และเข้าสู่ระบบด้วยบัญชี LINE ของคุณ

2. **สร้าง Official Account ใหม่**:
   - กดปุ่ม `Create Account` (สร้างบัญชี) เลือก `LINE Official Account`
   - กรอกข้อมูลที่จำเป็น เช่น ชื่อและหมวดหมู่บัญชี
   - หลังจากสร้างเสร็จจะเข้าสู่หน้าแดชบอร์ด LINE Official Account ของคุณ

3. **สร้าง LINE Messaging API**:
   - ในหน้าแดชบอร์ดของ LINE Official Account ไปที่ `Settings` > `Account Details` > `Messaging API` > `Enable Messaging API`
   - สร้าง New Provider หรือถ้ามีอยู่แล้วเลือกที่มีอยู่แล้วได้เลย หลักจากนั้นกด OK ไปจดสุด
   - กรณีที่ยังไม่เคยสมัครเป็น LINE Developer จะมีให้กรอก `NAME` และ `EMAIL` แค่นี้เลยหลังจากนั้นก็สร้าง `New Provider` ได้เลย
 

## ขั้นตอนที่ 2: หลักจากสร้าง LINE Messaging API


1. **เข้าสู่ LINE Developers Console**:
   - ไปที่ [LINE Developers Console](https://developers.line.biz/) และเข้าสู่ระบบด้วยบัญชี LINE เดียวกัน
   - หรืออีกวิธีหลังจาก  `Enable Messaging API` เรียบร้อยจะมีข้อความนี้แสดงขึ้น You can find more related settings in `LINE Developers`. ให้กดที่ `LINE Developers` ก็ได้เช่นกัน

2. **บันทึก Channel Credentials**:
   - หลังจากสร้าง Channel เสร็จสิ้น ให้ไปที่ `Channel Settings` เพื่อบันทึกข้อมูลสำคัญ:
     - `Channel ID`
     - `Channel Secret`
     - `Channel Access Token`

## ขั้นตอนที่ 3: เชื่อมต่อ LINE Messaging API กับ LINE OA

1. **ตั้งค่า Webhook URL**:
   - ใน `Messaging API` ของ LINE Developers Console ไปที่ `Webhook settings`
   - ตั้งค่า `Webhook URL` ให้ชี้ไปยังเซิร์ฟเวอร์ของคุณ (เช่น `https://yourdomain.com/webhook`)
   - เปิดใช้งาน `Use Webhook`

2. **ตั้งค่า Webhook Event**:
   - ในส่วน `Bot settings` ให้เปิดใช้งาน `Auto-reply messages` และ `Greeting messages` เพื่อทดสอบการทำงานของบอทเบื้องต้น

3. **ทดสอบการเชื่อมต่อ**:
   - ใช้ LINE App สแกน `QR Code` จาก `Messaging API` Console เพื่อเพิ่มบัญชี Official Account ของคุณและทดสอบการส่งข้อความ
