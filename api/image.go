package api

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

var bucketName = "imagera"

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	})
	uploader := s3.New(sess)

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer f.Close()

	// Read the file into a buffer
	size := file.Size
	buffer := make([]byte, size)
	f.Read(buffer)

	_, err = uploader.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(file.Filename),
		Body:   bytes.NewReader(buffer),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}
	imageURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, file.Filename)
	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully", "url": imageURL})
}

// Dummy function to return a list of images
func ListImages(c *gin.Context) {
	// You may retrieve image metadata from a database and return it here.
	// For now, return a dummy list of image URLs.
	images := []string{
		"https://imagera.s3.amazonaws.com/image1.png",
		"https://imagera.s3.amazonaws.com/image2.png",
		"https://imagera.s3.amazonaws.com/image3.png",
	}

	c.JSON(http.StatusOK, gin.H{
		"images": images,
	})
}

// DownloadImage provides a direct URL for image download
func DownloadImage(c *gin.Context) {
	imageID := c.Param("imageId")

	// Create a direct link to the image in S3
	imageURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, imageID)

	// Return the image URL
	c.JSON(http.StatusOK, gin.H{
		"image_url": imageURL,
	})
}
