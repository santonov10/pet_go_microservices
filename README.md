# Pet-проект "Менеджер задач" на микросервисах

запуск
- make docker_compose_up

микросервисы:
- api-gateway: предоставляет REST API для клиентского приложения
- user: авторизация на JWT, информация о пользователе. Так же при запуске применяет последние миграции
- tasks: микросервис работающий с задачами
- web-client: SPA приложение на React

в планах:
- rabbitmq
- swagger
- grpc web
- верстка
- auth0