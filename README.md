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

### GET `/rooms`

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
 

## WebSocket Document

websocket endpoint

```sh
http://localhost:8000/ws
```

## Client Events


### create_room

部屋を作成して参加する。作成の際にプレイ人数を指定できる。指定した人数集まった場合には自動的にスタートする。
```json
{
  "user_name": "Elwin Ebert",
  "room_name": "possimus",
  "room_capacity": 3
}
```

### join_room

部屋の名前を指定してユーザーを参加させる。しっかりと`/rooms`の結果で取得したroomにjoinすること。

```json
{
  "user_name": "Ms. Lilliana Walker",
  "room_name": "porro"
}
```

### leave_room

ほとんど使わないと思うが一応。ちなみに退出した場合にはそのゲームは即時終了する。
socket.ioは接続が切れた際に自動的にdisconnectしてくれるが、その際にも参加していたゲームは強制終了する。

```json
{
  "room_name": "laudantium"
}

```



### dahai

牌を手牌から捨てる際に使う。このイベントを送信するとnew_statusが走る。
基本的にはnew_statusのturn_player_indexに対応するplayerがtsumoを持っている。
これを捨てるために使用する。

```json
{
  "room_name": "laudantium",
  "dahai": "chun"
}

```


### naki

打牌に対して鳴きを行うか否かを選択する。
鳴きを行わない場合に関してもcancelというイベントを送信する。
actionsの種類に関しては`/naki/naki.go`のNakiActionを参照。


```json
{
  "room_name": "laudantium",
  "action": "cancel"
}
```

```json
{
  "room_name": "laudantium",
  "action": "pon",
  "tiles": ["manzu1", "manzu1"]
}
```


## Server Events


### create_room_error

```
```

### join_room_error

```
```

### new_room_status

```json
{
  "room_name": "modi",
  "room_member_count": 1,
  "room_capacity": 3
}
```

### game_start

ゲームの開始を通知するだけ。現状は空文字を返すのみ。

```json
```

### game_end

ゲームの終了を通知する。正常に終了した場合と、誰かが退出して強制的に終了した場合も含む。

```json
```

### new_status

他のプレイヤーの打牌、鳴きなど状況に変更があった際には更新がこのイベントで通知される。

```json
{
  "bakaze": "east",
  "deck_card_count": 36, 
  "oya_player_index": 1,
  "turn_player_index": 0,
  "jicha_player_index": 2,
  "players": [
    {
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
