# Golang Simple Service
A simple Golang service with a REST API, using PostgreSQL as the database, and Docker for containerization.

# Notes

- Handler -> Business [-> Repository] -> Storage (tầng trên dùng interface của tầng dưới)
- 2 layer cannot perform unit test: Handler, Storage (only can integration test)
- Business (implement logics) can unit test
- Interface: tui không biết bạn là ai, nhưng tôi biết bạn có thể làm được cái tôi cần

# Build & Run

## With minimal Dockerfile

### Dockerfile

```dockerfile
# Base image sử dụng alpine, nhẹ, phù hợp production
FROM alpine

# thư mục làm việc trong container.
WORKDIR /app/

# thêm binary app đã build sẵn từ local vào container.
ADD ./app /app/

# chạy binary đó khi container khởi động.
ENTRYPOINT ["./app"]

# Nếu binary chưa có quyền chạy, thêm:
RUN chmod +x /app/app

```

### Run

```shell
# 1. Build binary Go to file ./app
go build -o app .

# 2. Build Docker image
docker build -t social_todo_app .

# 3. Run Docker image
docker run --rm --env-file .env -p 3000:3000 social_todo_app
```

&rarr; Đây là kiểu Dockerfile đơn giản, dùng khi app đã được build sẵn bên ngoài (có thể CI/CD hoặc build tay).
&rar; Cách này cực kỳ nhanh, dễ hiểu. Nhưng phụ thuộc vào việc bạn phải build trước thủ công (go build -o app .).

### Notes
- Có thể sửa file pg_hba.conf của postgres nếu bị lỗi `no pg_hba.conf entry for host "172.17.0.2"`
- Trường hợp dùng postgres, trong file env localhost phải thay bằng `172.17.0.1`. Đây là IP của host trong container hoặc có thể chạy lệnh này để lấy IP host:
```shell
docker run --rm alpine sh -c "ip route | awk '/default/ { print \$3 }'"
```

## Dockerfile multi stage

**Ván đề với cách build đầu tiên:** Khi code thay đổi -> 

### Dockerfile

```dockerfile
# Dockerfile with multi stage
#FROM golang:1.20-alpine as builder
 #RUN mkdir /app
 
 FROM social-todo-service-cached as builder
 
 ADD . /app/
 WORKDIR /app
 RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o demoApp .
 
 
 FROM alpine
 WORKDIR /app/
 COPY --from=builder /app/demoApp .
 ENTRYPOINT ["/app/demoApp"]
```