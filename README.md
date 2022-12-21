# Go2-Todo

# CMD Golang
- install : go get -u gorm.io/driver/sqlite
- add module requirements and sums : go mod tidy
- PORT 8081 go run main.go : change PORT



# Request API
- Use extendtion in VS : REST Client
- file .http
- JWT : ใช้ทำ Authentication ก่อนเข้าถึง resource
>>> have 3 part : 
1. Headers (Algorytm encryption) 
2. Payloads (claim data) 
3. Signature (ลายเซ็น)
>>> การตรวจสอบ Token
1. token หมดอายุหรือยัง
2. signature เป็นของที่เรา gen หรือไม่
3. Audian คือใคร Trusted ไหม

# Database
- Use extendtion in VS
- Create connect with file test.db


# Command
# Build with parameter --ldflags
1. go build -o app --ldflags "-X main.buildCommit=`git rev-parse --short HEAD` -X main.buildTime=`date "+%Y-%m-%dT%H:%M:%S%Z:00"`"
2. app
