import React, { useEffect, useState } from "react";
import { ClipLoader } from "react-spinners";
import GameBoard from "./GameBoard";
import GameAIVersion from "./GameAIVersion";
import { useLocation } from "react-router-dom";

function GameRoomPage() {
  const [socket, setSocket] = useState(null);
  const [roomJoined, setRoomJoined] = useState(false);
  const [roomId, setRoomId] = useState("");
  const location = useLocation();

  // 讀取 mode
  const searchParams = new URLSearchParams(location.search);
  const isAI = searchParams.get("mode") === "ai";

  useEffect(() => {
    const token = localStorage.getItem("token");
    const endpoint = isAI ? "/api/ai/start" : "/api/ws";
    const ws = new WebSocket(`ws://localhost:3000${endpoint}?token=${token}`);
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
  }, [isAI]);

  return (
    <div style={{ textAlign: "center", padding: "20px" }}>
      {roomJoined && socket ? (
        isAI ? (
          <GameAIVersion socket={socket} roomId={roomId} />
        ) : (
          <GameBoard socket={socket} roomId={roomId} />
        )
      ) : (
        <div>
          <p>{isAI ? "與 AI 建立連線中..." : "等待配對中..."}</p>
          <ClipLoader size={50} color={"#36d7b7"} loading={!roomJoined} />
        </div>
      )}
    </div>
  );
}

export default GameRoomPage;


