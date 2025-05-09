import { useEffect, useRef } from "react";
import { io } from "socket.io-client";

export default function useSocket(roomId, onMessage) {
  const socketRef = useRef(null);

  useEffect(() => {
    const socket = io("http://localhost:3000", {
      query: { roomId },
    });
    socketRef.current = socket;

    socket.on("message", onMessage);

    return () => {
      socket.disconnect();
    };
  }, [roomId, onMessage]);

  return socketRef;
}
