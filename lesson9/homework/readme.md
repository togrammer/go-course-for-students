## Домашняя работа №9

Необходимо доработать сервис объявлений из ДЗ №9.
Для этого аккуратно перенесите его реализацию в `lesson9/homework`.

## Требования

- добавить методы на получение/удаление пользователя, удаления объявления (только для автора)
- добавить к вашему сервису gRPC интерфейс, таким образом чтобы одновременно была возможность взаимодействовать
  с помощью REST и gRPC.
- на каждый gRPC вызов написать хотя бы один тест
- доавить logger и panic interceptor с собственным логгером
- сделать graceful shutdown

## Критерии оценки

- новые методы - до 2-х баллов
- gRPC интерфейс - до 2-х баллов
- тесты - до 1-х баллов
- интерсепторы - до 2-х баллов
- graceful shutdown - до 3-х баллов

## Примечание

В ports/grpc описана примерная схема proto контракта, если вам необходимо изменить ее для вашего сервиса - меняйте,
подстраиваться под этот контакт не нужно.

Скорее всего при реализации graceful shutdown для http сервиса вам необходимо будет использовать стандартный http.Server,
пример изменения сервера так же можете посмотреть в ports/httpgin/server.go. Эти измеения так же затронут тесты, а именно создание
тестового сервера, вместо метода server.Handler() используйте просто поле server.Handler, вместо метода server.Listen() - server.ListenAndServe()

Пример теста для grpc можно посмотреть в tests/grpc_test.go

Сгенерированные с помощью protoc файлы и реализацию сервиса положите рядом с service.proto.