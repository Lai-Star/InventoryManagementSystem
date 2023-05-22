type Data = {
  UserId: number;
  Username: string;
  Email: string;
  IsActive: number;
  OrganisationName: string;
};

type TableProps = {
  config: ConfigItem[];
  data: Data[];
  keyFn: (rowData: Data) => number;
};

interface ConfigItem {
  label: string;
  render: (account: Data) => React.ReactNode;
}

export type { Data, ConfigItem, TableProps };
