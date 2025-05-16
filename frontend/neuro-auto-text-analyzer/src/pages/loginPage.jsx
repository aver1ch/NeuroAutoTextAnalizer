import React from 'react';
import './LoginPage.css'; // если будешь использовать внешний CSS

const LoginPage = () => {
  return (
    <div className="login-container">
      <h2 className="login-title">Вход</h2>

      <div className="form-group">
        <label htmlFor="login">Логин:</label>
        <input type="text" id="login" name="login" />
      </div>

      <div className="form-group">
        <label htmlFor="password">Пароль:</label>
        <input type="password" id="password" name="password" />
      </div>
    </div>
  );
};

export default LoginPage;
