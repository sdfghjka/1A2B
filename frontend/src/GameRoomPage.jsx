import React, { useEffect, useState } from "react";
import { ClipLoader } from "react-spinners"; // 載入 Spinner 動畫
import GameBoard from "./GameBoard";

function GameRoomPage() {
  const [socket, setSocket] = useState(null);
  const [roomJoined, setRoomJoined] = useState(false);
  const [roomId, setRoomId] = useState("");

  useEffect(() => {
    const token = localStorage.getItem("token");
    const ws = new WebSocket(`ws://localhost:3000/api/ws?token=${token}`);
    setSocket(ws);

    ws.onmessage = (event) => {
      const msg = JSON.parse(event.data);
      if (msg.type === "roomJoined") {
        setRoomJoined(true);
        if (msg.payload?.roomId) {
          setRoomId(msg.payload.roomId);
        }
      }
    };

    ws.onerror = (e) => console.error("WebSocket error:", e);
    ws.onclose = () => console.log("WebSocket closed");

    return () => ws.close();
  }, []);

  return (
    <div style={{ textAlign: "center", padding: "20px" }}>
      {roomJoined && socket ? (
        <GameBoard socket={socket} roomId={roomId} />
      ) : (
        <div>
          <p>等待配對中...</p>
          <ClipLoader size={50} color={"#36d7b7"} loading={!roomJoined} />
        </div>
      )}
    </div>
  );
}

export default GameRoomPage;


