import './App.css';
import Route from './component/Route';
import Accounts from './pages/accounts/AccountsPage';
import Login from './pages/auth/LoginPage';
import SignUp from './pages/auth/SignUpPage';
import Home from './pages/Home/HomePage';
import NavigationBar from './pages/Home/NavigationBar';

function App() {
  // FOR DEV ONLY!!
  // SHIFTED NAVIGATION BAR HERE TEMPORARILY!!

  return (
    <div>
      <NavigationBar />
      <Route path="/">
        <Login />
      </Route>
      <Route path="/signup">
        <SignUp />
      </Route>
      <Route path="/home">
        <Home />
      </Route>
      <Route path="/accounts">
        <Accounts />
      </Route>
    </div>
  );
}

export default App;
