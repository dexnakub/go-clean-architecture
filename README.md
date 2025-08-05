# 🚀 วิธีการรันโปรเจกต์ #

## 1. ติดตั้ง dependencies ##
```sh
go mod tidy
```
##  2. รันโปรเจกต์ ##
```sh
go run main.go
```
ระบบจะเชื่อมต่อฐานข้อมูลและทำ AutoMigrate ตารางให้โดยอัตโนมัติ

## 3. เข้าใช้งาน API ##
เข้า http://localhost:8080

## 4. Collection PostMan  ##
ไฟล์สำหรับทดสอบ API ด้วย Postman ถูกจัดเก็บไว้ในโฟลเดอร์ `postman/` ที่ root ของโปรเจกต์


# 🧾 API Documentation (API & Database Design Reference)

🔗 [API & SQL Schema Design Document](https://docs.google.com/spreadsheets/d/1Vn6F9ymwzw1One0BZszeplP8DRBTqnV-Fk2nRsAjolE/edit?usp=sharing)


# 🗂️ Project Structure #
cmd/
  └─ api.go                 # จุดเริ่มต้นสำหรับรันเซิร์ฟเวอร์

internal/                 
  ├─ adapters/              # ตัวเชื่อมต่อกับ library ภายนอก (เช่น GORM) หรือ external service (wrapper)
  ├─ app/                   # ส่วนจัดการของ App Layer โดยอิงตาม Framework ปัจจุบันใช้ gin
  │   ├─ handlers/          # ตัวจัดการ HTTP
  │   ├─ helpers/           # Utility Function ที่มีการใช้ร่วมกันหลายๆจุด ของ App Layer
  │   ├─ middleware/        # Function จัดการ request/response ระหว่าง client และ handler
  │   ├─ models/            # Struct Request และ Response ตาม Framework
  │   └─ routes/            # API routes
  ├─ configs/               # Config ต่างๆของ Application
  ├─ data/                  # ส่วนการจัดการข้อมูลต่างๆ (Data Access Layer)
  │   ├─ entities/          # Schema ของ Database
  │   ├─ helpers/           # Utility Function ที่มีการใช้ร่วมกันหลายๆจุด ของ Data Layer
  │   ├─ migrations/        # Function จัดการ Database (Create, Update, Delete)
  │   └─ repositories/      # Function สำหรับจัดการข้อมูลใน Database
  └─ domain/                # ส่วนจัดการ Business Logic
      ├─ models/            # Struct Business Logic
      └─ usecase/           # Business Logic
      
postman/                    # collection postman       

.env                        # Environment variables
go.mod                      # กำหนดโมดูลและจัดการ dependencies ของโปรเจกต์
go.sum                      # ตรวจสอบความถูกต้องของ dependencies
main.go                     # จุดเริ่มต้นโปรแกรม