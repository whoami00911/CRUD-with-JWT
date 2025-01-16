## CRUD Приложение предоставляющее web API к данным
### Стэк
- **Go**: Версия 1.22.5

### Начало работы

Для запуска программы используйте следующую команду:

```bash
go run main.go
```

### Использование
**Для легкого взаимодействия с API вы можете перейти на:**
```http://localhost:8080/swagger/index.html```

**Взаимодействие через терминал:**
- **Добавить сущность:**
Добавить сущность с уже существующим IP нельзя ни в кеш ни в БД
```
curl -v --request POST "http://localhost:8080/Abuseip/" -H "Content-Type: application/json" -d "{\"ipAddress\": \"192.25.6.39\", \"isPublic\": true, \"ipVersion\": 4, \"isWhitelisted\": false, \"abuseConfidenceScore\": 100, \"countryCode\": \"CN\", \"countryName\": \"China\", \"usageType\": \"WebHosting\", \"isp\": \"ChinChinarem\"}"
```

- **Получить все сущности из БД:**
```
curl -v --request GET http://localhost:8080/Abuseip/
```

- **Получить одну сущность из кеша(Хранится 15 секунд после добавления) или БД:**
- Комманда: ```curl -v --request GET http://localhost:8080/Abuseip/192.25.6.39```
- Вывод:
Из БД: "isFromDB":true | "IsCache":false
```
{"ipAddress":"192.25.6.39","isPublic":true,"ipVersion":4,"isWhitelisted":false,"abuseConfidenceScore":100, <...>  "IsCache":false,"isFromDB":true}*
```
Из кеша: "IsCache":true | "isFromDB":false
```
{"ipAddress":"192.25.6.39","isPublic":true,"ipVersion":4,"isWhitelisted":false,"abuseConfidenceScore":100, <...> "IsCache":true,"isFromDB":false}*
```

- **Удалить все сущности в БД и кеше:**
```
curl -v --request DELETE http://localhost:8080/Abuseip/
```
- **Удалить сущность по IP:**
```
curl -v --request GET http://localhost:8080/Abuseip/192.25.6.39
```

- **Обновить сущность по IP**
```
curl -v --request PUT "http://localhost:8080/Abuseip/" -H "Content-Type: application/json" -d "{\"ipAddress\": \"192.25.6.39\", \"isPublic\": true, \"ipVersion\": 4, \"isWhitelisted\": false, \"abuseConfidenceScore\": 100, \"countryCode\": \"CN\", \"countryName\": \"Chinazes\", \"usageType\": \"WebHosting\", \"isp\": \"ChinChinarem\"}"
```
