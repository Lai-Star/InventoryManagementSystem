import { useState } from 'react';
import { SortProps, ConfigItem } from '../models/table-model';

enum SORT_VALUE {
  ASC = 'asc',
  DESC = 'desc',
}

// Our table is currently capable of sorting based on these types
export type SortType = string | number | null;

function useSort<Data>(config: ConfigItem<Data>[], data: Data[]) {
  // 2 states to keep track of the sort order and which column we are sorting by
  const [sortOrder, setSortOrder] = useState<string | null>(null);
  const [sortBy, setSortBy] = useState<SortType>(null);

  // Sets the sort order for column and the current sort label
  const setSortColumn = (label: string): void => {
    // sort on a new column if the old column is currently sorting
    if (sortBy && label !== sortBy) {
      setSortOrder(SORT_VALUE.ASC);
      setSortBy(label);
      return;
    }

    if (sortOrder === null) {
      setSortOrder(SORT_VALUE.ASC);
      setSortBy(label);
    } else if (sortOrder === SORT_VALUE.ASC) {
      setSortOrder(SORT_VALUE.DESC);
      setSortBy(label);
    } else if (sortOrder === SORT_VALUE.DESC) {
      // In the last cycle, reset column to original state
      setSortOrder(null);
      setSortBy(null);
    }
  };

  // Only sort data if sortOrder && sortBy are not null
  // Make a copy of the `data` prop
  // Find the correct sortValue function and use it for sorting
  let sortedData = data;
  if (sortOrder && sortBy) {
    // find sortValue in config
    const column = config.find((column) => column.label === sortBy);

    if (column && column.sortValue) {
      const { sortValue } = column;

      sortedData = [...data].sort((a, b) => {
        const valueA = sortValue(a);
        const valueB = sortValue(b);

        const reverseOrder = sortOrder === SORT_VALUE.ASC ? 1 : -1;

        if (typeof valueA === 'string') {
          return valueA.localeCompare(valueB) * reverseOrder;
        } else {
          return (valueA - valueB) * reverseOrder;
        }
      });
    }
  }

  return { sortOrder, sortBy, sortedData, setSortColumn };
}

export default useSort;
