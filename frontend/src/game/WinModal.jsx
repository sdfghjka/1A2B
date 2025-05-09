// components/WinModal.jsx
import React from "react";
import { useNavigate } from "react-router-dom";

function WinModal({ open, onClose, data }) {
  const navigate = useNavigate();

  if (!open) return null;

  const handleClose = () => {
    onClose(); // 如果你還想清除 modal 狀態的話
    navigate("/gamestart"); // 導向 StartPage
  };

  return (
    <div
      style={{
        position: "fixed",
        top: 0,
        left: 0,
        width: "100vw",
        height: "100vh",
        backgroundColor: "rgba(0,0,0,0.5)",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        zIndex: 999,
      }}
    >
      <div
        style={{
          backgroundColor: "#fff",
          padding: "30px",
          borderRadius: "8px",
          maxWidth: "400px",
          textAlign: "center",
          boxShadow: "0 4px 10px rgba(0,0,0,0.2)",
        }}
      >
        <h2 style={{ marginBottom: "20px" }}>🎉 恭喜你猜對了！</h2>
        <p><strong>猜的次數：</strong> {data.count}</p>
        <p><strong>總花費時間：</strong> {data.duration}</p>
        <p><strong>起始時間：</strong> {new Date(data.startTime).toLocaleString()}</p>
        <p><strong>最後一次猜的數字：</strong> {data.guess}</p>
        <button
          onClick={handleClose}
          style={{
            marginTop: "20px",
            padding: "10px 20px",
            backgroundColor: "#4CAF50",
            color: "#fff",
            border: "none",
            borderRadius: "4px",
            cursor: "pointer",
          }}
        >
          關閉
        </button>
      </div>
    </div>
  );
}

export default WinModal;
