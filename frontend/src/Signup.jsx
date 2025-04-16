import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

function Signup() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');
  const [phone, setPhone] = useState('');
  const [userType, setUserType] = useState('USER');
  const [errorMessage, setErrorMessage] = useState(''); // 用來顯示錯誤訊息
  const navigate = useNavigate(); // 用來處理導航

  const handleRegisterSubmit = async (e) => {
    e.preventDefault();

    const userData = {
      First_name: firstName,
      Last_name: lastName,
      Password: password,
      Email: email,
      Phone: phone,
      User_type: userType,
    };

    try {
      const response = await fetch('http://localhost:3000/api/user/signup', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(userData),
      });

      const data = await response.json();

      // 如果後端回應錯誤
      if (!response.ok) {
        setErrorMessage(data.error || '註冊失敗，請再試一次');
        return;
      }

      // 註冊成功後處理
      setErrorMessage(''); // 清除錯誤訊息
      navigate('/', { state: { successMessage: '註冊成功！請登入您的帳號' } }); // 傳遞成功訊息到登入頁面
    } catch (error) {
      setErrorMessage(error.message || '發生錯誤，請稍後再試');
    }
  };

  return (
    <div className="container d-flex justify-content-center align-items-center min-vh-100">
      <div className="card p-4 shadow text-center" style={{ width: '400px' }}>
        <h2 className="mb-3">註冊</h2>
        {/* 顯示錯誤訊息 */}
        {errorMessage && (
          <div className="alert alert-danger" role="alert">
            {errorMessage}
          </div>
        )}

        <form onSubmit={handleRegisterSubmit}>
          <div className="mb-3 text-start">
            <label className="form-label">名字</label>
            <input
              type="text"
              className="form-control"
              value={firstName}
              onChange={(e) => setFirstName(e.target.value)}
              required
            />
          </div>
          <div className="mb-3 text-start">
            <label className="form-label">姓氏</label>
            <input
              type="text"
              className="form-control"
              value={lastName}
              onChange={(e) => setLastName(e.target.value)}
              required
            />
          </div>
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
          <div className="mb-3 text-start">
            <label className="form-label">電話</label>
            <input
              type="text"
              className="form-control"
              value={phone}
              onChange={(e) => setPhone(e.target.value)}
              required
            />
          </div>
          <div className="mb-3 text-start">
            <label className="form-label">用戶類型</label>
            <select
              className="form-control"
              value={userType}
              onChange={(e) => setUserType(e.target.value)}
              required
            >
              <option value="USER">USER</option>
              <option value="ADMIN">ADMIN</option>
            </select>
          </div>
          <button type="submit" className="btn btn-primary w-100">註冊</button>
        </form>
      </div>
    </div>
  );
}

export default Signup;
