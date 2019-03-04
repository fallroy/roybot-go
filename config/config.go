package config

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/line/line-bot-sdk-go/linebot"

	yaml "gopkg.in/yaml.v2"
)

type Database struct {
	Url          string
	MaxIdleConns int
	MaxOpenConns int
}

type Release struct {
	Version string
	Time    string
}

type Linebot struct {
	AdminID string `yaml:"adminID"`
	Channel struct {
		Secret string
		Token  string
	}
}

type Config struct {
	Database   Database
	Release    Release
	ListenAddr string `yaml:"listenAddr"`
	Linebot    Linebot
}

var DB *sql.DB
var Conf *Config
var Bot *linebot.Client

//Init is func
func init() {
	initConf()
	initDB()
	initLinebot()
	fmt.Println("Init finished...")
}

func initConf() {
	fmt.Println("Init Conf")
	data, err := ioutil.ReadFile("config/config.yml")
	if err != nil {
		log.Panicf("yamlFile.Get err.\n%+v ", err)
	}
	// fmt.Println(string(data))
	yaml.Unmarshal(data, &Conf)
	fmt.Printf("\n####\nConf : %+v\n###\n\n", Conf)
}

func initDB() {
	fmt.Println("Init DB")
	db, err := sql.Open("mysql", Conf.Database.Url)
	if err != nil {
		log.Panicf("DB init Get err.\n%+v ", err)
	} else {
		db.SetMaxIdleConns(Conf.Database.MaxIdleConns)
		db.SetMaxOpenConns(Conf.Database.MaxOpenConns)
		DB = db
	}
}

func initLinebot() {
	fmt.Println("Init Linebot")
	bot, err := linebot.New(Conf.Linebot.Channel.Secret, Conf.Linebot.Channel.Token)
	if err != nil {
		log.Panicf("Linebot init Get err.\n%+v ", err)
	} else {
		Bot = bot
		_, err := Bot.PushMessage(Conf.Linebot.AdminID, linebot.NewTextMessage("I'm reborn!!")).Do()
		if err != nil {
			log.Panicf("Linebot Send message Get err.\n%+v ", err)
		}
	}
}
