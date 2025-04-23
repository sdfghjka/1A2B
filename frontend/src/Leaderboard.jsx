import { useEffect, useState } from "react";

function Leaderboard() {
  const [leaders, setLeaders] = useState([]);

  useEffect(() => {
    const token = localStorage.getItem("token");
    const fetchLeaderboard = async () => {
      try {
        const res = await fetch("http://localhost:3000/api/users/leaderboard", {
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        const data = await res.json();

        if (res.ok) {
          setLeaders(Array.isArray(data) ? data : []); // 保證 data 是陣列
        } else {
          console.error("無法取得排行榜", data.error);
        }
      } catch (err) {
        console.error("發生錯誤", err);
      }
    };

    fetchLeaderboard();
  }, []);

  return (
    <div
      style={{
        width: "250px",
        backgroundColor: "#fff",
        marginLeft: "30px",
        padding: "20px",
        borderRadius: "8px",
        boxShadow: "0 4px 10px rgba(0, 0, 0, 0.1)",
        maxHeight: "80vh",
        overflowY: "auto",
      }}
    >
      <h3 style={{ textAlign: "center", marginBottom: "15px" }}>排行榜</h3>
      {leaders.length === 0 ? (
        <p style={{ textAlign: "center" }}>暫無資料</p>
      ) : (
        leaders.map((user, index) => (
          <div
            key={user.id}
            style={{
              marginBottom: "10px",
              fontWeight: "bold",
              display: "flex",
              justifyContent: "space-between",
            }}
          >
            <span>{index + 1}. 玩家 {user.id.slice(0, 6)}</span>
            <span>{user.count} 次 / {user.duration.toFixed(1)} 秒</span>
          </div>
        ))
      )}
    </div>
  );
}

export default Leaderboard;
