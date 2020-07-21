# AWS-Mahjong


## Description
https://c9katayama.hatenablog.com/entry/2014/12/15/002712


## Test


`count=1`などで見かけ上はcacheされていないように見えるが正しくテストできていない。
最も確実な方法が今のところcleanすること。

```sh
go clean -testcache && go test -v ./... -failfast
```

```sh
go clean -testcache && go test -v ./server -tags=integration
```


## Api Document

api endpoint

```sh
http://localhost:8000/
```

---

### [x]  GET `/rooms`

現在開かれている部屋の一覧を取得する。

**RESPONSE**

`status 200`


```json
[
  {
    "room_name": "provident",
    "room_capacity": 3,
    "room_member_count": 1
  },
  {
    "room_name": "fugiat",
    "room_capacity": 4,
    "room_member_count": 1
  }
]
```

`default`

NO CONTENT

---

### [ ] POST `/rooms`

部屋を作成する。作成の際にプレイ人数を指定できる。指定した人数集まった場合には自動的にスタートする。
```json
{
  "user_name": "Elwin Ebert",
  "room_name": "possimus",
  "room_capacity": 3
}
```
 
**RESPONSE**

`status 200`


```json
{
  "room_name": "fugiat",
  "room_capacity": 4
}
```

`default`

NO CONTENT

---

### [ ] PUT `/rooms/{room_name}/join`

部屋に参加する。作成の際にプレイ人数を指定できる。指定した人数集まった場合には自動的にスタートする。
`user_id`に関してはクライエント側で適当に生成する。他のプレイヤーとかぶらなければ正直なんの値でも良い。
```json
{
  "user_id": "ba214aff-365b-398f-8492-80e3057f0d44"
}
```
 
**RESPONSE**

`status 200`


```json
{
  "room_capacity": 4,
  "room_member_count": 1
}
```

`default`

NO CONTENT

---

### [ ] PUT `/rooms/{room_name}/leave`

現在joinしている部屋から退出する。ちなみにゲームが開始されてからの退出は受け付けていない。

```json
{
  "user_id": "0b13481b-3006-31bd-ab96-4790616e0af1"
}

```


**RESPONSE**

`status 200`


```json
{
  "room_capacity": 4,
  "room_member_count": 1
}
```

`default`

NO CONTENT

---


### [ ] POST `/rooms/{room_name}/dahai`

haiの箇所に入る値はサーバーから送信される牌の名前をそのまま使用で問題ない。


```json
{
  "user_id": "0b13481b-3006-31bd-ab96-4790616e0af1",
  "hai": "chun"
}
```

**RESPONSE**

`status 200`


```json
{
    "name": "",
    "zikaze": "east",
    "tsumo": "manzu3",
    "tehai": ["hanzu3", "souzu3"],
    "kawa": [
      {"isSide": false, "name": "manzu3"},
      {"isSide": false, "name": "manzu3"}
    ],
    "naki_actions": {
      "pon": [["manzu1", "manzu1"], ["manzu2", "manzu2"]],
      "kan": [],
      "chii": [["manzu1", "manzu2"]]
    },
    "naki": [
      [
        {"isOpen": false, "isSide": false, "name": "hatu"},
        {"isOpen": false, "isSide": false, "name": "hatu"},
        {"isOpen": false, "isSide": true, "name": "hatu"}
      ],
      [
        {"isOpen": true, "isSide": false, "name": "manzu1"},
        {"isOpen": false, "isSide": false, "name": "manzu1"},
        {"isOpen": false, "isSide": false, "name": "manzu1"},
        {"isOpen": true, "isSide": false, "name": "manzu1"}
      ]
    ]
}
```

`default`

NO CONTENT

---


### [ ] POST `/rooms/{room_name}/naki`

```json
{
  "user_id": "0b13481b-3006-31bd-ab96-4790616e0af1",
  "naki": "cancel"
}
```

```json
{
  "user_id": "0b13481b-3006-31bd-ab96-4790616e0af1",
  "naki": "pon",
  "mentsu": []
}
```

**RESPONSE**

`status 200`


```json
{
    "name": "",
    "zikaze": "east",
    "tsumo": "manzu3",
    "tehai": ["hanzu3", "souzu3"],
    "kawa": [
      {"isSide": false, "name": "manzu3"},
      {"isSide": false, "name": "manzu3"}
    ],
    "naki_actions": {
      "pon": [["manzu1", "manzu1"], ["manzu2", "manzu2"]],
      "kan": [],
      "chii": [["manzu1", "manzu2"]]
    },
    "naki": [
      [
        {"isOpen": false, "isSide": false, "name": "hatu"},
        {"isOpen": false, "isSide": false, "name": "hatu"},
        {"isOpen": false, "isSide": true, "name": "hatu"}
      ],
      [
        {"isOpen": true, "isSide": false, "name": "manzu1"},
        {"isOpen": false, "isSide": false, "name": "manzu1"},
        {"isOpen": false, "isSide": false, "name": "manzu1"},
        {"isOpen": true, "isSide": false, "name": "manzu1"}
      ]
    ]
}
```

`default`

NO CONTENT

---

## WebSocket Document

websocket endpoint

```sh
http://localhost:8000/ws
```

## Server Events


### [ ] room_status

```json
{
  "room_name": "modi",
  "room_member_count": 1,
  "room_capacity": 3
}
```

### [ ] game_status

他のプレイヤーの打牌、鳴きなど状況に変更があった際には更新がこのイベントで通知される。

```json
{
  "bakaze": "east",
  "deck_count": 36, 
  "turn_player_index": 0,
  "oya_player_index": 0,
  "players": [
    {
      "id": "b3448161-f9dc-343c-a3d5-fae551d1158a",
      "name": "",
      "zikaze": "east",
      "tsumo": "manzu3",
      "hand": ["hanzu3", "souzu3"],
      "kawa": [
        {"isSide": false, "name": "manzu3"},
        {"isSide": false, "name": "manzu3"}
      ],
      "naki_actions": {
        "pon": [["manzu1", "manzu1"], ["manzu2", "manzu2"]],
        "kan": [],
        "chii": [["manzu1", "manzu2"]]
      },
      "naki": [
        [
          {"isOpen": false, "isSide": false, "name": "hatu"},
          {"isOpen": false, "isSide": false, "name": "hatu"},
          {"isOpen": false, "isSide": true, "name": "hatu"}
        ],
        [
          {"isOpen": true, "isSide": false, "name": "manzu1"},
          {"isOpen": false, "isSide": false, "name": "manzu1"},
          {"isOpen": false, "isSide": false, "name": "manzu1"},
          {"isOpen": true, "isSide": false, "name": "manzu1"}
        ]
      ]
    }
  ]
}
```
