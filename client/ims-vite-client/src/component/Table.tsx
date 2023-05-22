import styles from './Table.module.css'

function Table({ config, data, keyFn }: any) {
  // Table Headers
  const renderedHeaders = config.map((column: any) => {
    return (
      <th key={column.label} className={styles.table_header}>
        {column.label}
      </th>
    );
  });

  // Table Rows
  const renderedRows = data.map((rowData: any) => {
    const renderedCells = config.map((column: any) => {
      return (
        <td key={column.label} className={styles.table_cell}>
          {column.render(rowData)}
        </td>
      );
    });

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
