// import React from 'react';
// import useNavigation from '../hooks/use-navigation';

// function Link({ to, children, className, activeClassName }) {
//   const { navigate, currentPath } = useNavigation();

//   const handleClick = (event: React.MouseEvent<HTMLElement>) => {
//     // to open up new window with CTRL + Enter
//     if (event.metaKey || event.ctrlKey) {
//       return;
//     }

//     event.preventDefault();

//     navigate(to);
//   };

//   return (
//     <a className={classes} href={to} onClick={handleClick}>
//       {children}
//     </a>
//   );
// }
