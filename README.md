### Запуск проекта

Для начала необходимо создать файл `.env` и вставить туда содержимое `example.env`.

Если вы хотите запустить платформу в develop режиме, то выполните команду `go mod vendor`. А после `docker compose -f docker-compose.yaml up --build`.  
Если вы хотите использовать релизный режим, то выполните команду `docker compose -f release-compose.yaml up --build`.

После запуска композа перейдите в Minio и создайте bucket с названием, который вы указали в `.env` (MINIO\_BUCKET). А после дайте этому bucket права на чтение и запись.  

Всё готово к работе.

### Кодогенераторы

* swag init - генерация swagger
* sqlc generate - генерация sql запросов по `./queries/*`  
* goverter gen path\_to\_package - генерация преобразователя структур
