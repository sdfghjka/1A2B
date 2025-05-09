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
        const res = await fetch("http://localhost:3000/api/users/admin?recordPerPage=10&page=1", {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
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
                <td>{user.first_name} {user.last_name}</td>
                <td>{user.email}</td>
                <td>{user.user_type}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </Layout>
  );
}

export default AdminUserList;
