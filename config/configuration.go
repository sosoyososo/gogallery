package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
	About    AboutConfiguration
	Gallery  GalleryConfiguration
	IG       InstagramConfiguration
}

type DatabaseConfiguration struct {
	Baseurl string
}

type ServerConfiguration struct {
	Port  string
	Debug bool
}

type InstagramConfiguration struct {
	Username string
	Password string
	Enable   bool
	SyncRate int
}

type GalleryConfiguration struct {
	Name             string
	Basepath         string
	Url              string
	Theme            string
	ImagesPerPage    int
	QueThreshold     int
	AlbumBlacklist   []string
	PictureBlacklist []string
	Renderer         string
}

type AboutConfiguration struct {
	Enable          bool
	Twitter         string
	Facebook        string
	Email           string
	Instagram       string
	Description     string
	Footer          string
	Photographer    string
	ProfilePhoto    string
	BackgroundPhoto string
	Blog            string
	Website         string
}

var config = Configuration{}

func LoadConfig() *Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("GLLRY")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return &config
}

func (c *AboutConfiguration) Save() {
	log.Println("Saving About Config")
	viper.Set("about", c)
	viper.WriteConfig()
	config.About = *c
}
func (c *GalleryConfiguration) Save() {
	log.Println("Saving Gallery Config")
	viper.Set("gallery", c)
	viper.WriteConfig()
	config.Gallery = *c
}
