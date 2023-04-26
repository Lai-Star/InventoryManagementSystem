import { useState } from 'react'
import './Login.css'

function Login() {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")

  const handleEmailChange = (event) => {
    setEmail(event.target.value)
  }

  const handlePasswordChange = (event) => {
    setPassword(event.target.value)
  }

  const handleSubmit = (event) => {
    event.preventDefault()

    console.log(email, password)
    setEmail("")
    setPassword("")
  }

  return (
    <div className="login-container">
      <form onSubmit={handleSubmit}>
        <h2 className="login-title">Login</h2>
        <div>
          <label>Email</label>
          <input value={email} onChange={handleEmailChange} autoFocus />
        </div>
        <div>
          <label>Password</label>
          <input value={password} onChange={handlePasswordChange} />
        </div>
        <button type="submit">Login</button>
      </form>
    </div>
  );
}

export default Login;
