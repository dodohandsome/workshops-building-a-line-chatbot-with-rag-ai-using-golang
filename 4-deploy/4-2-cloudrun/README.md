# Cloud Run Deployment Guide

## การตั้งค่าและการ Deploy บน Google Cloud Run

1. **เปลี่ยนค่าคีย์ใน `deploy.sh` และ `cloud-build.yaml`:**
   - ในไฟล์ `deploy.sh` และ `cloud-build.yaml` ให้ทำการแก้ไข Project ID เป็นของคุณเอง
   - สำหรับ `cloud-build.yaml`:
     - ค้นหา `_PROJECT_ID` และแทนที่ด้วย Project ID ของคุณ เช่น `"your-project-id"`
   - สำหรับ `deploy.sh`:
     - แก้ไขค่าของตัวแปรที่เกี่ยวข้อง เช่น `PROJECT_ID="your-project-id"`

2. **Deploy บน Cloud Run:**
   สามารถดูรายละเอียดและวิธีการ Deploy บน Cloud Run ได้ที่บทความนี้:  
   [ดูวิธีการ Deploy บน Cloud Run](https://medium.com/p/98dda588406c)

---

หลังจากปรับคอนฟิกเสร็จเรียบร้อย ให้ทำการรันสคริปต์ `deploy.sh` เพื่อทำการ Deploy ไปยัง Cloud Run 
