import Table from './Table';
import { TableProps, ConfigItem, SortProps } from '../models/table-model';
import { GoArrowSmallDown, GoArrowSmallUp } from 'react-icons/go';
import useSort from '../hooks/use-sort';
import { SortType } from '../hooks/use-sort';

interface SortCustomHook<Data> {
  sortOrder: string | null;
  sortBy: SortType;
  sortedData: Data[];
  setSortColumn: (label: string) => void;
}

function SortableTable<Data>(props: SortProps<Data>) {
  const { config, data }: { config: ConfigItem<Data>[]; data: Data[] } = props;
  const { sortOrder, sortBy, sortedData, setSortColumn }: SortCustomHook<Data> = useSort(
    config,
    data
  );

  const updatedConfig: (ConfigItem<Data> | void)[] = config.map(
    (column: ConfigItem<Data>) => {
      if (!column.sortValue) {
        return; // returning null when column.sortValue is falsy
      }

      return {
        ...column,
        header: (): React.ReactNode => {
          return (
            <th onClick={() => setSortColumn(column.label)}>
              <div></div>
            </th>
          );
        },
      };
    }
  );
}

const getIcons = (
  label: string,
  sortBy: string,
  sortOrder: string
): React.ReactElement => {
  // the selected column is currently not sorted
  if (label !== sortBy) {
    return (
      <div>
        <GoArrowSmallUp />
        <GoArrowSmallDown />
      </div>
    );
  }

  if (sortOrder === null) {
    // data not sorted yet
    return (
      <div>
        <GoArrowSmallUp />
        <GoArrowSmallDown />
      </div>
    );
  } else if (sortOrder === 'asc') {
    return (
      <div>
        <GoArrowSmallUp />
      </div>
    );
  } else if (sortOrder === 'desc') {
    return (
      <div>
        <GoArrowSmallDown />
      </div>
    );
  }

  throw new Error(`Invalid sortOrder: ${sortOrder}`);
};

export default SortableTable;
