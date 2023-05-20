import './App.css';
import Route from './component/Route';
import Login from './pages/auth/LoginPage';
import SignUp from './pages/auth/SignUpPage';
import Home from './pages/Home/HomePage';

function App() {
  return (
    <div>
      <Route path="/">
        <Login />
      </Route>
      <Route path="/signup">
        <SignUp />
      </Route>
      <Route path="/home">
        <Home />
      </Route>
    </div>
  );
}

export default App;
