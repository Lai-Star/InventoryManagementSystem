import { useState, useEffect } from 'react';
import { getAccountsRoute } from '../../golang-api/accounts';
import Table from '../../component/Table';

type Data = {
  UserId: number;
  Username: string;
  Email: string;
  IsActive: number;
  OrganisationName: string;
};

type ConfigItem = {
  label: string;
  render: (account: Data) => React.ReactNode;
};

function Accounts() {
  const [data, setData] = useState<Data[]>([]);

  useEffect(() => {
    const fetchAccountsData = async (): Promise<void> => {
      const response = await getAccountsRoute();
      setData(response.data.Result);
    };

    fetchAccountsData();
  }, []);

  const config: ConfigItem[] = [
    {
      label: 'User ID',
      render: (account) => account.UserId,
    },
    {
      label: 'Username',
      render: (account) => account.Username,
    },
    {
      label: 'Email',
      render: (account) => account.Email,
    },
    {
      label: 'IsActive',
      render: (account) => account.IsActive,
    },
    {
      label: 'Organisation Name',
      render: (account) => account.OrganisationName,
    },
  ];

  // Function used to generate key for mapping function in rows in Table component
  const keyFn = (account: any) => {
    return account.UserId;
  };

  return (
    <div>
      <Table config={config} data={data} keyFn={keyFn} />
    </div>
  );
}

export default Accounts;
