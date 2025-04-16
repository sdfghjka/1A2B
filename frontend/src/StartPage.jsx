// StartPage.js
import { useNavigate } from "react-router-dom";

function StartPage() {
  const navigate = useNavigate();

  const startGame = async () => {
    const token = localStorage.getItem("token"); // 取得 token
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
  return (
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        minHeight: "100vh",
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
        開始遊戲
      </button>
    </div>
  );
}

export default StartPage;
