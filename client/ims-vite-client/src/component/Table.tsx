import styles from './Table.module.css';
import { ConfigItem, TableProps } from '../models/table-model';

function Table<Data>({ config, data, keyFn }: TableProps<Data>): React.ReactElement {
  // Table Headers
  const renderedHeaders = config.map(
    (column: ConfigItem<Data>): React.ReactElement => {
      return (
        <th key={column.label} className={styles.table_header}>
          {column.label}
        </th>
      );
    }
  );

  // Table Rows
  const renderedRows = data.map((rowData: Data): React.ReactElement => {
    const renderedCells = config.map(
      (column: ConfigItem<Data>): React.ReactElement => {
        return (
          <td key={column.label} className={styles.table_cell}>
            {column.render(rowData)}
          </td>
        );
      }
    );

    return (
      <tr key={keyFn(rowData)} className={styles.table_row}>
        {renderedCells}
      </tr>
    );
  });

  return (
    <table>
      <thead>
        <tr>{renderedHeaders}</tr>
      </thead>
      <tbody>{renderedRows}</tbody>
    </table>
  );
}

export default Table;
