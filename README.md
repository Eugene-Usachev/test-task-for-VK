# Test task for VK.

<details>
    <summary>Task description (in Russian)</summary>

    Для постоянного мониторинга своих контейнеров необходимо приложение, которое будет постоянно отслеживать их состояние.

    Необходимо написать приложение на языках программирования Go и JavaScript(TS), которое получает ip адреса контейнеров docker, пингует их с определенным интервалом и помещает данные в базу данных.
    
    Получение данных о состоянии контейнеров доступно на динамически формируемой веб-странице.

    В результате выполнения задания должны появиться 4 сервиса:

        - Backend-сервис обеспечивает RESTful API для запроса данных из DB и добавления туда новых данных.

        - Frontend-сервис должен быть написан на JS с использованием любой библиотеки пользовательских интерфейсов (предпочтительно React). Берет данные через API Backend и отображает данные по всем IP адресам в виде таблицы: IP адрес, время пинга, дата последней успешной попытки. Для отображения данных в html можно использовать bootstrap или antd или подобное.

        - База данных PostgreSQL.

        - Сервис Pinger. Получает список всех docker-контейнеров, пингует их и отправляет данные в базу через API frontend.

    * Дополнительная сложность: добавление nginx, сервис очередей, использовать netns, отдельный конфиг для сервиса с верификацией.
    
    Результат:
    
        - В результате выполнения задания должны быть созданы Dockerfile для каждого сервиса и общий файл compose, которые собирает эти образы из исходников и запускает их, после чего можно зайти через http на определенный порт и увидеть данные о статусе машин, когда произойдет первый цикл опроса контейнеров.

        - Всё это размещено на github/gitlab в отдельном репозитории c README.md, в котором описан кратко функционал и шаги запуска.
    
![images/api_schema.png](images/api_schema.png)
</details>

# Running the application

1. Call `make` to generate `.go` files from `.proto` files;
2. Run `go mod tidy` in `backend` and `pinger` directories to generate `go.sum` files and install dependencies (for dev);
3. Run `docker-compose up`.

It is fine if `pinger` or `nginx` containers are restarted and 
if `pinger` can't ping `migrate` container.

Next you can see results at http://localhost (if you don't have conflicts). Every 30 seconds the page is updated,
and you can see the table of ping's results.

# On conflicts

Change Nginx port in `docker-compose.yaml` and `nginx.conf`. 

# About done work

All main tasks are done:

1. Backend service is ready;
2. Frontend service is ready with using React + TypeScript + Next.js;
3. Database is ready with its migrations;
4. Pinger service is ready.

# Additional complexity

Of the additional complexity, the task of adding Nginx
(for Reverse-Proxy, Caching and Load-Balance) has been implemented.

# Configuration

You can configure it in `docker-compose.yaml` file (for example, add more containers to ping list).

# For development

It uses `golang-migrate`. When you run `docker-compose up`, it will create the database and run migrations.
After it, you can call
`docker exec -i <container name> pg_dump -U admin -d container_monitoring --schema-only > ./db/schema.sql`
to get the schema.

All DTOs are described in `pkg/model` in `.proto` files. To generate `.go` files run `make`.