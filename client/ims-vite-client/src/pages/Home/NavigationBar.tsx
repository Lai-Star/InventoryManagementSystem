import Link from '../../component/Link';
import './NavigationBar.css';

const NavigationBar = () => {
  return (
    <header>
      <h4 className="logo">
        <Link to={'/home'}>IMS</Link>
      </h4>
      <nav>
        <ul className="nav__links">
          <li>
            <Link to={'/accounts'}>Accounts</Link>
          </li>
          <li>
            <Link to={'/transactions'}>Transactions</Link>
          </li>
          <li>
            <Link to={'/inventory'}>Inventory</Link>
          </li>
          <li>
            <Link to={'/products'}>Products</Link>
          </li>
          <li>
            <Link to={'/sales'}>Sales</Link>
          </li>
        </ul>
      </nav>
    </header>
  );
};

export default NavigationBar;
