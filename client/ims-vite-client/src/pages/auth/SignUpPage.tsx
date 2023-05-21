import { useState } from 'react';
import { signUpRoute } from '../../golang-api/auth';
import useNavigation from '../../hooks/use-navigation';
import { AxiosError } from 'axios';
import Button from '../../component/Button';

function SignUpPage() {
  const [username, setUsername] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [email, setEmail] = useState<string>('');
  const [errorMsg, setErrorMsg] = useState<string>('');

  const { navigate } = useNavigation()!;

  const handleSubmit = async (event: React.FormEvent): Promise<void> => {
    event.preventDefault();

    try {
      const response = await signUpRoute(username, password, email);
      setUsername('');
      setPassword('');
      setEmail('');
      if (response.data.Status) {
        navigate('/');
      }
    } catch (error) {
      if (error instanceof AxiosError) {
        setErrorMsg(error.response?.data?.Err || 'An unknown error occurred.');
      } else {
        setErrorMsg('An unknown error occurred.');
      }
    }
  };

  const handleClick = () => [navigate('/')];

  return (
    <div className="centered-container">
      <h2>Register Account</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Username</label>
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            autoFocus
          />
        </div>
        <div>
          <label>Password</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            autoFocus
          />
        </div>
        <div>
          <label>Email</label>
          <input
            type="text"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            autoFocus
          />
        </div>
        {errorMsg && <div>{errorMsg}</div>}
        <Button success rounded type="submit">
          Create Account
        </Button>
      </form>
      <Button light rounded onClick={handleClick}>
        Back to Login Page
      </Button>
    </div>
  );
}

export default SignUpPage;
