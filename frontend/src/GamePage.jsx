import { useState } from "react";
import WinModal from "./WinModal";
import Leaderboard from "./Leaderboard";
function GamePage() {
  const [guess, setGuess] = useState("");
  const [message, setMessage] = useState("");
  const [attempts, setAttempts] = useState(0);
  const [guessHistory, setGuessHistory] = useState([]);
  const [winData, setWinData] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (guess.length !== 4 || isNaN(guess)) {
      setMessage("請輸入有效的 4 位數字");
      return;
    }

    const token = localStorage.getItem("token");

    try {
      const res = await fetch(`http://localhost:3000/api/users/game/guess`, {
        method: "POST",
        headers: {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ guess }),
      });

      const data = await res.json();

      if (res.ok) {
        setMessage(data.result || "成功");
        setGuessHistory((prev) => [...prev, data.guess]);

        if (data.result.includes("Congratulations")) {
          setWinData(data); // 顯示 Modal
        }

        setAttempts((prev) => prev + 1);
      } else {
        setMessage(data.error || "錯誤");
      }
    } catch (err) {
      setMessage(err.message || "連線失敗");
    }

    setGuess("");
  };

  return (
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        minHeight: "100vh",
        backgroundColor: "#f0f4f8",
      }}
    >
      <div
        style={{
          backgroundColor: "white",
          padding: "30px",
          borderRadius: "8px",
          boxShadow: "0 4px 10px rgba(0, 0, 0, 0.1)",
          textAlign: "center",
          width: "300px",
        }}
      >
        <h1 style={{ fontSize: "24px", marginBottom: "20px" }}>1A2B 遊戲</h1>
        <p>請輸入一個 4 位數字來猜：</p>

        <form onSubmit={handleSubmit}>
          <input
            type="text"
            value={guess}
            onChange={(e) => setGuess(e.target.value)}
            maxLength={4}
            placeholder="輸入 4 位數字"
            required
            style={{
              width: "100%",
              padding: "10px",
              marginBottom: "10px",
              borderRadius: "4px",
              border: "1px solid #ddd",
            }}
          />
          <button
            type="submit"
            style={{
              width: "100%",
              padding: "10px",
              backgroundColor: "#4CAF50",
              color: "white",
              border: "none",
              borderRadius: "4px",
              cursor: "pointer",
              fontSize: "16px",
            }}
          >
            提交猜測
          </button>
        </form>

        <div
          style={{ marginTop: "15px", fontSize: "18px", fontWeight: "bold" }}
        >
          {message && <p>{message}</p>}
        </div>

        <div style={{ marginTop: "10px", fontSize: "14px", color: "#666" }}>
          已猜測次數：{attempts}
        </div>

        <div style={{ marginTop: "20px", textAlign: "left" }}>
          <p
            style={{
              fontWeight: "bold",
              fontSize: "16px",
              marginBottom: "8px",
            }}
          >
            猜過的答案：
          </p>
          <div
            style={{
              maxHeight: "150px",
              overflowY: "auto",
              border: "1px solid #ddd",
              borderRadius: "4px",
              padding: "10px",
              backgroundColor: "#f9f9f9",
            }}
          >
            {guessHistory.length === 0 ? (
              <p style={{ color: "#888" }}>尚未猜測</p>
            ) : (
              guessHistory.map((item, index) => (
                <div
                  key={index}
                  style={{
                    padding: "6px 10px",
                    borderBottom: "1px solid #eee",
                    color: "#333",
                  }}
                >
                  第 {index + 1} 次：{item}
                </div>
              ))
            )}
          </div>
          <Leaderboard/>
        </div>

      </div>

      {/* Win Modal */}
      <WinModal
        open={!!winData}
        onClose={() => setWinData(null)}
        data={winData}
      />
    </div>
  );
}

export default GamePage;
