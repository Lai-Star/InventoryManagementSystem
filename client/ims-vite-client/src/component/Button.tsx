import className from 'classnames';
import styles from './Button.module.css';

type Props = {
  children: React.ReactNode;
  primary?: boolean;
  secondary?: boolean;
  success?: boolean;
  warning?: boolean;
  danger?: boolean;
  outline?: boolean;
  rounded?: boolean;
  [rest: string]: any;
};

const Button: React.FC<Props> = ({
  children,
  primary,
  secondary,
  success,
  warning,
  danger,
  info,
  light,
  dark,
  outline,
  rounded,
  ...rest
}) => {
  // Number(!!undefined) = 0
  const count =
    Number(!!primary) +
    Number(!!secondary) +
    Number(!!success) +
    Number(!!warning) +
    Number(!!danger);

  // Checks if count > 1. Only 1 button type can be specified for each button component
  // This error will only be thrown during develop on button component creation
  if (count > 1) {
    throw new Error(
      'Only 1 of primary, secondary, success, warning, danger, info, light and dark can be true!'
    );
  }

  const classes: string = className(
    rest.className,
    'flex items-center px-3 py-1.5 border',
    {
      [styles.primary]: primary,
      [styles.secondary]: secondary,
      [styles.success]: success,
      [styles.warning]: warning,
      [styles.danger]: danger,
      [styles.info]: info,
      [styles.light]: light,
      [styles.dark]: dark,
      [styles.rounded]: rounded,
      [styles.outline]: outline,
      [styles.primary_text]: outline && primary,
      [styles.secondary_text]: outline && secondary,
      [styles.success_text]: outline && success,
      [styles.warning_text]: outline && warning,
      [styles.danger_text]: outline && danger,
      [styles.info_text]: outline && info,
      [styles.light_text]: outline && light,
      [styles.dark_text]: outline && dark,
    }
  );

  return (
    <button {...rest} className={classes}>
      {children}
    </button>
  );
};

export default Button;
