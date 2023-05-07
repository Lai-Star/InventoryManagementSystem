import { createContext, useState, useEffect } from 'react';

const NavigationContext = createContext();

function NavigationProvider({ children }) {
  const [currentPath, setCurrentPath] = useState(window.location.pathname);

  useEffect(() => {
    // to handle users clicking forward and back buttons on browser
    const handler = () => {
      setCurrentPath(window.location.pathname);
    };

    window.addEventListener('popstate', handler);

    // cleanup
    return () => {
      window.removeEventListener('popstate', handler);
    };
  }, []);

  // for link navigation
  const navigate = (to) => {
    window.history.pushState({}, '', to);
    setCurrentPath(to);
  };

  return (
    <NavigationProvider value={{ currentPath, navigate }}>
      {children}
    </NavigationProvider>
  );
}

export { NavigationProvider };
export default NavigationContext;
