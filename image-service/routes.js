const express = require('express');
const { uploadImage } = require('./uploadController');  

const router = express.Router();


router.post('/image', uploadImage);  

module.exports = router;
