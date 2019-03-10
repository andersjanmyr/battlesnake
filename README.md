# battlesnake

A [Battlesnake AI](http://battlesnake.io) written in Go.

## Getting started

Download the [battlesnake engine](https://github.com/battlesnakeio/engine/releases). Start it locally with `engine dev`. Open UI in browser at [http://localhost:3010/](http://localhost:3010).


## Install Dependencies

```
make install
```

## Run

```
Make run
```

4) Open in browser: [http://127.0.0.1:9000/](http://127.0.0.1:9000/)

### Snakes

The API supports multiple snakes and they can be access through the URL.
`http://127.0.0.1:9000/{kind}/{id}`

* `kind` - `empty` (no implementation), `horry` (avoids immediate obstacles and prefers horizontal movement), or `randy` (avoid immediate obstacles, otherwise random movement).
* `id` - Any value, used to separate snakes of the same `kind`.

#### Example

URLs for four different snakes. One `empty`, one, `horry` and two `randy`.

```
http://127.0.0.1:9000/emtpy/1
http://127.0.0.1:9000/horry/1
http://127.0.0.1:9000/randy/1
http://127.0.0.1:9000/randy/2
```
