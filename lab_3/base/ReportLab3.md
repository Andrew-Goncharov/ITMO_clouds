# Лабораторная работа № 3
### Команда
- Соболевская Надежда K34212
- Осипова Валерия K34202
- Гончаров Андрей K34211
- Донина Дарья К34202

## Задание

Сделать, чтобы после пуша в репозиторий автоматически собирался докер образ и результат его сборки сохранялся на сервере.

## Основная часть
Перед началом выполнения работы воспользуемся возможностями GitHub Action Secrets, в котором сохраним переменные ADDRESS, PASS, PATH, USERNAME, SSH_KEY.

![image](https://github.com/Andrew-Goncharov/ITMO_clouds/assets/64967406/b5dc6e2f-47b9-4330-aedd-47dfee808d01)

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

      - run: docker build ./lab_3/base > builder.log 2>&1

      - uses: appleboy/scp-action@v0.1.4
        with:
          host: ${{ secrets.ADDRESS }}
          username: ${{ secrets.USERNAME }}
          password: ${{ secrets.PASS }}
          key: ${{ secrets.SSH_KEY }}
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

`- run: docker build ./lab_3/base > builder.log 2>&1` - сборка с логированием в файл builder.log. `2>&1` это уровень логирования.

`- uses: appleboy/scp-action@v0.1.4` - далее мы используем проект scp-action который позволяет отправить файлы по ssh.

В Uses идет вызов сторонних проектов. В with аргументы: 
```
- uses: appleboy/scp-action@v0.1.4
        with:
          host: ${{ secrets.ADDRESS }}
          username: ${{ secrets.USERNAME }}
          key: ${{ secrets.SSH_KEY }}
          password: ${{ secrets.PASS }}
          port: 22
          source: "builder.log"
          target: ${{ secrets.PATH }}
```

В данном случае: адрес и порт хоста, логин и пароль пользователя, файлы для передачи и путь сохранения.

#### 2. Тестирование
Сделаем пуш в репозиторий. Видим успешный билд: 

<img width="956" alt="image_2023-12-17_19-54-13" src="https://github.com/Andrew-Goncharov/ITMO_clouds/assets/64967406/1d7e2011-4c74-47cb-bdbf-d6a58477e944">

Подключимся к серверу и проверим наличие файла build.log:

<img width="273" alt="image" src="https://github.com/Andrew-Goncharov/ITMO_clouds/assets/64967406/22a62829-01e5-47e8-869f-0186c66b9126">

## Вывод
В результате выполнения данной лабораторной работы была реализована автоматическая сборка докер образа после пуша в репозиторий и результат его сборки сохраняется на сервер. Удалось настроить CI/CD с помощью GitHub Actions и протестировать работу.
