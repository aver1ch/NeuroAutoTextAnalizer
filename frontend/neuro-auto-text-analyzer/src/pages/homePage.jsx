import React, { useState, useRef, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import '../styles/stylesOfPages/HomePage.css';

const Home = () => {
  const navigate = useNavigate();
  const [menuOpen, setMenuOpen] = useState(false);
  const [sideBarOpen, setSideBarOpen] = useState(false);
  const menuRef = useRef(null);
  const sideBarRef = useRef(null);
  const login = localStorage.getItem('login');

  const handleLogout = () => {
    localStorage.removeItem('auth');
    localStorage.removeItem('login');
    navigate('/login');
  };

  useEffect(() => {
    const handleClickOutside = (event) => {
      if (menuRef.current && !menuRef.current.contains(event.target)) {
        setMenuOpen(false);
      }
      if (sideBarRef.current && !sideBarRef.current.contains(event.target) && !event.target.closest('.sidebar-toggle')) {
        setSideBarOpen(false);
      }
    };
    document.addEventListener('mousedown', handleClickOutside);
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, []);

  return (

    <div className={`home-wrapper ${sideBarOpen ? 'sidebar-open' : ''}`}>
    <header className="top-bar">
      <input
        type="text"
        placeholder="Поиск"
        className="search-input"
      />
      <button
        className="sidebar-toggle"
        onClick={() => setSideBarOpen(!sideBarOpen)}
        aria-label="Toggle sidebar"
      >
        ☰
      </button>

      <div className="header-icons">
            <span className="icon" title="Помощь">❓</span>

            <div className="account-menu" ref={menuRef}>
              <span
                className="icon"
                style={{ cursor: 'pointer' }}
                onClick={() => setMenuOpen(!menuOpen)}
                title="Аккаунт"
              >
                👤
              </span>

              {menuOpen && (
                <div className="dropdown-menu">
                  <div className="dropdown-login">Логин: {login}</div>

                  <button className="logout-button" onClick={handleLogout}>
                    Выйти
                  </button>
                </div>
              )}
            </div>
          </div>
    </header>


      <div className="main-content">

      <aside className="sidebar" ref={sideBarRef}>
        <div className="menu-item">
          <span className="icon">📄</span>
          <span>Свойства документа</span>
        </div>
        <div className="menu-item">
          <span className="icon">🗂️</span>
          <span>Структура документа</span>
        </div>
        <div className="menu-item">
          <span className="icon">📏</span>
          <span>Текстовые метрики</span>
        </div>
        <div className="menu-item">
          <span className="icon">⚙️</span>
          <span>Параметры экспертизы</span>
        </div>
        <div className="menu-item">
          <span className="icon">📐</span>
          <span>Размер документа</span>
        </div>
      </aside>


        <div className="home-container">
          <h1 className="home-title">Добро пожаловать, {login}!</h1>
          <p className="home-subtitle">Вы успешно вошли в систему.</p>

          <div className="card-grid">
            <div className="card">
              <h2 className="card-title">Свойства документа</h2>
              <p>Просмотр и редактирование основных метаданных документа.</p>
            </div>

            <div className="card">
              <h2 className="card-title">Загрузка файла</h2>
              <p>Перетащите файл или выберите из проводника для анализа.</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Home;
