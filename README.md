# Go Gin Base

<div>
<img src="https://github.com/truongbo17/go-base/blob/main/readme-logo.png?raw=true" align="right">

A development boilerplate based on the Gin framework, quickly build and develop web applications.
</div>

----

## Document

- Golang: https://go.dev
- Gin framework: https://gin-gonic.com

----

## Introduction

So why I use choose Gin framework?

- Based on benchmarks: https://gin-gonic.com/docs/benchmarks/
- Gin doesn't build everything from scratch, but relies heavily on Go's standard net/http (also one reason why i don't
  use fiber, even though fiber has better stats).
- Big community, really the biggest. With twice the stars and contributors of the second place competitor.
- Suitable for building RESTFul API, microservices or realtime applications.
- There is a full tutorial to build up a web server, even in Go Dev official blog!

----

## Installation

### Setup

```shell
  git clone https://github.com/truongbo17/go-base.git
  go mod download
```

```shell
  cp .env.example .env
```

### Run the Application

Run with air(hot reload):

```shell
  air server
```

Or simple:

```shell
  go run main.go server
```

Docs

```text
{{host}}/docs/swagger/index.html
```

----

## Feature

* I18N
* Config ✅
* Command / Console ✅
* Schedule
* Queue
* Swagger ✅
* View
* Logger ✅
* Database (DocumentDB(MongoDB), Relation DB(MYSQL)) ✅
* Mail
* File upload
* Upload file (local, s3)
* Kafka (producer, consumer)
* Authentication (JWT access token, refresh token, Google auth)
* Cache (local, redis) ✅
* Redis ✅
* GraphQL
* Http call service
* Middleware ✅
* Filter
* Push notify to telegram
* Router
* Worker

----

## Overview

----

## Struct