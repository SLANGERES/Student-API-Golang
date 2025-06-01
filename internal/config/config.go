package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpConfig struct {
	Addr string
}
//Setting up the configuration file
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HttpConfig  `yaml:"http_server"`
}


//Config function that must be run first 

func MustDone() *Config{
	var configPath string

	//Checking the env for the config file path
	configPath=os.Getenv("CONFIG_PATH")

	if configPath==""{
		//checking if the path of config file present on the flags or not 
		flags:=flag.String("config","","path to the configuration file")
		flag.Parse()
		configPath=*flags

		//if flag empty return fatal error
		if configPath==""{
			log.Fatal("Config file needed")
		}
	}
	//if the file not exist according to either env or flag
	if _,err:=os.Stat(configPath);os.IsNotExist(err){
		log.Fatalf("Config file does not exist:%s",configPath)


	}
	//making the config instance 
	var cfg Config

	//Seralize the config file 
	err:=cleanenv.ReadConfig(configPath,&cfg)

	if err!=nil{	
		log.Fatal("Unable to parse the config data")
	}

	//return the config file by reference
	return &cfg

}
