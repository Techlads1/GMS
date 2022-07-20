package config

import (
	"fmt"
	"gateway/package/log"
	"gateway/package/util"
	"os"

	"github.com/spf13/viper"
)

// Config defines configurations to be exportes
type Config struct {
	WebServer   WebServerConfig
	DicomServer DicomServerConfig
	Database    DatabaseConfig
	Secret      SecretKey
}

// ServerConf defines server configurations
type WebServerConfig struct {
	Host string
	Port int
}

// DicomServerConfig defines server configurations
type DicomServerConfig struct {
	Host        string
	Port        int32
	AETitle     string
	StoragePath string
}

// DatabaseConf defines database configuration
type DatabaseConfig struct {
	Name     string
	User     string
	Password string
}

type SecretKey struct {
	Secret string
}

var connectionString string

// New configuration
// inputs filePath should contain path with file name and file extension e.g ./storage/config.yml
func New() (*Config, error) {
	//assumed cofiguration path
	configFile := "config.yml"
	confPath, err := os.Getwd()
	if util.CheckError(err) {
		log.Error("error getting a working directory:%v", err)
		return nil, err
	}
	configPath := fmt.Sprintf("%s/%s", confPath, configFile)

	viper.SetConfigFile(configPath) //file name with extension
	//viper.AddConfigPath(filePath) //config file apth
	viper.AutomaticEnv() //enable reading environmental variables

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return nil, err
	}

	cfg := Config{}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return nil, err
	}
	return &cfg, nil
}

func (c *Config) GetDatabaseConnection() string {
	conn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%d", c.WebServer.Host, c.Database.Name, c.Database.User, c.Database.Password, c.WebServer.Port)
	return conn
}

func LoggerPath() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting file path %s\n", err)
	}
	return path + "/.logs"
}

func (c *Config) GetSecret() string {
	return c.Secret.Secret
}

//LogoPath returns logo path
func LogoPath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Error(err)
		return "", err
	}
	path = fmt.Sprintf("%s/webserver/public/images/dit-logo.png", path)
	return path, nil
}

//ReportDir returns .storage/reports path
func ReportDir() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Error(err)
		return "", err
	}
	// path = fmt.Sprintf("%s\\.storage\\reports", path)
	path = fmt.Sprintf("%s/.storage/reports/", path)
	return path, nil
}

//UploadsDir returns .storage/uploads
func UploadsDir() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Error(err)
		return "", err
	}
	path = fmt.Sprintf("%s/.storage/uploads/", path)
	return path, nil
}

//DownloadsDir returns .storage/downloads path
func DownloadsDir() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Error(err)
		return "", err
	}
	path = fmt.Sprintf("%s/.storage/downloads/", path)
	return path, nil
}

//DicomDir returns .storage/dicom/ path
func DicomDir() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		log.Error(err)
		return "", err
	}
	path = fmt.Sprintf("%s/.storage/dicom/", path)
	return path, nil
}
