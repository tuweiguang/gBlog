package db

import (
	"fmt"
	"gBlog/common/log"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

//官网地址 https://github.com/lionsoul2014/ip2region

type IP2Region struct{}

func (i *IP2Region) init(dbI SQL) interface{} {
	//ip2region.db是从官网下载的,地址https://github.com/lionsoul2014/ip2region/tree/master/data,需要定期更新
	region, err := ip2region.New("./conf/ip2region.db")
	if err != nil {
		log.Log.Error(fmt.Sprintf("init ip2region.db failed,err = %v", err.Error()))
		panic("init ip2region.db failed")
	}

	return &GIP2Region{region}
}

type GIP2Region struct {
	*ip2region.Ip2Region
}

func (g *GIP2Region) Select(ip string) *ip2region.IpInfo {
	//三种方式查询
	//ipInfo, err = g.BinarySearch("127.0.0.1") //线性不安全
	//ipInfo, err = g.BtreeSearch("127.0.0.1") //线性不安全
	ipInfo, err := g.MemorySearch(ip) // 该方法线性安全,查询最快
	if err != nil {
		log.GetLog().Error(fmt.Sprintf("GIP2Region select ip:%v error:%v", ip, err))
		return nil
	}
	return &ipInfo
}
