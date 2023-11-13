import { useState, useEffect } from 'react';
import { getAccountsRoute } from '../../golang-api/accounts';
import { ConfigItem, TableProps } from '../../models/table-model';
import Table from '../../component/Table';

type AccountsData = {
  UserId: number;
  Username: string;
  Email: string;
  IsActive: number;
  OrganisationName: string;
};

function Accounts(): React.ReactElement {
  const [data, setData] = useState<AccountsData[]>([]);

  useEffect(() => {
    const fetchAccountsData = async (): Promise<void> => {
      const response = await getAccountsRoute();
      setData(response.data.Result);
    };

    fetchAccountsData();
  }, []);

  const config: ConfigItem<AccountsData>[] = [
    {
      label: 'User ID',
      render: (account: AccountsData): number => account.UserId,
    },
    {
      label: 'Username',
      render: (account: AccountsData): string => account.Username,
    },
    {
      label: 'Email',
      render: (account: AccountsData): string => account.Email,
    },
    {
      label: 'IsActive',
      render: (account: AccountsData): number => account.IsActive,
      sortValue: (account: AccountsData): number => account.IsActive,
    },
    {
      label: 'Organisation Name',
      render: (account: AccountsData): string => account.OrganisationName,
    },
  ];

  // Function used to generate key for mapping function in rows in Table component
  const keyFn = (account: AccountsData): number => {
    return account.UserId;
  };

  const tableProps: TableProps<AccountsData> = {
    config,
    data,
    keyFn,
  };

  return (
    <div>
      <Table {...tableProps} />
    </div>
  );
}

export default Accounts;
