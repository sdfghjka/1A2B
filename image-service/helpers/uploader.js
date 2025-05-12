const imgur = require('imgur');  // 使用 imgur 套件上傳圖片

// 如果不需要 API Key，則設定 null；如果需要，請將其設為你從 imgur 獲得的 Client ID
imgur.setClientId(null);  // 如果不需要 API Key

const imgurFileHandler = (file) => {
  return new Promise((resolve, reject) => {
    if (!file) return resolve(null);

    // 使用 imgur 上傳檔案
    imgur.uploadFile(file.path)
      .then(img => {
        if (img && img.link) {
          resolve(img.link);  // 回傳圖片的 URL
        } else {
          resolve(null);
        }
      })
      .catch(err => {
        console.error("Imgur Upload Error:", err);
        reject("無法儲存圖片到 Imgur");
      });
  });
};

module.exports = { imgurFileHandler };
