import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App.tsx';
import './index.css';
import { NavigationProvider } from './context/navigation.tsx';

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <NavigationProvider>
      <App />
    </NavigationProvider>
  </React.StrictMode>
);

// import ReactDOM from 'react-dom/client';
// import App from './App.jsx';
// import { NavigationProvider } from './context/navigation.jsx';

// const el = document.getElementById('root');
// const root = ReactDOM.createRoot(el);

// root.render(
//   <NavigationProvider>
//     <App />
//   </NavigationProvider>
// );
