# soy-log-collector
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/7f4951281ac54cd2b75a5c23d1fb9cb5)](https://www.codacy.com/gh/soyoslab/soy_log_collector/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=soyoslab/soy_log_collector&amp;utm_campaign=Badge_Grade)
[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/7f4951281ac54cd2b75a5c23d1fb9cb5)](https://www.codacy.com/gh/soyoslab/soy_log_collector/dashboard?utm_source=github.com&utm_medium=referral&utm_content=soyoslab/soy_log_collector&utm_campaign=Badge_Coverage)

## Introduction

This project sends the messages got from soy_log_generator to soy_log_explorer.
The internal process is below.
1. Collect messages received from soy_log_generator according to hot/cold and pushes them to the corresponding queue.
2. Background daemon pops the message from the queue and unzips it.
3. Background daemon push the messages into redis-server with caching.
4. If hot messages, send the unzipped message to soy_log_explorer.
5. If cold messages, send zipped messages to soy_log_explorer.

## Installation

```bash
$ git clone https://github.com/soyoslab/soy_log_collector.git
$ cd soy_log_collector
```

## Usage

Set enviroment variables:
```bash
$ export RPCSERVER=0.0.0.0:YYYY			# Server Address
$ export EXPLORERSERVER=X.X.X.X:YYYY    # soy_log_explorer's RPC server Address
$ export DBADDR=X.X.X.X:YYYY			# Redis-server's Address
```

Example:
```
RPCSERVER=0.0.0.0:8972
EXPLORERSERVER=localhost:8973
DBADDR=localhost:6379
```

Run soy_log_collector:
```bash
$ go run cmd/server/server.go
```
