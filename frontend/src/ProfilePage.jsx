import { useEffect, useState } from "react";
import { useUser } from "./contexts/UserContext";
import Layout from "./components/Layout";

function ProfilePage() {
  const { user } = useUser(); // 從 context 拿 user 資訊
  const [profile, setProfile] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    // 確保 user 已經存在並且有 id
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
        console.log("Response data:", data);  // 檢查 API 回應

        if (res.ok) {
          setProfile(data); // 更新資料
        } else {
          setError(data.error || "取得資料失敗");
        }
      } catch (err) {
        setError(err.message || "連線錯誤");
      } finally {
        setLoading(false);
      }
    };

    fetchProfile(); // 發送請求
  }, [user]); // 依賴 user，當 user 改變時會重新執行

  return (
    <Layout>
      <div className="container mt-5">
        <h2 className="mb-4">個人資料</h2>
        {loading && <p>載入中...</p>}
        {error && !loading && <p className="text-danger">{error}</p>} 
        {profile && (
          <div className="card p-4 shadow">
            <p><strong>姓名：</strong> {profile.first_name} {profile.last_name}</p>
            <p><strong>信箱：</strong> {profile.email}</p>
            <p><strong>使用者類型：</strong> {profile.user_type}</p>
            {/* 如有其他欄位可補充 */}
          </div>
        )}
      </div>
    </Layout>
  );
}

export default ProfilePage;

