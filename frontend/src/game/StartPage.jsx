import { useNavigate } from "react-router-dom";
import Layout from "../components/Layout";

function StartPage() {
  const navigate = useNavigate();

  const startGame = async () => {
    const token = localStorage.getItem("token");
    try {
      const res = await fetch("http://localhost:3000/api/users/game/random", {
        method: "GET",
        headers: {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        },
      });

      const data = await res.json();

      if (res.ok) {
        navigate(`/game?gameId=${data.gameId}`);
      } else {
        alert(data.error || "發生錯誤，請稍後再試");
      }
    } catch (err) {
      alert(err.message || "連線失敗");
    }
  };

  const startVsAI = async () => {
    const token = localStorage.getItem("token");
    try {
      const res = await fetch("http://localhost:3000/api/ai/start", {
        method: "POST",
        headers: {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        },
      });

      const data = await res.json();

      if (res.ok) {
        navigate(`/game/ai`);
      } else {
        alert(data.error || "發生錯誤，請稍後再試");
      }
    } catch (err) {
      alert(err.message || "連線失敗");
    }
  };

  const goToMultiplayer = () => {
    navigate("/multiplayer");
  };

  const user = {
    email: "test@example.com",
    id: "abc123",
    is_admin: false,
  };

  return (
    <Layout user={user}>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          justifyContent: "center",
          alignItems: "center",
          minHeight: "100vh",
          gap: "20px",
        }}
      >
        <button
          onClick={startGame}
          style={{
            fontSize: "20px",
            padding: "15px 30px",
            backgroundColor: "#007bff",
            color: "white",
            border: "none",
            borderRadius: "5px",
          }}
        >
          單人模式
        </button>

        <button
          onClick={startVsAI}
          style={{
            fontSize: "20px",
            padding: "15px 30px",
            backgroundColor: "#666666",
            color: "white",
            border: "none",
            borderRadius: "5px",
          }}
        >
          電腦對戰
        </button>
        <button
          onClick={goToMultiplayer}
          style={{
            fontSize: "20px",
            padding: "15px 30px",
            backgroundColor: "#28a745",
            color: "white",
            border: "none",
            borderRadius: "5px",
          }}
        >
          即時對戰
        </button>
      </div>
    </Layout>
  );
}

export default StartPage;
