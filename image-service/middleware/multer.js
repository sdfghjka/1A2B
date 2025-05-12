const multer = require('multer');

// 設定儲存路徑及檔名規則
const upload = multer({ dest: 'uploads/' });  // 儲存圖片的資料夾

module.exports = upload;
