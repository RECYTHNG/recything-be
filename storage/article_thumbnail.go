package storage

import (
	"context"
	"encoding/base64"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func UploadThumbnail(image *multipart.FileHeader) (string, error) {
	ctx := context.Background()

	// Decode Google Cloud credentials from Base64
	encodedCredentials := os.Getenv("STORAGE_KEY")
	decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
	if err != nil {
		logrus.Error("Gagal melakukan decode pada Google Cloud Credentials:", err)
		return "", err
	}

	client, err := storage.NewClient(ctx, option.WithCredentialsJSON(decodedCredentials))
	if err != nil {
		logrus.Error("Gagal membuat client GCP:", err)
		return "", err
	}
	defer client.Close()

	bucketName := "article_thumbnail"
	extension := filepath.Ext(image.Filename)
	allowedExtensions := map[string]bool{".jpg": true, ".png": true}

	if !allowedExtensions[strings.ToLower(extension)] {
		return "", errors.New("error : format file tidak diizinkan")
	}

	imagePath := "thumbnail/" + uuid.New().String() + extension

	wc := client.Bucket(bucketName).Object(imagePath).NewWriter(ctx)
	defer wc.Close()

	file, err := image.Open()
	if err != nil {
		logrus.Error("gagal membuka file:", err)
		return "", err
	}

	if _, err := io.Copy(wc, file); err != nil {
		logrus.Error("gagal menyalin file:", err)
		return "", err
	}

	imageURL := "https://storage.googleapis.com/" + bucketName + "/" + imagePath

	return imageURL, nil
}
