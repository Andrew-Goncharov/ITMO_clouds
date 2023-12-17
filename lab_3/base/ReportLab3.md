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

<img width="955" alt="image" src="https://github.com/Andrew-Goncharov/ITMO_clouds/assets/64967406/5b3ede5b-bd61-4592-90ae-3e840e5bed1a">

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
В начале задаем имя сценария: build. Далее указывается, что сценарий должен запускаться при пуше:

```
on:
  push:
```

Далее указывается, какие действия будут выполняться. 
В строчке `runs-on: ubuntu-latest` указываем на какой операционной системе будут вызываться следующие скрипты.

Далее по шагам:
`- uses: actions/checkout@master` - загружает виртуалку в репозиторий.

`- run: docker build . > builder.log 2>&1` - сборка с логированием в файл builder.log. `2>&1` это уровень логирования.

`- uses: appleboy/scp-action@v0.1.4` - далее мы используем проект scp-action который позволяет отправить файлы по ssh.

В Uses идет вызов сторонних проектов. В with аргументы: 
```
- uses: appleboy/scp-action@v0.1.4
        with:
          host: ${{ secrets.ADDRESS }}
          username: ${{ secrets.USERNAME }}
          password: ${{ secrets.PASS }}
          port: 22
          source: "builder.log"
          target: ${{ secrets.PATH }}
```

В данном случае: адрес и порт хоста, логин и пароль пользователя, файлы для передачи и путь сохранения.

