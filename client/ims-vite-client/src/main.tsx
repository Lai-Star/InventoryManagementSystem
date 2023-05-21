import ReactDOM from 'react-dom/client';
import App from './App.tsx';
import './index.css';
import { NavigationProvider } from './context/navigation.tsx';

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <NavigationProvider>
    <App />
  </NavigationProvider>
);
