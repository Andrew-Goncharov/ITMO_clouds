# Лабораторная работа № 2
### Команда
- Соболевская Надежда K34212
- Осипова Валерия K34202
- Гончаров Андрей K34211
- Донина Дарья К34202

## Задание

Поднять kubernetes кластер локально, в нём развернуть свой сервис, используя 2-3 ресурса kubernetes.

## Основная часть

1. Устанавливаем инструменты для работы с Kubernetes `kubectl`` с помощью команды:
```
curl -LO https://dl.k8s.io/release/v1.28.3/bin/windows/amd64/kubectl.exe
```

А также инструмент для запуска Kubernetes - Minikube по прямой ссылке из браузера.

2. Запустим minikube с помощью команды:

```
minikube start --driver=docker
```
<img src="./img/pic1.jpg"/>

3. Создадим файл web.yml:

```
apiVersion: apps/v1
kind: Deployment

metadata:
  name: webserver
  labels:
    app: webserver

spec:
  selector:
    matchLabels:
      app: webserver

  replicas: 2

  template:
    metadata:
      name: webserver
      labels:
        app: webserver
    spec:
      containers:
        - name: webserver
          image: web-server
```
В файле указан kind - Deployment. Он отвечает за развертывание подов, следит за их состоянием. Параметр replicas отвечает за количетсво экземпляров объекта. В нашем случае их будет 2.

Для создания файла выполним команду:

```
kubectl create -f ./web.yml
```

Для получения deployments и подов выполним команды:
```
kubectl get deployments 
kubectl get pods
```
<img src="#"/>

4. Создаем файл для сервиса ser.yml со следующим содержимым:

```
apiVersion: v1
kind: Service

metadata:
  name: service

spec:
  type: LoadBalancer
  ports:
    - targetPort: 443
      port: 443
      nodePort: 32115
  selector:
    app: webserver
```

В качестве сервиса был выбран LoadBalancer. Он обеспечивает баланс нагрузки для сервиса.

Для создания файла выполним команду:

```
kubectl create -f ./ser.yml
```
Для получения сервисов выполним команду:

```
kubectl get service
```
<img src="#"/>

5. Проверим работу сервиса.

Для этого выполним команду:

```
./minikube.exe service ser
```
<img src="./img/pic2.jpg"/>

Переходим по адресу и в браузере видим ответ от сервера.

<img src="./img/pic5.jpg"/>


6. Вывод

В результате выполнения данной лабораторной работы был поднят kubernetes кластер локально, в нём был развернут http сервер.
