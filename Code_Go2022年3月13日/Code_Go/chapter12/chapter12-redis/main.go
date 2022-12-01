package main

import "GoBook/code/chapter12-redis/db"

func main() {
	db.RedisConnect("tcp", "127.0.0.1", "gopher2020", 6379, 0)
}
