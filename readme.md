#Установка

```
git clone https://github.com/eonias189/calculationService.git <path>

cd <path>
scripts/install
```
path - путь для установки

>Для установки необходим компилятор go, nodejs и yarn (yarn можно установить через `npm i -g yarn`)

#Запуск

```
scripts/run-orchestrator
scripts/run-frontend
```

####Для запуска агентов:
```
go run agent/cmd/main.go <port> <threads>
```

где port - четырёхзначное число (не 8081), threads - количество потоков агента

>Каждый скрипт нужно запускать в отдельном терминале. Агентов можно запустить несколько, но у них должны быть разные порты!

Сервер с оркестратором работает на http://loaclhost:8081
Сервер с фронтендом работает на http://localhost:3000
Сервера агентов работают на http://localhost:\<port\>, где port - порт, с которым запущен агент


##Взаимодействие с серверами
Эндпоинты для взаимодействия с серверами указаны в [endpoints.txt](https://github.com/eonias189/calculationService/blob/main/endpoints.txt).</br>
Протокол взаимодействия указан в [.proto](https://github.com/eonias189/calculationService/blob/main/.proto)</br>
Сервер с фронтендом имеет графический интерфейс для взаимодействия

>Если сервер агента был отключен во время выполнения задач, то перед тем, как подключать его обратно, необходимо дождаться, когда задачи, который выполнял этот агент, перейдут в режим ожидания (их цвет изменится на серый), иначе они не выполнятся никогда

##Как Работает
Пользователь через сервер с фронтендом посылает на сервер оркестратора запросы на добавление задачи/изменение таймаутов/получение информации о задачах/агентах/таймаутах.</br>
Информацию о задачах/агентах/таймаутах оркестратор берёт из базы данных. В отдельном потоке оркестратор с определённым интервалом делает запросы к агентам, чтобы обновлять информацию о них. Также в отдельном потоке он проверяет ping агентов и при обнаружении отключенного агента изменяет статус всех задач, которые выполняет этот агент на "ожидание распределения".</br>

Агенты в отдельном потоке с определённым интервалом (при наличии свободного потока для выполнения) делают запрос к оркестратору на получение новой задачи и после выполнения делают запрос на установление результата.
</br>
#Возможные вопросы

###Где документация к коду?
>Код написан по принципам самодокументирования

###Зачем нужен модуль backend?
>В этом модуле хранится код, который используется и в модуле оркестратора, и в модуле агента