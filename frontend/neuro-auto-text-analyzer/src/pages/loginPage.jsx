import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import '../styles/stylesOfPages/LoginPage.css';

const LoginPage = () => {
  const [login, setLogin] = useState('');
  const [password, setPassword] = useState('');
  const [status, setStatus] = useState('');

  const navigate = useNavigate();

  const handleSubmit = (e) => {
    e.preventDefault();

    if (login === 'admin' && password === '1234') {
      localStorage.setItem('auth', 'true');
      localStorage.setItem('login', login);
      setStatus("Успешно!");
      navigate('/');
    } else if (login === '') {
      setStatus("Введите логин")
    } else if (password === '') {
      setStatus("Введите пароль")
    } else {
      setStatus('Неверный логин или пароль');
    }
  };

  return (
    <form onSubmit={handleSubmit} className="login-form">
      <p1 className="autorisation-text">Авторизация</p1>
      
      <input
        type="text"
        placeholder="Логин"
        value={login}
        onChange={e => setLogin(e.target.value)}
        className="login-input"
      />
      <input
        type="password"
        placeholder="Пароль"
        value={password}
        onChange={e => setPassword(e.target.value)}
        className="login-input"
      />
      <button type="submit" className="login-button">
        Войти
      </button>

      {status && <div className="status-message">{status}</div>}
    </form>
  );
};

export default LoginPage;
