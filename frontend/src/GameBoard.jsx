import React, { useState, useEffect } from "react";

function GameBoard({ socket, roomId }) {
  const [guess, setGuess] = useState("");
  const [chatInput, setChatInput] = useState("");
  const [messages, setMessages] = useState([]);

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

  useEffect(() => {
    if (!socket) return;

    socket.onmessage = (event) => {
      const msg = JSON.parse(event.data);
      console.log("📩 收到訊息:", msg);

      switch (msg.type) {
        case "guessResult":
          setMessages((prev) => [...prev, msg]);
          break;
        case "chat":
          setMessages((prev) => [...prev, { chat: msg.payload }]);
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

      <form onSubmit={handleChatSubmit} style={{ marginTop: "1rem" }}>
        <input
          type="text"
          value={chatInput}
          onChange={(e) => setChatInput(e.target.value)}
          placeholder="💬 傳送訊息"
        />
        <button type="submit">送出</button>
      </form>

      <div style={{ marginTop: "1rem" }}>
        {messages.map((msg, index) => {
          if (msg.system) return <p key={index} style={{ color: "gray" }}>{msg.system}</p>;
          if (msg.chat) return <p key={index} style={{ color: "blue" }}>💬 {msg.chat}</p>;
          return (
            <p key={index}>
              🧠 {msg.guess} ➜ 🎯 {msg.result}
            </p>
          );
        })}
      </div>
    </div>
  );
}

export default GameBoard;


