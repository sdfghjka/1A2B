import { createContext, useContext, useState, useEffect } from "react";

const UserContext = createContext();

export function UserProvider({ children }) {
  const [user, setUser] = useState(null);

  // 自動載入 user
  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) return;

    fetch("http://localhost:3000/api/users/me", {
      headers: { Authorization: `Bearer ${token}` },
    })
      .then(res => res.json())
      .then(data => {
        if (data.email) {
          setUser(data); // 確保設定正確的 user 資料
        } else {
          console.log("用戶資料無效", data);
        }
      })
      .catch(err => console.error("驗證失敗:", err));
  }, []); // 這裡只會在首次載入時執行

  const logout = () => {
    localStorage.removeItem("token");
    setUser(null);
  };

  return (
    <UserContext.Provider value={{ user, setUser, logout }}>
      {children}
    </UserContext.Provider>
  );
}

export function useUser() {
  return useContext(UserContext);
}

