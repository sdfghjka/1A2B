package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "backend/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UploadImageToImageService(filePath string) (string, error) {
	imageData, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("讀檔錯誤: %v", err)
	}
	conn, err := grpc.Dial("host.docker.internal:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", fmt.Errorf("建立 gRPC 連線失敗: %v", err)
	}
	defer conn.Close()

	client := pb.NewImageUploaderClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.UploadImage(ctx, &pb.UploadImageRequest{
		ImageData: imageData,
		Filename:  "uploaded_from_go.jpg",
	})
	if err != nil {
		return "", fmt.Errorf("gRPC 呼叫失敗: %v", err)
	}

	log.Printf("上傳成功，圖片網址：%s", res.ImageUrl)
	return res.ImageUrl, nil
}
