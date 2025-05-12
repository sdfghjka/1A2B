const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const fs = require('fs');
const path = require('path');

// 載入 proto 檔案
const PROTO_PATH = path.join(__dirname, '..', 'proto', 'image.proto');

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});
const imageProto = grpc.loadPackageDefinition(packageDefinition).image;


const client = new imageProto.ImageUploader('localhost:50051', grpc.credentials.createInsecure());

const imagePath = path.join(__dirname, '..', 'upload', 'img01.jpg');

const imageData = fs.readFileSync(imagePath);  

client.UploadImage({ image_data: imageData, filename: 'img01.jpg' }, (error, response) => {
  if (error) {
    console.error('gRPC 請求錯誤:', error);
  } else {
    console.log('上傳成功，圖片 URL:', response.image_url);
  }
});

