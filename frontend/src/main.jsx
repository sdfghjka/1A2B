import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './index.css';
import App from './App.jsx';
import Signup from './Signup.jsx';
import GamePage from './game/GamePage.jsx';
import StartPage from './game/StartPage.jsx';
import AuthCallback from './AuthCallback.jsx';
import GameRoomPage from './game/GameRoomPage.jsx';
import 'bootstrap/dist/css/bootstrap.min.css';
import ProfilePage from './ProfilePage.jsx';
import { UserProvider } from "./contexts/UserContext";
import AdminUserList from './users/AdminUserList.jsx';
// 正確渲染並處理路由
createRoot(document.getElementById('root')).render(
  <StrictMode>
    <Router>
      <UserProvider>
        <Routes>
          {/* 定義各路由 */}
          <Route path="/login" element={<App />} />
          <Route path="/signup" element={<Signup />} />
          <Route path="/gamestart" element={<StartPage />} />
          <Route path="/game" element={<GamePage />} />
          <Route path="/auth/callback" element={<AuthCallback />} />
          <Route path="/multiplayer" element={<GameRoomPage />} />
          <Route path="/user/profile" element={<ProfilePage />} />
          <Route path="/users/admin" element={<AdminUserList/>} />
        </Routes>
      </UserProvider>
    </Router>
  </StrictMode>
);


