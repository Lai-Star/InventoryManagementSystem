import React, { createContext, useState, useEffect } from 'react';

export type NavigationContextType = {
  currentPath: string;
  navigate: (to: string) => void;
};

interface Props {
  children: React.ReactNode;
}

const NavigationContext = createContext<NavigationContextType | null>(null);

const NavigationProvider: React.FC<Props> = ({ children }) => {
  const [currentPath, setCurrentPath] = useState<string>(
    window.location.pathname
  );

  useEffect(() => {
    // to handle users clicking forward and back buttons on browser
    const handler = (): void => {
      setCurrentPath(window.location.pathname);
    };

    window.addEventListener('popstate', handler);

    // cleanup
    return () => {
      window.removeEventListener('popstate', handler);
    };
  }, []);

  // for link navigation
  const navigate = (to: string): void => {
    window.history.pushState({}, '', to);
    setCurrentPath(to);
  };

  const contextValue: NavigationContextType = {
    currentPath,
    navigate,
  };

  return (
    <NavigationContext.Provider value={contextValue}>
      {children}
    </NavigationContext.Provider>
  );
};

export { NavigationProvider };
export default NavigationContext;
