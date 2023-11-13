type TableProps<Data> = {
  config: ConfigItem<Data>[];
  data: Data[];
  keyFn: (rowData: Data) => number;
};

interface ConfigItem<Data> {
  label: string;
  render: (account: Data) => React.ReactNode;
  sortValue?: (account: Data) => string | number;
}

type SortProps<Data> = {
  config: ConfigItem<Data>[];
  data: Data[];
};

export type { ConfigItem, TableProps, SortProps };
