import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

function Signup() {
  const [form, setForm] = useState({
    firstName: '',
    lastName: '',
    email: '',
    password: '',
    phone: '',
    userType: 'USER',
  });

  const [errorMessage, setErrorMessage] = useState('');
  const navigate = useNavigate();

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleRegisterSubmit = async (e) => {
    e.preventDefault();

    const userData = {
      First_name: form.firstName,
      Last_name: form.lastName,
      Email: form.email,
      Password: form.password,
      Phone: form.phone,
      User_type: form.userType,
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

      if (!response.ok) {
        setErrorMessage(data.error || '註冊失敗，請再試一次');
        return;
      }

      setErrorMessage('');
      navigate('/login', { state: { successMessage: '註冊成功！請登入您的帳號' } });
    } catch (error) {
      setErrorMessage(error.message || '發生錯誤，請稍後再試');
    }
  };

  return (
    <div className="container d-flex justify-content-center align-items-center min-vh-100">
      <div className="card p-4 shadow text-center" style={{ width: '400px' }}>
        <h2 className="mb-3">註冊</h2>

        {errorMessage && (
          <div className="alert alert-danger" role="alert">
            {errorMessage}
          </div>
        )}

        <form onSubmit={handleRegisterSubmit}>
          {[
            { label: '名字', name: 'firstName', type: 'text' },
            { label: '姓氏', name: 'lastName', type: 'text' },
            { label: 'Email', name: 'email', type: 'email' },
            { label: '密碼', name: 'password', type: 'password' },
            { label: '電話', name: 'phone', type: 'text' },
          ].map(({ label, name, type }) => (
            <div key={name} className="mb-3 text-start">
              <label className="form-label">{label}</label>
              <input
                type={type}
                name={name}
                className="form-control"
                value={form[name]}
                onChange={handleChange}
                required
              />
            </div>
          ))}

          <div className="mb-3 text-start">
            <label className="form-label">用戶類型</label>
            <select
              name="userType"
              className="form-control"
              value={form.userType}
              onChange={handleChange}
              required
            >
              <option value="USER">USER</option>
              <option value="ADMIN">ADMIN</option>
            </select>
          </div>

          <button type="submit" className="btn btn-primary w-100">註冊</button>
          <button
            type="button"
            className="btn btn-outline-secondary w-100 mt-2"
            onClick={() => navigate('/login')}
          >
            返回登入頁
          </button>
        </form>
      </div>
    </div>
  );
}

export default Signup;
