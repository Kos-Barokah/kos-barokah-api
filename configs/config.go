package configs

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type ProgrammingConfig struct {
	ServerPort            int
	Environment           string
	DBPort                uint16
	DBHost                string
	DBUser                string
	DBPass                string
	DBName                string
	Secret                string
	RefSecret             string
	BaseURL               string
	BucketAccessKeyID     string
	BucketSecretAccessKey string
	BucketRegion          string
	BucketEndpoint        string
	BucketName            string
}

func InitConfig() *ProgrammingConfig {
	var res = new(ProgrammingConfig)
	err := godotenv.Load(".env")

	if err != nil {
		err := godotenv.Load(".env.staging")

		if err != nil {
			return nil
		}

	}

	res, errorRes := loadConfig()

	logrus.Error(errorRes)
	if res == nil {
		logrus.Error("Config: Cannot start program, failed to load configuration")
		return nil
	}

	return res
}

func ReadData() *ProgrammingConfig {
	var data = new(ProgrammingConfig)
	data, _ = loadConfig()

	if data == nil {
		err := godotenv.Load(".env")
		data, errorData := loadConfig()

		fmt.Println(errorData)

		if err != nil || data == nil {
			return nil
		}
	}
	return data
}

func loadConfig() (*ProgrammingConfig, error) {
	var error error
	var res = new(ProgrammingConfig)
	var permit = true

	if val, found := os.LookupEnv("SERVER"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config: Invalid port value,", err.Error())
			permit = false
		}
		res.ServerPort = port
	} else {
		permit = false
		error = errors.New("Port undefined")
	}

	if val, found := os.LookupEnv("DB_PORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : Invalid port value,", err.Error())
			permit = false
		}
		res.DBPort = uint16(port)
	} else {
		permit = false
		error = errors.New("DB Port undefined")
	}

	if val, found := os.LookupEnv("DB_HOST"); found {
		res.DBHost = val
	} else {
		permit = false
		error = errors.New("DB Host undefined")
	}

	if val, found := os.LookupEnv("DB_USER"); found {
		res.DBUser = val
	} else {
		permit = false
		error = errors.New("DB User undefined")
	}

	if val, found := os.LookupEnv("DB_PASS"); found {
		res.DBPass = val
	} else {
		// permit = false
		// error = errors.New("DB Pass undefined")
		res.DBPass = ""
	}

	if val, found := os.LookupEnv("DB_NAME"); found {
		res.DBName = val
	} else {
		permit = false
		error = errors.New("DB Name undefined")
	}

	if val, found := os.LookupEnv("BASE_URL"); found {
		res.BaseURL = val
	} else {
		res.BaseURL = ""
		// permit = false
		// error = errors.New("BASE_URL undefined")
	}
	if val, found := os.LookupEnv("BUCKET_ACCESS_KEY_ID"); found {
		res.BucketAccessKeyID = val
	} else {
		permit = false
		error = errors.New("Config : Invalid BUCKET ACCESS KEY ID undefined")
	}

	if val, found := os.LookupEnv("BUCKET_SECRET_ACCESS_KEY"); found {
		res.BucketSecretAccessKey = val
	} else {
		permit = false
		error = errors.New("Config : Invalid BUCKET SECRET ACCESS KEY undefined")
	}

	if val, found := os.LookupEnv("BUCKET_REGION"); found {
		res.BucketRegion = val
	} else {
		permit = false
		error = errors.New("Config : Invalid BUCKET REGION undefined")
	}

	if val, found := os.LookupEnv("BUCKET_ENDPOINT"); found {
		res.BucketEndpoint = val
	} else {
		permit = false
		error = errors.New("Config : Invalid BUCKET ENDPOINT undefined")
	}

	if val, found := os.LookupEnv("BUCKET_NAME"); found {
		res.BucketName = val
	} else {
		permit = false
		error = errors.New("Config : Invalid BUCKET NAME undefined")
	}

	if !permit {
		return nil, error
	}

	return res, nil
}
