import styles from './Table.module.css';
import { Data, ConfigItem, TableProps } from '../models/table-model';

function Table({ config, data, keyFn }: TableProps): React.ReactElement {
  // Table Headers
  const renderedHeaders = config.map(
    (column: ConfigItem): React.ReactElement => {
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
      (column: ConfigItem): React.ReactElement => {
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
