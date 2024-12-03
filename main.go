package main

import (
    "context"
    "fmt"
    "os"

    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
    bucket := "dfs-storage"
    key := "test.txt"
    file := "test.txt"

    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        fmt.Println("Error loading AWS config:", err)
        return
    }

    client := s3.NewFromConfig(cfg)

    f, err := os.Open(file)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer f.Close()

    _, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
        Bucket: &bucket,
        Key:    &key,
        Body:   f,
    })
    if err != nil {
        fmt.Println("Error uploading file:", err)
        return
    }

    fmt.Println("File uploaded successfully!")
}
