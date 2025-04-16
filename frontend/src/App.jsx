import { useState } from "react";
import { useLocation, useNavigate } from "react-router-dom"; // 引入 useLocation 和 useNavigate
import "bootstrap/dist/css/bootstrap.min.css";
import { FaFacebook, FaGoogle } from "react-icons/fa";
import { Link } from "react-router-dom";

function App() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const location = useLocation(); // 使用 useLocation 來讀取傳遞過來的狀態
  const navigate = useNavigate(); // 用來進行頁面導航

  const handleSubmit = async (e) => {
    e.preventDefault();

    const userData = {
      Email: email,
      Password: password,
    };

    try {
      const response = await fetch("http://localhost:3000/api/user/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(userData),
      });

      const data = await response.json();

      if (response.ok) {
        // 登入成功，將 token 儲存到 localStorage 或 sessionStorage
        localStorage.setItem("token", data.token); // 儲存 token
        navigate("/gamestart"); // 假設登入成功後跳轉到 dashboard 頁面
      } else {
        // 登入失敗，顯示錯誤訊息
        setErrorMessage(data.error || "登入失敗，請檢查帳號和密碼");
      }
    } catch (error) {
      setErrorMessage(error.message || "發生錯誤，請稍後再試");
    }
  };

  return (
    <div className="container d-flex justify-content-center align-items-center min-vh-100">
      <div className="card p-4 shadow text-center" style={{ width: "400px" }}>
        <h2 className="mb-3">登入</h2>

        {/* 顯示註冊成功的訊息 */}
        {location.state?.successMessage && (
          <div className="alert alert-success" role="alert">
            {location.state.successMessage}
          </div>
        )}

        {/* 顯示錯誤訊息 */}
        {errorMessage && (
          <div className="alert alert-danger" role="alert">
            {errorMessage}
          </div>
        )}

        <form onSubmit={handleSubmit}>
          <div className="mb-3 text-start">
            <label className="form-label">Email</label>
            <input
              type="email"
              className="form-control"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />
          </div>
          <div className="mb-3 text-start">
            <label className="form-label">密碼</label>
            <input
              type="password"
              className="form-control"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>
          <button type="submit" className="btn btn-primary w-100">
            登入
          </button>
        </form>

        <div className="mt-3">
          <p>或使用其他方式登入</p>
          <div className="d-grid gap-2">
            <button className="btn btn-outline-primary d-flex align-items-center justify-content-center">
              <FaFacebook size={20} className="me-2" /> 使用 Facebook 登入
            </button>
            <button className="btn btn-outline-danger d-flex align-items-center justify-content-center">
              <FaGoogle size={20} className="me-2" /> 使用 Google 登入
            </button>
          </div>
        </div>

        {/* 註冊頁面按鈕 */}
        <div className="mt-3">
          <Link to="/signup" className="btn btn-link">
            還沒有帳號? 註冊
          </Link>
        </div>
      </div>
    </div>
  );
}

export default App;
