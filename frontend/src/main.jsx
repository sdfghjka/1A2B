import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'; // 注意更新為 Routes 和 Route
import './index.css';
import App from './App.jsx';
import Signup from './Signup.jsx';
import GamePage from './GamePage.jsx';
import StartPage from './StartPage.jsx';
import AuthCallback from './AuthCallback.jsx';
import 'bootstrap/dist/css/bootstrap.min.css';
import GameRoomPage from './GameRoomPage.jsx';

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <Router>
      <Routes> {/* 使用 Routes 而不是 Switch */}
        <Route path="/login" element={<App />} />  {/* 使用 element 而不是 component */}
        <Route path="/signup" element={<Signup />} />  {/* 使用 element 而不是 component */}
        <Route path="/gamestart" element={<StartPage />} />
        <Route path="/game" element={<GamePage />} />
        <Route path="/auth/callback" element={<AuthCallback />} />
        <Route path="/multiplayer" element={<GameRoomPage />} /> 
      </Routes>
    </Router>
  </StrictMode>
);


