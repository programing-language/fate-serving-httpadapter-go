# 代码运行测试
```
go run main.go
```
```
curl -v localhost:8080/feature -H 'Content-Type: application/json' -d '{}'
```
# 镜像测试
```
docker build -t cjx/adapter:1.0.0 .
docker run -d --name adapter -p 20000:8080 cjx/adapter:v1
docker logs -f adapter
```
```
curl -v localhost:20000/feature -H 'Content-Type: application/json' -d '{}'
```
# 请求分析测试
```
curl -v http://localhost:8080/feature?token=1234&page=1 \
-H 'Content-Type: application/json' \
-H 'app_key: 123456' \
-H 'access_token: qwertyuiop' \
-H 'expire_time: 1670395896' \
-d '{
    "head": {
        "serviceId": "chenjinxin.cn#1995"
    },
    "body": {
        "featureData": {
            "x0": -2.123,
            "x1": -1.456,
            "x2": 0.789,
            "x3": 1.101112,
            "x4": 2.131415
        },
        "sendToRemoteFeatureData": {
            "age": 18
        }
    }
}'
```