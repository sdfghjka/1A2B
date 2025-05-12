const fs = require("fs");
const path = require("path");
const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");
const { imgurFileHandler } = require("./helpers/uploader");

// 載入 proto 檔
const PROTO_PATH = path.join(__dirname, "proto", "image.proto");

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

const imageProto = grpc.loadPackageDefinition(packageDefinition).image;

function uploadImage(call, callback) {
  const { image_data, filename } = call.request;

  // 儲存到暫存資料夾
  const tempPath = path.join(__dirname, "temp", filename);
  fs.writeFile(tempPath, image_data, async (err) => {
    if (err) {
      console.error("寫入圖片失敗:", err);
      return callback({ code: grpc.status.INTERNAL, message: "寫入失敗" });
    }

    try {
      const imageUrl = await imgurFileHandler({ path: tempPath });

      fs.unlink(tempPath, () => {}); // 上傳完畢後刪除暫存圖片

      if (!imageUrl) {
        return callback(null, { image_url: "" });
      }

      return callback(null, { image_url: imageUrl });
    } catch (err) {
      console.error("Imgur 上傳失敗:", err);
      return callback({ code: grpc.status.INTERNAL, message: "Imgur 錯誤" });
    }
  });
}

function main() {
  const server = new grpc.Server();
  server.addService(imageProto.ImageUploader.service, { UploadImage: uploadImage });
  server.bindAsync("0.0.0.0:50051", grpc.ServerCredentials.createInsecure(), () => {
    server.start();
    console.log("gRPC server running at http://localhost:50051");
  });
}

main();
