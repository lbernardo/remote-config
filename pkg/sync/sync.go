package sync

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Sync struct {
	client      *s3.Client
	bucket      string
	environment string
	project     string
	configFile  string
	namespace   string
}

func New() *Sync {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)
	return &Sync{
		client:      client,
		bucket:      os.Getenv("GO_REMOTE_CONFIG_BUCKET"),
		environment: os.Getenv("ENVIRONMENT"),
		project:     os.Getenv("PROJECT"),
		namespace:   os.Getenv("NAMESPACE"),
		configFile:  "config.yaml",
	}
}

func (s *Sync) SetBucket(bucket string) *Sync {
	s.bucket = bucket
	return s
}

func (s *Sync) SetProject(project string) *Sync {
	s.project = project
	return s
}

func (s *Sync) SetEnvironment(environment string) *Sync {
	s.environment = environment
	return s
}

func (s *Sync) SetConfigFile(configFile string) *Sync {
	s.configFile = configFile
	return s
}

func (s *Sync) SetNamespace(namespace string) *Sync {
	s.namespace = namespace
	return s
}

func (s *Sync) Sync() {
	key := fmt.Sprintf("%v/%v/%v/%v", s.project, s.environment, s.namespace, s.configFile)
	object, err := s.client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Fatalf("error to get object '%v' : %v\n", key, err.Error())
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(object.Body)

	if err := os.WriteFile(s.configFile, buf.Bytes(), 0755); err != nil {
		log.Fatalf("error to create config.yaml in root: %v", err.Error())
	}
	//defer os.Remove(s.configFile)

	viper.SetConfigFile(s.configFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error to read config: %v", err.Error())
	}

}
