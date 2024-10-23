## Install Golang

### For Mac

1. Open home directory
```bash
cd ~
```

2. Install Golang on Mac using Homebrew
```bash
brew update
brew install go
```

3. Set environment variables
```bash
echo "export GOROOT=$(brew --prefix golang)/libexec" >> ~/.zshrc
echo "export PATH=$GOROOT/bin:$PATH" >> ~/.zshrc
source ~/.zshrc
```

4. After install, open terminal and run command
```bash
go version
```

### For Windows
1. Download Golang

- ไปที่ https://golang.org/dl/ และดาวน์โหลดไฟล์ .msi สำหรับ Windows (เวอร์ชันล่าสุด เช่น go1.21.1.windows-amd64.msi)

2. Install Golang

- ดับเบิลคลิกที่ไฟล์ .msi ที่ดาวน์โหลดมา และทำตามขั้นตอนใน wizard เพื่อติดตั้ง

3. Set environment variables

- เปิด **Control Panel > System and Security > System > Advanced system settings**
- คลิกที่ **Environment Variables**
- ใน **System Variables** ค้นหา `Path` และคลิก **Edit**
- เพิ่ม `C:\Go\bin` ใน `Path` แล้วคลิก OK

4. Open Command Prompt or PowerShell and run command
```bash
go version
```