import { useState } from 'react';
import reactLogo from './assets/react.svg';
import viteLogo from '/vite.svg';
import './App.css';
import Login from './pages/auth/Login';
import { NavigationProvider } from './context/navigation';

function App() {
  return (
    <NavigationProvider>
      <Login />
    </NavigationProvider>
  );
}

export default App;
