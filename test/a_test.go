package test

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
)

func Test1(t *testing.T) {
	op()
	//aa(3)
	//ar := []int{1,2,4,4,5}
	//ret := upper_bound_(5,4,ar)
	//fmt.Println(ret)
}

func aa(b int) int {
	fmt.Println(b)
	return b * b
}

func op() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	err := client.Do("Set", "xxx", 100).Err()
	if err != nil {
		panic("set v error")
	}
	//获取值不存在时
	data, err := client.Do("Get", "bbb").Result()
	if err == redis.Nil {
		fmt.Println("xxx does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("xxx:>", data)
	}

	if err := client.Do("expire", "xxx", 10).Err(); err != nil {
		panic("error")
	}
	if err3 := client.Do("Mset", "aa", 123, "bb", 234, "cc", 456).Err(); err3 != nil {
		panic("mset error")
	}
	data1, err4 := client.Do("Mget", "aa", "bb", "cc").Result()
	if err4 != nil {
		panic("mget error")
	}
	//直接调用方法  设置值 .Err；获取值 .Result
	data2, _ := client.MGet("aa", "bb", "cc").Result()
	fmt.Println(data1, data2)

	//队列
	if eee := client.LPush("mylll", 1, 2, "ss", "zz").Err(); eee != nil {
		panic("入队错误")
	}
	data4, err5 := client.LRange("mylll", 0, -1).Result()
	if err5 == redis.Nil {
		fmt.Println("mylll does not exist")
	}
	if err5 != nil {
		panic("mylll err")
	}
	fmt.Println(data4)

	//hash
	//client.HMSet()
	if err6 :=client.Do("hset","book","name","zhs","aa",1,"bb",2).Err();err6 != nil {
		panic("hashset err")
	}
	data5,err7 := client.HGetAll("book").Result()

	if err7 == redis.Nil {
		fmt.Println("book does not exist")
	}
	if err7 != nil {
		panic("hget error")
	}
	fmt.Println(data5)


}


func upper_bound_( n int ,  v int ,  a []int ) int {
	// write code here
	left := 0
	right  := n-1
	for left < right {
		mid  := left + (right-left)/2
		if a[mid] == v  {
			return mid+1
		}else{
			if a[mid]>v {
				right--
			}else{
				left++
			}
		}
	}
	return n+1
}