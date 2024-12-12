Проект

Этот проект представляет собой RESTful API, написанный на языке Go, для управления кошельками и их операциями. Проект использует PostgreSQL в качестве базы данных и Docker для контейнеризации.

Структура директории

    main.go: Основной файл проекта, содержащий функцию main.

    DataContext: Пакет, содержащий функции для работы с базой данных.
        DataContext.go: Файл, содержащий функции для работы с базой данных.

    RestApi: Пакет, содержащий функции для работы с RESTful API.
        AddWallet.go: Файл, содержащий функцию для добавления кошелька.
        GetWallet.go: Файл, содержащий функцию для получения информации о кошельке.
        UpdateWallet.go: Файл, содержащий функцию для обновления кошелька.
        errorJson.go: Файл, содержащий функции для работы с ошибками в формате JSON.

    Models: Пакет, содержащий структуры данных для кошельков и операций.
        Wallet.go: Файл, содержащий структуру данных для кошелька.

    config.env: Файл, содержащий переменные окружения для базы данных.

    docker-compose.yml: Файл, содержащий конфигурацию для Docker.

    main_api_test.go, Файл, содержащий функцию TestWalletAPI для тестирования REST API
    main_getWallet_test.go, Файл, содержащий функцию TestGetRPSLoad для тестирования RPS
    main_postWallet_test.go, Файл, содержащий функцию TestPostRPSLoad для тестирования RPS

Функциональность
    Добавление кошелька
    Получение информации о кошельке
    Обновление кошелька
    Запуск проекта

Установите Docker и Docker Compose на вашем компьютере.
    1. Откройте/Скачайте Docker Desktop.
    2. Выполните в директоии команду "docker-compose --env-file=config.env up --build".
    3. Проект будет доступен по адресу http://localhost:8080.

Тестирование
    Тесты для проекта:
    1. main_api_test.go (полный цикл тестирования REST API:
        1. Создание кошелька,
        2. Пополнение баланса,
        3. Получение баланса.
    2. main_getWallet_test.go (запускает тест для имитации нагрузки 100 GET запросов в секунду в течение 1 секунды)
    3. main_postWallet_test.go (запускает тест для имитации нагрузки 100 POST запросов в секунду в течение 1 секунды)
    Чтобы запустить тесты, выполните команду "go test -v" в директории проекта.
    PS:Количество запросов и время в тестах можно настраивать
Автор
    Солохин Иван Алексеевич
