package config

import (
	"io/ioutil"
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

// PostgresConfig ...
type PostgresConfig struct {
	DriverName                 string `yaml:"driver_name"`
	DataSource                 string `yaml:"data_source"`
	MaxOpenConns               int    `yaml:"max_open_conns"`
	MaxIdleConns               int    `yaml:"max_idle_conns"`
	ConnMaxLifeTimeMiliseconds int64  `yaml:"conn_max_life_time_ms"`
	MigrationConnURL           string `yaml:"migration_conn_url"`
	IsDevMode                  bool   `yaml:"is_dev_mode"`
}

// KafkaConfig ...
type KafkaConfig struct {
	Brokers  []string `yaml:"brokers"`
	Version  string   `yaml:"version"`
	ClientID string   `yaml:"client_id"`
	RackID   string   `yaml:"rack_id"`
	GroupID  string   `yaml:"group_id"`
}

// WorkerConfig ...
type WorkerConfig struct {
	Interval  int `yaml:"interval"`
	BatchSize int `yaml:"batch_size"`
}

// AppConfig ...
type AppConfig struct {
	DB     *PostgresConfig `yaml:"db"`
	Kafka  *KafkaConfig    `yaml:"kafka"`
	Worker *WorkerConfig   `yaml:"worker"`
}

func Load(filePath string) (*AppConfig, error) {
	if len(filePath) == 0 {
		filePath = os.Getenv("CONFIG_FILE")
	}

	fields := []interface{}{
		"func",
		"config.readFromFile",
		"filePath",
		filePath,
	}

	sugar := zap.S().With(fields...)

	sugar.Debug("Load config...")
	zap.S().Debugf("CONFIG_FILE=%v", filePath)

	configBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		sugar.Error("Failed to load config file")
		return nil, err
	}
	configBytes = []byte(os.ExpandEnv(string(configBytes)))

	cfg := &AppConfig{}

	err = yaml.Unmarshal(configBytes, cfg)
	if err != nil {
		sugar.Error("Failed to parse config file")
		return nil, err
	}
	zap.S().Debugf("config: %+v", cfg)
	zap.S().Debug("======================================")
	zap.S().Debugf("database config: %+v", cfg.DB)

	return cfg, nil
}
