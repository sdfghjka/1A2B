import { useState, useEffect } from "react";
import { useNavigate, useLocation, Link } from "react-router-dom";
import "bootstrap/dist/css/bootstrap.min.css";
import { FaFacebook, FaGoogle } from "react-icons/fa";

function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const [successMessage, setSuccessMessage] = useState("");
  const navigate = useNavigate();
  const location = useLocation();

  // 處理註冊成功訊息
  useEffect(() => {
    if (location.state?.successMessage) {
      setSuccessMessage(location.state.successMessage);
      // 清除 state，避免重整還顯示
      window.history.replaceState({}, document.title);
    }
  }, [location.state]);

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
        localStorage.setItem("token", data.token);
        localStorage.setItem("user", JSON.stringify(data));
        navigate("/gamestart");
      } else {
        setErrorMessage(data.error || "登入失敗，請檢查帳號和密碼");
      }
    } catch (error) {
      setErrorMessage(error.message || "發生錯誤，請稍後再試");
    }
  };

  const handleGoogleLogin = () => {
    window.location.href = "http://localhost:3000/api/auth/google";
  };

  const handleFacebookLogin = () => {
    setErrorMessage("Facebook 登入功能尚未實現");
  };

  return (
    <div className="container d-flex justify-content-center align-items-center min-vh-100">
      <div className="card p-4 shadow text-center" style={{ width: "400px" }}>
        <h2 className="mb-3">登入</h2>

        {/* 成功註冊提示 */}
        {successMessage && (
          <div className="alert alert-success">{successMessage}</div>
        )}

        {/* 錯誤訊息 */}
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
          <p>還沒有帳號？<Link to="/signup">註冊</Link></p>
          <p>或使用其他方式登入</p>
          <div className="d-grid gap-2">
            <button
              className="btn btn-outline-primary d-flex align-items-center justify-content-center"
              onClick={handleFacebookLogin}
            >
              <FaFacebook size={20} className="me-2" /> 使用 Facebook 登入
            </button>
            <button
              className="btn btn-outline-danger d-flex align-items-center justify-content-center"
              onClick={handleGoogleLogin}
            >
              <FaGoogle size={20} className="me-2" /> 使用 Google 登入
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Login;
