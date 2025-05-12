const { imgurFileHandler } = require('./helpers/uploader'); 
const upload = require('./middleware/multer'); 


const uploadImage = async (req, res) => {
  upload.single('image')(req, res, async (err) => {
    if (err) return res.status(400).json({ message: '上傳圖片時出現錯誤' });

    try {
      const imgLink = await imgurFileHandler(req.file);

      if (!imgLink) {
        return res.status(500).json({ message: '無法儲存圖片' });
      }

      return res.status(200).json({ imageUrl: imgLink });
    } catch (error) {
      console.error('Error during image upload:', error);
      return res.status(500).json({ message: '無法儲存圖片到 Imgur' });
    }
  });
};

module.exports = { uploadImage };
