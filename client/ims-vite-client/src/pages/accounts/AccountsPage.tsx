import { getAccountsRoute } from '../../golang-api/accounts';

function Accounts() {
  const handleClick = async () => {
    const response = await getAccountsRoute();

    console.log(response);
  };

  return (
    <div>
      <button onClick={handleClick}>Test Get Accounts</button>
    </div>
  );
}

export default Accounts;
