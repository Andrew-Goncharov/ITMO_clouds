# Лабораторная работа № 3
### Команда
- Соболевская Надежда K34212
- Осипова Валерия K34202
- Гончаров Андрей K34211
- Донина Дарья К34202

## Задание

Сделать, чтобы после пуша в репозиторий автоматически собирался докер образ и результат его сборки сохранялся на сервер.

## Основная часть
Перед началом выполнения работы воспользуемся возможностями GitHub Action Secrets, в котором сохраним переменные ADDRESS, PASS, PATH.

![image](https://github.com/Andrew-Goncharov/ITMO_clouds/assets/64967406/a454c71c-1872-4fee-bd3c-5166b211b2b4)

#### 1.  Настройка CI/CD с помощью GitHub Actions.
Необходимо создать `.yml` файл в директории `/.github/workflows`. 
Содержимое файла `lab3.yml`

```
name: build

on:
  push:

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master

      - run: docker build . > builder.log 2>&1

      - uses: appleboy/scp-action@v0.1.4
        with:
          host: ${{ secrets.ADDRESS }}
          username: ${{ secrets.USERNAME }}
          password: ${{ secrets.PASS }}
          port: 22
          source: "builder.log"
          target: ${{ secrets.PATH }}
```
