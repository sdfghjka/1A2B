import React, { useEffect, useState } from "react";
import GameBoard from "./GameBoard";

function GameRoomPage() {
  const [roomId, setRoomId] = useState(null);
  const [socket, setSocket] = useState(null);

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      console.error("Token 不存在，請先登入！");
      return;
    }

    const ws = new WebSocket(`ws://localhost:3000/api/ws?roomId=default&token=${token}`);

    ws.onopen = () => {
      console.log("✅ WebSocket connected");
    };

    ws.onmessage = (event) => {
      const msg = JSON.parse(event.data);
      if (msg.type === "roomJoined") {
        setRoomId(msg.data.roomId);
      }
    };

    ws.onerror = (err) => {
      console.error("❌ WebSocket error:", err);
    };

    ws.onclose = () => {
      console.warn("❎ WebSocket disconnected");
    };

    setSocket(ws);

    return () => {
      ws.close();
    };
  }, []);

  return (
    <div>
      <h2>多人對戰房間</h2>
      {roomId && socket ? (
        <GameBoard socket={socket} />
      ) : (
        <p>等待加入房間...</p>
      )}
    </div>
  );
}

export default GameRoomPage;


