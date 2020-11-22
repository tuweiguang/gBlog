package config

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

type DBConfig struct {
	Title string
	Db    []*DB `toml:"db"`
}
type DB struct {
	DbType       string `toml:"dbtype"`
	Host         string `toml:"host"`
	Port         int    `toml:"port"`
	User         string `toml:"user"`
	Password     string `toml:"password"`
	DbName       string `toml:"dbName"`
	DbCharset    string `toml:"dbCharset"`
	Active       int    `toml:"active"`
	Idle         int    `toml:"idle"`
	IdleTimeout  string `toml:"idleTimeout"`
	QueryTimeout string `toml:"queryTimeout"`
	ExecTimeout  string `toml:"execTimeout"`
	TranTimeout  string `toml:"tranTimeout"`
	DbNum        int    `toml:"dbNum"`
}

func (db *DB) String() string {
	return fmt.Sprintf("{DbType:%v,Host:%v,Port:%v,User:%v,Password:%v,DbName:%v,DbCharset:%v,Active:%v,Idle:%v,IdleTimeout:%v,QueryTimeout:%v,ExecTimeout:%v,TranTimeout:%v,DbNum:%v}",
		db.DbType, db.Host, db.Port, db.User, db.Password, db.DbName, db.DbCharset, db.Active, db.Idle, db.IdleTimeout, db.QueryTimeout, db.ExecTimeout, db.TranTimeout, db.DbNum)
}

var db *DBConfig

type APPConfig struct {
	Title string
	App   *APP `toml:"app"`
}

type APP struct {
	HttpAddr      string `toml:"httpAddr"`
	SessionExpire int    `toml:"sessionExpire"`
	LogMode       bool   `toml:"logMode"`
}

var app *APPConfig

type LOGConfig struct {
	Title string
	Log   *LOG `toml:"log"`
}

type LOG struct {
	Path  string `toml:"path"`
	Level string `toml:"level"`
}

var logConfig *LOGConfig

var confPath = flag.String("conf", "./conf", "配置文件目录")

func Init() {
	dbFile := *confPath + "/db.toml"
	log.Printf("config init ,dbFile %v", dbFile)
	db = new(DBConfig)
	_, err := toml.DecodeFile(dbFile, db)
	if err != nil {
		panic("加载db配置文件错误:" + err.Error())
	}
	log.Printf("config init, DBConfig:%v", db.Db)

	logFile := *confPath + "/log.toml"
	logConfig = new(LOGConfig)
	_, err = toml.DecodeFile(logFile, logConfig)
	if err != nil {
		panic("加载log配置文件错误:" + err.Error())
	}
	log.Printf("config init, LOGConfig:%v", logConfig)

	appFile := *confPath + "/app.toml"
	app = new(APPConfig)
	_, err = toml.DecodeFile(appFile, app)
	if err != nil {
		panic("加载app配置文件错误:" + err.Error())
	}
	log.Printf("config init, APPConfig:%v", app)
}

func GetDBConfig() []*DB {
	if db.Db == nil {
		return []*DB{}
	}
	return db.Db
}

func GetAPPConfig() *APP {
	if app.App == nil {
		return &APP{}
	}
	return app.App
}

func GetLOGConfig() *LOG {
	if logConfig.Log == nil {
		return &LOG{}
	}
	return logConfig.Log
}
