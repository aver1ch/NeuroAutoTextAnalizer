package config

import (
	"os"

	errors "github.com/kerilOvs/backend/internal/errorsExt"

	yaml "gopkg.in/yaml.v3"
)

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type MinioConfig struct {
	Endpoint  string `yaml:"endpoint" env:"MINIO_ENDPOINT"`
	AccessKey string `yaml:"access_key" env:"MINIO_ACCESS_KEY"`
	SecretKey string `yaml:"secret_key" env:"MINIO_SECRET_KEY"`
	Bucket    string `yaml:"bucket" env:"MINIO_BUCKET"`
	UseSSL    bool   `yaml:"use_ssl" env:"MINIO_USE_SSL"`
}

type RabbitConfig struct {
	Url            string `yaml:"endpoint" env:"MINIO_ENDPOINT"`
	QueuePhotoName string `yaml:"queue_photo_name" env:"RABBIT_PHOTO_NAME"`
	QueueTagsName  string `yaml:"queue_tags_name" env:"RABBIT_TAGS_NAME"`
	QueueAnketName string `yaml:"queue_anket_name" env:"RABBIT_ANKET_NAME"`
}

type LogConfig struct {
	LogLevel  string `yaml:"endpoint" env:"MINIO_ENDPOINT"`
	LogFormat string `yaml:"queue_photo_name" env:"RABBIT_PHOTO_NAME"`
}

type Config struct {
	Database DBConfig     `yaml:"database"`
	Server   ServerConfig `yaml:"server"`
	Minio    MinioConfig  `yaml:"minio"`
	Rabbit   RabbitConfig `yaml:"rabbit"`
}

func ReadConfig() (Config, error) {

	var config Config

	data, err := os.ReadFile(fileName)
	if err != nil {
		return config, errors.ErrorLocate(err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, errors.ErrorLocate(err)
	}
	return config, nil
}
