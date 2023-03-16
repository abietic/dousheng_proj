package config

type MinioConfigStruct struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	Region          string
	MinioAccessUrl  string
}
