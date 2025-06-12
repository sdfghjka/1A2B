import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { ClipLoader } from "react-spinners";

function GameAIVersion({ socket }) {
  const [guess, setGuess] = useState("");
  const [systemMessages, setSystemMessages] = useState([]);
  const [loading, setLoading] = useState(false);
  const [aiGuess, setAiGuess] = useState("");
  const navigate = useNavigate();

  const handleSubmit = (e) => {
    e.preventDefault();
    if (guess.length !== 4 || isNaN(guess)) return;
    setLoading(true);
    socket.send(JSON.stringify({ type: "guess", payload: guess }));
    setGuess("");
  };

  const handleLeave = () => {
    if (socket) {
      socket.send(JSON.stringify({ type: "leave" }));
      socket.close();
    }
    navigate("/gamestart");
  };

  useEffect(() => {
    if (!socket) return;

    socket.onmessage = (event) => {
      const msg = JSON.parse(event.data);
      console.log("📩 Received:", msg);

      switch (msg.type) {
        case "aiGuess":
          setAiGuess(msg.payload.guess);
          break;
        case "playerResult":
        case "aiResult":
        case "gameOver":
        case "system":
        case "guessResult": 
          setSystemMessages((prev) => [...prev, msg]);
          break;
        default:
          break;
      }

      setLoading(false);
    };

    return () => {
      socket.onmessage = null;
    };
  }, [socket, navigate]);

  const renderSystemMessage = (msg, index) => {
    switch (msg.type) {
      case "guessResult":
        return (
          <p key={index}>
            🧠 <strong>{msg.from}</strong> guessed{" "}
            <span style={{ color: "green" }}>{msg.payload}</span>
          </p>
        );
      case "roomJoined":
        return (
          <p key={index} style={{ color: "purple" }}>
            🔗 <strong>{msg.from}</strong> joined room{" "}
            <strong>{msg.payload.roomId}</strong>
          </p>
        );
      case "gameOver":
        return (
          <p key={index} style={{ color: "red" }}>
            🎉 Game Over: {msg.payload}
          </p>
        );
      case "system":
        return (
          <p key={index} style={{ color: "gray" }}>
            🔔 {msg.payload}
          </p>
        );
      default:
        return null;
    }
  };

  return (
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "flex-start",
        padding: "50px",
        backgroundColor: "#f0f4f8",
        gap: "30px",
      }}
    >
      <div
        style={{
          backgroundColor: "white",
          padding: "30px",
          borderRadius: "8px",
          boxShadow: "0 4px 10px rgba(0, 0, 0, 0.1)",
          textAlign: "center",
          width: "400px",
        }}
      >
        <h1 style={{ fontSize: "24px", marginBottom: "10px" }}>
          1A2B 遊戲（對戰電腦）
        </h1>

        {aiGuess && (
          <div style={{ marginBottom: "10px", color: "#666" }}>
            🤖 最新電腦猜測：<strong>{aiGuess}</strong>
          </div>
        )}

        <form onSubmit={handleSubmit} style={{ marginBottom: "1rem" }}>
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

        <button
          onClick={handleLeave}
          style={{
            backgroundColor: "#dc3545",
            color: "white",
            padding: "10px",
            border: "none",
            borderRadius: "4px",
            width: "100%",
          }}
        >
          離開遊戲
        </button>

        <div
          style={{
            marginTop: "1.5rem",
            background: "#f8f8f8",
            padding: "1rem",
            borderRadius: "8px",
            textAlign: "left",
            maxHeight: "300px",
            overflowY: "auto",
          }}
        >
          {loading ? (
            <div style={{ textAlign: "center" }}>
              <ClipLoader size={30} color={"#4CAF50"} loading={loading} />
              <p>處理中...</p>
            </div>
          ) : (
            systemMessages.map((msg, index) => renderSystemMessage(msg, index))
          )}
        </div>
      </div>
    </div>
  );
}

export default GameAIVersion;
