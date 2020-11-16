package service


import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"singo/model"
	"singo/serializer"
	"strconv"
)


/**
商品秒杀
1.下订单，更新库存 判断库存-1>0，则更新成功 【mysql：qps 500-600】
2.增加redis缓存，将商品的库存添加到redis中，则每次更新先更新redis,然后异步更新mysql 【redis qps:万级别】
做法：a、将商品key_id作为key,商品数量作为value存入redis中
     b、先创建订单,然后对redis中商品的库存进行减一操作
     c、更新MySQL中的库存数量
缺点:
	a、当更新MySQL失败时，会导致MySQL中的数据和redis中的数据不一致问题
	b、当库存变为0时，不能再进行下单操作
解决：
	a.当更新mysql失败时，则对redis中对应的商品key进行加1操作
    b.当redis中的库存变为0时,则不能进行下单操作

3.缓存升级：商品数量只有100个  请求10万+时，如何处理
做法：基于内存处理，当商品的库存减到0时,设置isEmpty[flag] = true ,每次请求时，先判断是否库存为0，库存为0时，直接返回
 */

var db *gorm.DB

//库存设置的key
const goodsIds = "prefix_key_goods_ids_"

//库存为0标志
var isEmpty = map[string]bool{"isFlag":false}

type OrderServer struct {
	GoodsId int `form:"goodsId" json:"goodsId" binding:"required"`
}

//TODO 每次先初始化缓存
func initCache() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//获取所有的商品

	type result struct {
		Id    string
		Stock string
	}
	var goods []result
	//temp_slice := make([]map[string]int,3)
	model.DB.Raw("SELECT id, stock FROM goods").Scan(&goods)
	//goods [切片]
	//fmt.Println(goods)

	//将商品的数量写入redis中
	for _, v := range goods {
		err := client.Set(goodsIds+v.Id, v.Stock, 0).Err()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (service *OrderServer) CreateAndUpdateStock() serializer.Response {
	//下订单
	//initCache()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//商品key_id
	key := goodsIds+strconv.Itoa(service.GoodsId)
	//fmt.Println(key)

	//在内存中先判断isFlag=true还是false 是true了则直接返回
	if isEmpty["isFlag"] {
		return  serializer.Response{200,"","商品已经卖完",""}
	}
	nums,err2 := rdb.Get(key).Result()
	if err2 != nil {
		fmt.Println(err2)
	}
	//字符串转整数
	num,_ := strconv.Atoi(nums)
	////小于0时 设置状态码
	if num<=0{
		//在内存中,设置状态码
		isEmpty["isFlag"] = true
	}

	//恢复库存为0
	decrBy := rdb.DecrBy(key,1)
	if decrBy.Val() < 0 {
		rdb.IncrBy(key,1)
	}
	//******************redis操作*****************************
	//设置值
	//err := rdb.Set("score",100,0).Err()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	////获取值
	//val,err2 := rdb.Get("score").Result()
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	//
	//fmt.Println(val,err2)
	//如何开启一个事务
	uuid := 123

	order := model.Order{
		GoodsId: service.GoodsId,
		Uid:     uuid,
	}
	//减少库存,按照商品id查找，并且物品的stock大于0时，才去更新库存
	var goods model.GoodsModel

	tx := model.DB.Begin()

	//rollback ----没有生效
	if err := tx.Where("id =? and stock>0", service.GoodsId).First(&goods).Update("stock", goods.Stock-1).Error; err != nil {
		tx.Rollback()
		return serializer.ParamErr("更新数据失败", err)
	}
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return serializer.ParamErr("用户下单失败", err)
	}

	//if err := tx.Rollback().Error; err != nil {
	//	_ = fmt.Errorf("Rollback should not raise error")
	//}
	tx.Commit()
	return serializer.Response{}

}
