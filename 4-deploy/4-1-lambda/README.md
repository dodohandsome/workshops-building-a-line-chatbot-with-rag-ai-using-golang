# AWS Lambda Deployment Guide

## Step 1: Install AWS CLI

AWS CLI เป็นเครื่องมือที่ใช้ในการควบคุมบริการของ AWS ผ่านบรรทัดคำสั่ง เราจำเป็นต้องติดตั้งและตั้งค่าก่อนที่จะ Deploy Lambda Function ได้

1. **Download and Install AWS CLI**

   สำหรับ Windows, macOS, และ Linux ให้ดาวน์โหลดได้ที่: https://aws.amazon.com/cli/

   สำหรับ macOS คุณสามารถติดตั้งผ่าน `Homebrew` ได้เช่นกัน:

   ```bash
   brew install awscli
   ```
    ตรวจสอบว่า AWS CLI ติดตั้งสำเร็จแล้วด้วยคำสั่ง:
   ```bash
   aws --version
   ```

2. **Configure AWS CLI**
   รันคำสั่ง `aws configure `เพื่อทำการตั้งค่าการเข้าถึง AWS
   ```bash
   aws configure
   ```
    จากนั้นให้ใส่ข้อมูล:

    - `AWS Access Key ID`: ใส่ Access Key ของคุณ
    - `AWS Secret Access Key`: ใส่ Secret Key ของคุณ
    - `Default region name`: ใส่ Region ที่คุณต้องการ (เช่น `ap-southeast-1`)
    `Default output format`: เลือกเป็น json
    ข้อมูลเหล่านี้สามารถสร้างได้ที่ AWS IAM Console ในส่วนของ Access Management > Users


## Step 2: Install Serverless Framework

Serverless Framework ทำให้การ Deploy Lambda ง่ายขึ้น

1. **Install Node.js** (ถ้ายังไม่มีติดตั้ง)

   คุณสามารถดาวน์โหลดได้ที่ [Node.js](https://nodejs.org/)

2. **Install Serverless Framework**

   ติดตั้ง Serverless Framework ผ่าน npm:

   ```bash
   npm install -g serverless
   ```
3. **Check Serverless**
    ```bash
   serverless -v
   ```

## Step 3: Deploy Lambda
1. สร้างไฟล์ `.env` ใน root directory ของโปรเจกต์ จากนั้นใส่ค่าต่อไปนี้:

    ```plaintext
    USE_LAMBDA=TRUE
    OPENAI_API_KEY=YOUR_OPENAI_API_KEY
    CHANNEL_ID=YOUR_CHANNEL_ID
    CHANNEL_SECRET=YOUR_CHANNEL_SECRET
    PORT=YOUR_PORT
    INDEX_NAME=YOUR_INDEX_NAME
    PINECONE_API_KEY=YOUR_PINECONE_API_KEY
    ```
    หมายเหตุ: กรุณาแทนที่ `YOUR_OPENAI_API_KEY`, `YOUR_CHANNEL_ID`, `YOUR_CHANNEL_SECRET`, `YOUR_PORT`, `YOUR_INDEX_NAME`, และ `YOUR_PINECONE_API_KEY` ด้วยข้อมูลที่ได้รับจริง

### ขั้นที่ 2: ติดตั้ง Dependencies
รันคำสั่งต่อไปนี้ใน terminal เพื่อดาวน์โหลด dependencies ที่จำเป็น:

   ```bash
   go mod tidy
   ```
### ขั้นที่ 3: Deploy Lambda Function
   
   ใช้คำสั่ง `serverless deploy` เพื่อ Deploy Lambda Function ของคุณไปยัง AWS
   ```bash
   sh deploy.sh
   ```