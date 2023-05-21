import { useState } from 'react';
import './Login.css';
import { loginRoute } from '../../golang-api/auth';
import { AxiosError } from 'axios';
import useNavigation from '../../hooks/use-navigation';
import Link from '../../component/Link';
import Button from '../../component/Button';

function LoginPage() {
  const [username, setUsername] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [errorMsg, setErrorMsg] = useState<string>('');
  const SIGN_UP_PATH = '/signup';

  const { navigate } = useNavigation()!;

  const handleUsernameChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ): void => {
    setUsername(event.target.value);
  };

  const handlePasswordChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(event.target.value);
  };

  const handleSubmit = async (event: React.FormEvent): Promise<void> => {
    event.preventDefault();

    try {
      const response = await loginRoute(username, password);
      setUsername('');
      setPassword('');
      setErrorMsg('');
      if (response.data.Status) {
        navigate('/home');
      }
    } catch (error) {
      if (error instanceof AxiosError) {
        setUsername('');
        setPassword('');
        setErrorMsg(error.response?.data?.Err || 'An unknown error occurred.');
      } else {
        setErrorMsg('An unknown error occurred.');
      }
    }
  };

  return (
    <div className="centered-container">
      <form onSubmit={handleSubmit}>
        <h2 className="login-title">Login</h2>
        <div>
          <label>Username</label>
          <input value={username} onChange={handleUsernameChange} autoFocus />
        </div>
        <div>
          <label>Password</label>
          <input
            type="password"
            value={password}
            onChange={handlePasswordChange}
          />
        </div>
        {errorMsg && <div className="error-message">{errorMsg}</div>}
        <Button success rounded type="submit">
          Login
        </Button>
      </form>
      <Link to={SIGN_UP_PATH} className="" activeClassName="">
        Don't have an account? Click here to sign up!
      </Link>
    </div>
  );
}

export default LoginPage;
