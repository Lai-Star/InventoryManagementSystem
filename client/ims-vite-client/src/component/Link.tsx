import React from 'react';
import useNavigation from '../hooks/use-navigation';
import classNames from 'classnames';
import styles from './Link.module.css';

interface Props {
  to: string;
  children: React.ReactNode;
  className: string;
  activeClassName: string;
}

function Link({ to, children, className, activeClassName }: Props) {
  const { navigate, currentPath } = useNavigation()!;

  const classes = classNames(
    styles.link_blue,
    styles.link_size,
    className,
    currentPath === to && activeClassName
  );

  const handleClick = (event: React.MouseEvent<HTMLElement>) => {
    // to open up new window with CTRL + Enter
    if (event.metaKey || event.ctrlKey) {
      return;
    }

    event.preventDefault();

    navigate(to);
  };

  return (
    <a className={classes} href={to} onClick={handleClick}>
      {children}
    </a>
  );
}

export default Link;
