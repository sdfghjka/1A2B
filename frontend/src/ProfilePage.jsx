import { useEffect, useState } from "react";
import { useUser } from "./contexts/UserContext";
import Layout from "./components/Layout";

function ProfilePage() {
  const { user } = useUser();
  const [profile, setProfile] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  const [uploading, setUploading] = useState(false);
  const [uploadSuccess, setUploadSuccess] = useState("");
  const [uploadError, setUploadError] = useState("");
  const [uploadedImageUrl, setUploadedImageUrl] = useState("");

  useEffect(() => {
    if (!user || !user.user_id) {
      setLoading(false);
      setError("無法取得用戶資料，請重新登入");
      return;
    }

    const fetchProfile = async () => {
      const token = localStorage.getItem("token");
      if (!token) {
        setError("未登入，請重新登入");
        setLoading(false);
        return;
      }

      try {
        const res = await fetch(`http://localhost:3000/api/users/${user.user_id}`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        const data = await res.json();
        if (res.ok) {
          setProfile(data);
        } else {
          setError(data.error || "取得資料失敗");
        }
      } catch (err) {
        setError(err.message || "連線錯誤");
      } finally {
        setLoading(false);
      }
    };

    fetchProfile();
  }, [user]);

  const handleImageUpload = async (e) => {
    const token = localStorage.getItem("token");
    if (!token) {
      setError("未登入，請重新登入");
      setLoading(false);
      return;
    }

    const file = e.target.files[0];
    if (!file) return;

    const formData = new FormData();
    formData.append("image", file);
    formData.append("uploadType", "imgur");

    try {
      setUploading(true);
      setUploadSuccess("");
      setUploadError("");
      setUploadedImageUrl("");

      const res = await fetch("http://localhost:3000/api/users/upload/image", {
        method: "POST",
        body: formData,
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      const data = await res.json();
      if (res.ok) {
        setUploadSuccess("圖片上傳成功！");
        setUploadedImageUrl(data.imageUrl);
        console.log("圖片連結：", data.imageUrl);
        // 重新抓取使用者資料
        setProfile(prev => ({ ...prev, image_url: data.imageUrl }));
      } else {
        setUploadError(data.error || "上傳失敗");
      }
    } catch (err) {
      setUploadError("連線錯誤：" + err.message);
    } finally {
      setUploading(false);
    }
  };

  return (
    <Layout>
      <div className="container mt-5">
        <h2 className="mb-4">個人資料</h2>

        {loading && <p>載入中...</p>}
        {error && !loading && <p className="text-danger">{error}</p>}

        {profile && (
          <div className="card p-4 shadow">
            <p>
              <strong>姓名：</strong> {profile.first_name} {profile.last_name}
            </p>
            <p>
              <strong>信箱：</strong> {profile.email}
            </p>
            <p>
              <strong>使用者類型：</strong> {profile.user_type}
            </p>

            {profile.image_url && (
              <div className="mt-3">
                <p>目前大頭貼：</p>
                <img
                  src={profile.image_url}
                  alt="目前的大頭貼"
                  style={{
                    width: "150px",
                    height: "150px",
                    objectFit: "cover",
                    borderRadius: "50%",
                    border: "1px solid #ccc",
                  }}
                />
              </div>
            )}

            <hr />
            <div className="mt-3">
              <label className="form-label">上傳個人照片：</label>
              <input
                type="file"
                accept="image/*"
                onChange={handleImageUpload}
                className="form-control"
              />
            </div>

            {uploading && <p>上傳中...</p>}
            {uploadSuccess && (
              <div className="alert alert-success mt-2">{uploadSuccess}</div>
            )}
            {uploadError && (
              <div className="alert alert-danger mt-2">{uploadError}</div>
            )}

            {uploadedImageUrl && (
              <div className="mt-3">
                <p>預覽圖片：</p>
                <img
                  src={uploadedImageUrl}
                  alt="上傳的大頭貼"
                  style={{
                    width: "150px",
                    height: "150px",
                    objectFit: "cover",
                    borderRadius: "50%",
                    border: "1px solid #ccc",
                  }}
                />
              </div>
            )}
          </div>
        )}
      </div>
    </Layout>
  );
}

export default ProfilePage;


