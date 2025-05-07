import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { ClipLoader } from "react-spinners"; // 引入ClipLoader

function GameBoard({ socket, roomId }) {
  const [guess, setGuess] = useState("");
  const [chatInput, setChatInput] = useState("");
  const [chatMessages, setChatMessages] = useState([]);
  const [systemMessages, setSystemMessages] = useState([]);
  const [loading, setLoading] = useState(false); // 用來控制加載狀態
  const navigate = useNavigate();

  const handleSubmit = (e) => {
    e.preventDefault();
    if (guess.length !== 4 || isNaN(guess)) return;
    setLoading(true);  // 設定為加載中
    socket.send(JSON.stringify({ type: "guess", payload: guess }));
    setGuess("");
  };

  const handleChatSubmit = (e) => {
    e.preventDefault();
    if (chatInput.trim()) {
      setLoading(true); // 設定為加載中
      socket.send(JSON.stringify({ type: "chat", payload: chatInput }));
      setChatInput("");
    }
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

      if (msg.type === "playerLeft") {
        alert("對手已離開，返回首頁");
        navigate("/gamestart");
        return;
      }

      if (msg.type === "chat") {
        setChatMessages((prev) => [...prev, msg]);
      } else {
        setSystemMessages((prev) => [...prev, msg]);
      }

      setLoading(false); // 收到回應後停止加載
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
            🧠 <strong>{msg.from}</strong> guessed {" "}
            <span style={{ color: "green" }}>{msg.payload}</span>
          </p>
        );
      case "roomJoined":
        return (
          <p key={index} style={{ color: "purple" }}>
            🔗 <strong>{msg.from}</strong> joined room <strong>{msg.payload.roomId}</strong>
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

  const renderChatMessage = (msg, index) => (
    <p key={index}>
      💬 <strong>{msg.from}</strong>: {msg.payload}
    </p>
  );

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
      {/* 左邊：遊戲主體與系統訊息 */}
      <div
        style={{
          backgroundColor: "white",
          padding: "30px",
          borderRadius: "8px",
          boxShadow: "0 4px 10px rgba(0, 0, 0, 0.1)",
          textAlign: "center",
          width: "320px",
        }}
      >
        <h1 style={{ fontSize: "24px", marginBottom: "20px" }}>1A2B 遊戲</h1>
        <p>房間代碼：{roomId || "等待加入..."}</p>

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
            marginTop: "1rem",
            backgroundColor: "#dc3545",
            color: "white",
            padding: "10px 20px",
            border: "none",
            borderRadius: "5px",
            width: "100%",
          }}
        >
          離開房間
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
              <p>正在處理中...</p>
            </div>
          ) : (
            systemMessages.map((msg, index) => renderSystemMessage(msg, index))
          )}
        </div>
      </div>

      {/* 右邊：聊天室 */}
      <div
        style={{
          backgroundColor: "white",
          padding: "30px",
          borderRadius: "8px",
          boxShadow: "0 4px 10px rgba(0, 0, 0, 0.1)",
          width: "400px",
        }}
      >
        <h2 style={{ marginBottom: "20px" }}>聊天室</h2>
        <form onSubmit={handleChatSubmit} style={{ marginBottom: "1rem" }}>
          <input
            type="text"
            value={chatInput}
            onChange={(e) => setChatInput(e.target.value)}
            placeholder="輸入訊息..."
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
              backgroundColor: "#007BFF",
              color: "white",
              border: "none",
              borderRadius: "4px",
              cursor: "pointer",
              fontSize: "16px",
            }}
          >
            發送
          </button>
        </form>

        <div
          style={{
            background: "#f8f8f8",
            padding: "1rem",
            borderRadius: "8px",
            maxHeight: "400px",
            overflowY: "auto",
          }}
        >
          {chatMessages.map((msg, index) => renderChatMessage(msg, index))}
        </div>
      </div>
    </div>
  );
}

export default GameBoard;

