# holiday
查询指定日期是否为节假日
## 使用方法
```
go run main.go
holiday.json需要放在同一目录下
```
访问 http://localhost:9999/holiday?d=20200304 这样即可
返回值为0说明正常工作日
返回值为1说明是法定节假日
## 维护
日期在holiday.json中进行维护，每年根据国家法定假日自己修改
