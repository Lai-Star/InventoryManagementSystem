import { useState } from 'react';
import './Login.css';
import loginRoute from '../../golang-api/auth';

function Login() {
  const [username, setUsername] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [errorMsg, setErrorMsg] = useState<string>('');

  const handleUsernameChange = (event: React.ChangeEvent<HTMLInputElement>): void => {
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
      console.log(response);
    } catch (error: any) {
      console.log(error.response.data.Err);
      setErrorMsg(error.response.data.Err);
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

export default Login;
