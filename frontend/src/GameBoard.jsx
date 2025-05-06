import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";

function GameBoard({ socket, roomId }) {
  const [guess, setGuess] = useState("");
  const [chatInput, setChatInput] = useState("");
  const [messages, setMessages] = useState([]);
  const navigate = useNavigate();

  const handleSubmit = (e) => {
    e.preventDefault();
    if (guess.length !== 4 || isNaN(guess)) return;
    socket.send(JSON.stringify({ type: "guess", payload: guess }));
    setGuess("");
  };

  const handleChatSubmit = (e) => {
    e.preventDefault();
    if (chatInput.trim()) {
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

      setMessages((prev) => [...prev, msg]);
    };

    return () => {
      socket.onmessage = null;
    };
  }, [socket, navigate]);

  const renderMessage = (msg, index) => {
    switch (msg.type) {
      case "guessResult":
        return (
          <p key={index}>
            🧠 <strong>{msg.from}</strong> guessed ➜{" "}
            <span style={{ color: "green" }}>{msg.payload}</span>
          </p>
        );
      case "chat":
        return (
          <p key={index}>
            💬 <strong>{msg.from}</strong>: {msg.payload}
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
        return (
          <p key={index} style={{ color: "orange" }}>
            ⚠ Unknown message type: {msg.type}
          </p>
        );
    }
  };

  return (
    <div style={{ padding: "1rem", maxWidth: 600, margin: "auto" }}>
      <h3>Room ID: {roomId || "Waiting to join..."}</h3>

      <form onSubmit={handleSubmit} style={{ marginBottom: "1rem" }}>
        <input
          type="text"
          value={guess}
          onChange={(e) => setGuess(e.target.value)}
          maxLength={4}
          placeholder="Enter 4-digit guess"
          required
        />
        <button type="submit">Submit Guess</button>
      </form>

      <form onSubmit={handleChatSubmit}>
        <input
          type="text"
          value={chatInput}
          onChange={(e) => setChatInput(e.target.value)}
          placeholder="Send a chat message"
          required
        />
        <button type="submit">Send</button>
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
        }}
      >
        {messages.map((msg, index) => renderMessage(msg, index))}
      </div>
    </div>
  );
}

export default GameBoard;

