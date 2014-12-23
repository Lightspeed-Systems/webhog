package webhog

import (
	"io/ioutil"
	"launchpad.net/goamz/aws"
	"launchpad.net/goamz/s3"
	"strings"
)

func UploadEntity(dir string, entity *Entity) (string, error) {
	spl := strings.Split(dir, "/")
	endDir := spl[len(spl)-1]

	// auth, err := aws.EnvAuth()
	// if err != nil {
	// 	return "", err
	// }

	region := aws.USWest2
	if Config.AwsRegion == "USWest2" {
		region = aws.USWest2
	}

	// Open Bucket
	s := s3.New(aws.Auth{Config.AwsKey, Config.AwsSecret}, region)
	bucket := s.Bucket(Config.bucket)

	b, err := ioutil.ReadFile(dir)
	if err != nil {
		return "", err
	}

	err = bucket.Put("/"+endDir, b, "text/plain", s3.PublicRead)
	if err != nil {
		return "", err
	}

	awsLink := bucket.URL("/" + endDir)

	return awsLink, err
}
