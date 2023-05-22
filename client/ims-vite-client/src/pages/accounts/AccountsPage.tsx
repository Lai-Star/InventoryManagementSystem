import { useState, useEffect } from 'react';
import { getAccountsRoute } from '../../golang-api/accounts';
import { Data, ConfigItem } from '../../models/table-model';
import Table from '../../component/Table';

function Accounts(): React.ReactElement {
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
      render: (account: Data): number => account.UserId,
    },
    {
      label: 'Username',
      render: (account: Data): string => account.Username,
    },
    {
      label: 'Email',
      render: (account: Data): string => account.Email,
    },
    {
      label: 'IsActive',
      render: (account: Data): number => account.IsActive,
    },
    {
      label: 'Organisation Name',
      render: (account: Data): string => account.OrganisationName,
    },
  ];

  // Function used to generate key for mapping function in rows in Table component
  const keyFn = (account: Data): number => {
    return account.UserId;
  };

  return (
    <div>
      <Table config={config} data={data} keyFn={keyFn} />
    </div>
  );
}

export default Accounts;
