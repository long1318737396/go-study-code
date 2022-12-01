package db

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"strconv"
)

//输入协议名称、IP地址、端口号、密码、以及使用哪个index编号的数据集
func RedisConnect(protocol, ip, pwd string, port, userInfoIndex int) {
	//为Redis的拨号连接传入密码pwd
	options := redis.DialPassword(pwd)
	//对Redis数据库进行拨号，获取Redis连接对象
	conn, err := redis.Dial(protocol, ip+":"+strconv.Itoa(port), options)
	//在使用完成后，关闭Redis连接
	defer func() {
		if err = conn.Close(); err != nil {
			log.Println("close error:", err)
		}
	}()
	if err != nil {
		log.Println("dial error:", err)
		return
	}

	//每一个具体的Redis操作，本质上都是调用Do方法并传入相应的参数
	//切换到需要操作的Redis数据集，获取响应结果
	reply, err := conn.Do("select", userInfoIndex)
	if err != nil {
		log.Println("select error:", err)
		return
	}
	fmt.Println("select reply=", reply)
	//为字符串变量bar设置值为“3”
	//redis.String()方法用于将返回结果直接转换为Go语言的string类型，类似的还有redis.Float64(),redis.Int()等
	reply, err = redis.String(conn.Do("set", "bar", "3"))
	if err != nil {
		log.Println("set error:", err)
		return
	}
	fmt.Println("set reply=", reply)
	//获取为字符串变量bar的值为“3”
	reply, err = redis.String(conn.Do("get", "bar"))
	if err != nil {
		log.Println("get error:", err)
		return
	}
	fmt.Println("get reply=", reply)

	//为列表变量list1从右端添加“0”这个元素
	reply, err = redis.Int64(conn.Do("rpush", "list1", "777"))
	if err != nil {
		log.Println("rpush error:", err)
		return
	}
	fmt.Println("rpush reply=", reply)

	//获取list1列表第1个元素
	reply, err = redis.String(conn.Do("lindex", "list1", "0"))
	if err != nil {
		log.Println("lrange error:", err)
		return
	}
	fmt.Println("lrange reply=", reply)

	//为集合set1添加 "tom"
	reply, err = redis.Int64(conn.Do("sadd", "set1", "tom"))
	if err != nil {
		log.Println("sadd error:", err)
		return
	}
	fmt.Println("sadd reply=", reply)

	//获取集合set1中的所有元素对应的字节切片数组
	replyBytes, err := redis.ByteSlices(conn.Do("smembers", "set1"))
	if err != nil {
		log.Println("smembers error:", err)
		return
	}
	//将set1中的所有元素对应的字节切片数组中的第一个元素转换为字符串
	fmt.Println("smembers reply=", string(replyBytes[0]))

	//为有序集合zset1中添加分值为500的元素ele1
	reply, err = redis.Int64(conn.Do("zadd", "zset1", "500", "ele1"))
	if err != nil {
		log.Println("zadd error:", err)
		return
	}
	fmt.Println("zadd reply=", reply)

	//获取zset1中的所有元素和分值
	replyBytes, err = redis.ByteSlices(conn.Do("zrange", "zset1", "0","-1", "withscores"))
	if err != nil {
		log.Println("zrange error:", err)
		return
	}
	//将zset1中的所有元素对应的字节切片数组中的第一个元素转换为字符串
	fmt.Println("zrange 操作的元素名称为：", string(replyBytes[0]),"\nzrange 操作的元素的分值为：",string(replyBytes[1]))

	//为散列表hash1中添加字段name值为Tom
	reply, err = redis.Int64(conn.Do("hset", "hash1", "name","Tom"))
	if err != nil {
		log.Println("hset error:", err)
		return
	}
	fmt.Println("hset reply:",reply)

	//获取hash1中的第一个字段和值
	replyBytes, err = redis.ByteSlices(conn.Do("hgetall", "hash1"))
	if err != nil {
		log.Println("hset error:", err)
		return
	}
	fmt.Println("hset 获取得字段名称为:",string(replyBytes[0]),"\nhset 获取得字段值为:",string(replyBytes[1]))
}
