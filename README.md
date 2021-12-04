# multiple-docker-compose-communication-example
Example of how containers launched with different docker-compose can communicate with each other.

## How to

### カスタムブリッジネットワークを作る方法

- 事前にブリッジネットワークを作成しておく必要がある
- ホスト名での名前解決が可能

Front

```yaml
version: '3'
services:
  front:
    image: ubuntu:22.04
    tty: true
    command: "sleep infinity"
    networks:
      - shared-net
networks:
  shared-net:
    driver: bridge
```

Server

```yaml
version: '3'
services:
  server:
    build: .
    ports:
      - 8080:8080
    networks:
      - front_shared-net
networks:
  front_shared-net:
    external: true
```

### デフォルトブリッジネットワークに参加させる方法

- 事前にネットワークを作成しておく必要はない
- ホスト名での名前解決ができず、IPアドレスでしか通信ができない

Front

```yaml
version: '3'
services:
  front:
    image: ubuntu:22.04
    tty: true
    command: "sleep infinity"
    network_mode: bridge
```

Server

```yaml
version: '3'
services:
  server:
    build: .
    ports:
      - 8080:8080
    network_mode: bridge
```
