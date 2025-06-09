import { useEffect, useState } from "react";
import Layout from "../components/Layout";

function AdminUserList() {
  const [users, setUsers] = useState([]);
  const [total, setTotal] = useState(0);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchUsers = async () => {
      setLoading(true);
      const token = localStorage.getItem("token");
      try {
        const res = await fetch(
          "https://dc03-2401-e180-8820-8783-ed6c-8b9f-c725-1ce2.ngrok-free.app/api/users/admin?recordPerPage=10&page=1",
          {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        );
        const data = await res.json();
        if (res.ok) {
          setUsers(data.user_items || []);
          setTotal(data.total_count || 0);
        } else {
          setError(data.error || "取得資料失敗");
        }
      } catch (err) {
        setError(err);
      } finally {
        setLoading(false);
      }
    };

    fetchUsers();
  }, []);

  const handleUpgrade = async (email) => {
    try {
      const res = await fetch("https://e312-2401-e180-8820-8783-ed6c-8b9f-c725-1ce2.ngrok-free.app/api/payment", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email }),
      });
      const html = await res.text();

      const win = window.open("", "_blank");
      if (win) {
        win.document.open();
        win.document.write(html);
        win.document.close();
      } else {
        alert("無法開啟付款視窗，請檢查瀏覽器設定");
      }
    } catch (err) {
      alert(err || "升級失敗，請稍後再試");
    }
  };

  return (
    <Layout>
      <div className="container mt-5">
        <h2 className="mb-4">使用者管理 ({total} 人)</h2>
        {loading && <p>載入中...</p>}
        {error && <p className="text-danger">{error}</p>}
        <table className="table table-bordered">
          <thead>
            <tr>
              <th>姓名</th>
              <th>Email</th>
              <th>身分</th>
            </tr>
          </thead>
          <tbody>
            {users.map((user, idx) => (
              <tr key={idx}>
                <td>
                  {user.first_name} {user.last_name}
                </td>
                <td>{user.email}</td>
                <td>
                  {user.user_type}
                  {user.user_type !== "vip" && (
                    <button
                      className="btn btn-sm btn-success ms-2"
                      onClick={() => handleUpgrade(user.email)}
                    >
                      升級會員 $99
                    </button>
                  )}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </Layout>
  );
}

export default AdminUserList;
