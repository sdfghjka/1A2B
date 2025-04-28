import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

function AuthCallback() {
  const navigate = useNavigate();

  useEffect(() => {
    const query = new URLSearchParams(window.location.search);
    const token = query.get("token");

    if (token) {
      localStorage.setItem("token", token);
      navigate("/gamestart");
    } else {
      alert(token);
    }
  }, [navigate]);

  return <p>登入中，請稍候...</p>;
}

export default AuthCallback;
