Создание приватной блокчейн-сети:
1) Создаем несколько аккаунтов:
geth account new
2) создание genesis.json файла:
puppeth
3) инициализация genesis файла:
geth init genesis.json
    -если при создании возникает ошибка:
    удаляем существующую базу- geth removedb
    -если не помогает удаляем папку /home/taketa/.ethereum/geth/chaindata
4) запускаем ноду:
geth --networkid 123 --rpc console
5) копируем адрес ноды и подставляем в него свой ip (пример: enode://6ccd34b8f01c3ae3240eee1467dbf97217c85b866c2756e8eb1ddd67377f4c6209afff709e908b6863e44cf48b80ba215894da3931aa60f1d5a1bcb9fac239a8@192.168.88.62:30303)

Подключение пиров к ноде:
1) Инициализируем genesis.json:
geth init genesis.json
2) подключаемся к ноде:
geth --networkid 123 --bootnodes enode://6ccd34b8f01c3ae3240eee1467dbf97217c85b866c2756e8eb1ddd67377f4c6209afff709e908b6863e44cf48b80ba215894da3931aa60f1d5a1bcb9fac239a8@192.168.88.62:30303 console

Проверка подключились ли пиры:
admin.peers