package main

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"time"
)

/**
TODO 第二题：排行榜问题
    说明
	1.	整体设计，使用Redis 有序集合实现    ZADD key score member
	2.	设置：  score+(标准时间-入库时间)  位移操作， ZADD
		月底的时间戳为标准，减去当前戳，积分位移，拼接，这样相同分数的人，越晚入库，越排的靠后
	3.  获取： ZRANK MEMBER 、ZRANGE MEMBER START END
		通过UID 获取到自己的索引（排名），然后分别向前、向后  +10  获取到指定范围的排名信息
*/

type Client struct {
	conn redis.Conn
}

var (
	currentMonthEndTime = 1659283199
	scoreLimitMin       = 0
	scoreLimitMax       = 10000
	sortLimit           = 10
	RedisClient         = new(Client)
)

func main() {
	uid := "a-0-0-1"
	testScore := 100
	Get(uid)
	Set(uid, testScore)
}

func init() {
	conn, _ := redis.Dial("127.0.0.1", "6379")
	RedisClient.conn = conn
}

func Set(uid string, score int) error {
	if score < scoreLimitMin || score > scoreLimitMax {
		return errors.New("score is error")
	}
	RedisClient.set(uid, score)
	return nil
}

func Get(uid string) interface{} {
	//TODO score 位移26 =>score
	return RedisClient.get(uid)
}

func (this *Client) set(uid string, score int) error {
	dbScore := score<<26 + (currentMonthEndTime - int(time.Now().Unix()))
	this.conn.Do("ZADD", dbScore, uid)
	return nil
}

func (this *Client) get(uid string) interface{} {
	res, err := this.conn.Do("ZRANK", uid)
	if err == nil {
		return nil
	}
	index := res.(int)
	start := index - sortLimit
	after := index + sortLimit
	rangeList, _ := this.conn.Do("ZRANGE", start, after)
	return rangeList
}
