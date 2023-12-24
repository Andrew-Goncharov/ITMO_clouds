# Лабораторная работа № 4
### Команда
- Соболевская Надежда K34212
- Осипова Валерия K34202
- Гончаров Андрей K34211
- Донина Дарья К34202

## Задание

Сделать мониторинг сервиса, поднятого в kubernetes

## Основная часть

1. Запускаем сервис с предустановленным экспортером метрик prometheus


```
docker run -d --name prometheus-exporter -p 9100:9100 prom/node-exporter
```

<img src="./img/image_12.png"/>

2. Скачиваем prometheus для Windows, настраиваем prometheus.yml и далее запускаем prometheus.exe 

Меняю название и порт нашего сервиса

<img src="./img/image_7.png"/>

Запускаю exe

<img src="./img/image_13.png"/>

3. Открываем prometheus в браузере по адресу 127.0.0.1:9090

<img src="./img/image_1.png"/>

4. Проверяем доступность цели

<img src="./img/image_11.png"/>

5. После установки grafana, открываем grafana  по адресу localhost:3000

- уснановка графана
- открываем grafana  по адресу localhost:3000 
- меняем пароль от пользователя admin

далее видим следующее окно

<img src="./img/image_4.png"/>

6. Подключаем prometheus к grafana

<img src="./img/image_2.png"/>

7. Создаём dashboard

Можно создать с нуля

<img src="./img/image_9.png"/>

Но мы импортируем готовый шаблон в grafana

<img src="./img/image_6.png"/>

8. Открываем dashboard

<img src="./img/image_8.png"/>


## Вывод

В результате выполнения данной лабораторной работы был поднят сервис отправляющий метрики в kubernates, настроена связь grafana и kubernates, оформлено отображение метрик в dashboard grafana.