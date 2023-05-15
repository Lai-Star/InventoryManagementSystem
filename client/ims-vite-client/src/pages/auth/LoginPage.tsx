import { useState } from 'react';
import './Login.css';
import loginRoute from '../../golang-api/auth';
import { AxiosError } from 'axios';
import useNavigation from '../../hooks/use-navigation';

function LoginPage() {
  const [username, setUsername] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [errorMsg, setErrorMsg] = useState<string>('');

  const { navigate } = useNavigation()!;

  const handleUsernameChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ): void => {
    setUsername(event.target.value);
  };

  const handlePasswordChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(event.target.value);
  };

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();

    try {
      const response = await loginRoute(username, password);
      setUsername('');
      setPassword('');
      setErrorMsg('');
      if (response.data.Status) {
        navigate('/signup');
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
    <div className="login-container">
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
        <button type="submit">Login</button>
      </form>
    </div>
  );
}

export default LoginPage;
