import React, { useState, useEffect } from "react";

function GameBoard({ socket }) {
  const [guess, setGuess] = useState("");
  const [messages, setMessages] = useState([]);
  const [roomId, setRoomId] = useState("");

  const handleSubmit = (e) => {
    e.preventDefault();
    if (guess.length !== 4 || isNaN(guess)) return;

    if (socket.readyState === WebSocket.OPEN) {
      socket.send(JSON.stringify({ type: "guess", payload: guess }));
      setGuess("");
    }
  };

  useEffect(() => {
    if (!socket) return;

    socket.onmessage = (event) => {
      const msg = JSON.parse(event.data);
      console.log("📩 收到訊息:", msg);

      switch (msg.type) {
        case "roomJoined":
          setRoomId(msg.data.roomId);
          setMessages((prev) => [...prev, { system: `✅ 已加入房間 ${msg.data.roomId}` }]);
          break;
        case "guessResult":
          setMessages((prev) => [...prev, msg]);
          break;
        case "gameOver":
          alert(msg.data);
          break;
        default:
          setMessages((prev) => [...prev, { system: `📎 收到未知訊息: ${msg.type}` }]);
      }
    };

    return () => {
      socket.onmessage = null;
    };
  }, [socket]);

  return (
    <div>
      <h3>房間：{roomId || "未加入"}</h3>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          value={guess}
          onChange={(e) => setGuess(e.target.value)}
          maxLength={4}
          placeholder="輸入四位數"
        />
        <button type="submit">提交</button>
      </form>
      <div>
        {messages.map((msg, index) =>
          msg.system ? (
            <p key={index} style={{ color: "gray" }}>{msg.system}</p>
          ) : (
            <p key={index}>
              🧠 {msg.guess} ➜ 🎯 {msg.result}
            </p>
          )
        )}
      </div>
    </div>
  );
}

export default GameBoard;

