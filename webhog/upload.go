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

	region := aws.USEast
	switch Config.AwsRegion {
	case "us-east-1":
		region = aws.USEast
	case "us-west-2":
		region = aws.USWest2
	case "us-west-1":
		region = aws.USWest
	case "eu-west-1":
		region = aws.EUWest
	case "ap-southeast-1":
		region = aws.APSoutheast
	case "ap-southeast-2":
		region = aws.APSoutheast2
	case "ap-northeast-2":
		region = aws.APNortheast
	case "sa-east-1":
		region = aws.SAEast
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
