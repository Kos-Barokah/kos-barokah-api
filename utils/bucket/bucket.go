package bucket

import (
	"bytes"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	// "strconv"
	// "time"
	configs "kos-barokah-api/configs"
	"kos-barokah-api/helper"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	maxPartSize = int64(5 * 1024 * 1024)
	maxRetries  = 3
)

var (
	awsAccessKeyID     string
	awsSecretAccessKey string
	awsBucketRegion    string
	awsBucketEndpoint  string
	awsBucketName      string
)

type BucketInterface interface {
	UploadImageHelper(file multipart.FileHeader) (*string, error)
	DeleteFileHelper(fileName string) (bool, error)
}

type Bucket struct {
	cfg configs.ProgrammingConfig
}

func InitBucket(config configs.ProgrammingConfig) BucketInterface {
	return &Bucket{
		cfg: config,
	}
}

// func (bct *Bucket) generateRandomCode(length int) string {
// 	const charset = "0123456789"
// 	code := make([]byte, length)
// 	for i := range code {
// 		code[i] = charset[rand.Intn(len(charset))]
// 	}

// 	return string(code)
// }

func (bct *Bucket) UploadImageHelper(file multipart.FileHeader) (*string, error) {
	if file.Size > maxPartSize {
		return nil, errors.New("you have suceeded max file size")
	}

	config := configs.InitConfig()
	if config == nil {
		return nil, errors.New("failed to load configuration")
	}

	awsAccessKeyID = config.BucketAccessKeyID
	awsSecretAccessKey = config.BucketSecretAccessKey
	awsBucketRegion = config.BucketRegion
	awsBucketEndpoint = config.BucketEndpoint
	awsBucketName = config.BucketName

	creds := credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, "")
	_, err := creds.Get()
	if err != nil {
		fmt.Printf("bad credentials: %s", err)
		errMsg := fmt.Errorf("bad credential: %v", err)
		return nil, errMsg
	}

	cfg := aws.NewConfig().
		WithRegion(awsBucketRegion).
		WithEndpoint(awsBucketEndpoint).
		WithCredentials(creds)

	sessionAws, err := session.NewSession(cfg)

	if err != nil {
		return nil, err
	}

	svc := s3.New(sessionAws)

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	size := file.Size
	buffer := make([]byte, size)
	src.Read(buffer)

	fileType := http.DetectContentType(buffer)
	fileUniqueness, _ := helper.Generate(helper.RandomString(15))
	fileName := fmt.Sprintf("IMG%s", fileUniqueness)

	path := "/media/" + fileName

	input := &s3.CreateMultipartUploadInput{
		Bucket:      aws.String(awsBucketName),
		Key:         aws.String(path),
		ContentType: aws.String(fileType),
		ACL:         aws.String("public-read"),
	}

	resp, err := svc.CreateMultipartUpload(input)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println("[BUCKET] Created multipart upload request : ", time.Now().Local())

	var (
		curr, partLength int64
		completedParts   []*s3.CompletedPart
		remaining        = size
		partNumber       = 1
	)

	for curr = 0; remaining != 0; curr += partLength {
		if remaining < maxPartSize {
			partLength = remaining
		} else {
			partLength = maxPartSize
		}
		completedPart, err := uploadPart(svc, resp, buffer[curr:curr+partLength], partNumber)
		if err != nil {
			fmt.Println(err.Error())
			err := abortMultipartUpload(svc, resp)
			if err != nil {
				fmt.Println(err.Error())
			}
			return nil, err
		}
		remaining -= partLength
		partNumber++
		completedParts = append(completedParts, completedPart)
	}

	completeResponse, err := completeMultipartUpload(svc, resp, completedParts)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	uploadUrl := fmt.Sprintf("https://%v", *completeResponse.Location)

	return &uploadUrl, nil
}

func (bct *Bucket) DeleteFileHelper(fileName string) (bool, error) {
	config := configs.InitConfig()
	if config == nil {
		return false, errors.New("failed to load configuration")
	}

	awsAccessKeyID = config.BucketAccessKeyID
	awsSecretAccessKey = config.BucketSecretAccessKey
	awsBucketRegion = config.BucketRegion
	awsBucketEndpoint = config.BucketEndpoint
	awsBucketName = config.BucketName

	creds := credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, "")
	_, err := creds.Get()
	if err != nil {
		fmt.Printf("bad credentials: %s", err)
		errMsg := fmt.Errorf("bad credential: %v", err)
		return false, errMsg
	}

	cfg := aws.NewConfig().WithRegion(awsBucketRegion).WithEndpoint(awsBucketEndpoint).WithCredentials(creds)
	sessionAws, err := session.NewSession(cfg)
	if err != nil {
		return false, err
	}
	svc := s3.New(sessionAws)

	input := &s3.DeleteObjectInput{
		Bucket: aws.String(awsBucketName), // s3 bucket name
		Key:    aws.String(fileName),      // file name
	}

	_, err = svc.DeleteObject(input)
	if err != nil {
		os.Exit(1)
	}

	return true, nil
}

func completeMultipartUpload(svc *s3.S3, resp *s3.CreateMultipartUploadOutput, completedParts []*s3.CompletedPart) (*s3.CompleteMultipartUploadOutput, error) {
	completeInput := &s3.CompleteMultipartUploadInput{
		Bucket:   resp.Bucket,
		Key:      resp.Key,
		UploadId: resp.UploadId,
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: completedParts,
		},
	}
	return svc.CompleteMultipartUpload(completeInput)
}

func uploadPart(svc *s3.S3, resp *s3.CreateMultipartUploadOutput, fileBytes []byte, partNumber int) (*s3.CompletedPart, error) {
	tryNum := 1
	partInput := &s3.UploadPartInput{
		Body:          bytes.NewReader(fileBytes),
		Bucket:        resp.Bucket,
		Key:           resp.Key,
		PartNumber:    aws.Int64(int64(partNumber)),
		UploadId:      resp.UploadId,
		ContentLength: aws.Int64(int64(len(fileBytes))),
	}

	for tryNum <= maxRetries {
		uploadResult, err := svc.UploadPart(partInput)
		if err != nil {
			if tryNum == maxRetries {
				if aerr, ok := err.(awserr.Error); ok {
					return nil, aerr
				}
				return nil, err
			}
			fmt.Printf("[BUCKET] Retrying to upload part #%v\n", partNumber)
			tryNum++
		} else {
			fmt.Printf("[BUCKET] Uploaded part #%v\n", partNumber)
			return &s3.CompletedPart{
				ETag:       uploadResult.ETag,
				PartNumber: aws.Int64(int64(partNumber)),
			}, nil
		}
	}
	return nil, nil
}

func abortMultipartUpload(svc *s3.S3, resp *s3.CreateMultipartUploadOutput) error {
	fmt.Println("[BUCKET] Aborting multipart upload for Upload ID: ", *resp.UploadId)
	abortInput := &s3.AbortMultipartUploadInput{
		Bucket:   resp.Bucket,
		Key:      resp.Key,
		UploadId: resp.UploadId,
	}
	_, err := svc.AbortMultipartUpload(abortInput)
	return err
}
