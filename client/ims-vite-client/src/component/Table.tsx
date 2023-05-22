
function Table({ config, data, keyFn }: any) {
  // Table Headers
  const renderedHeaders = config.map((column: any) => {
    return <th key={column.label}>{column.label}</th>;
  });

  // Table Rows
  const renderedRows = data.map((rowData: any) => {
    const renderedCells = config.map((column: any) => {
      return <td key={column.label}>{column.render(rowData)}</td>;
    });

    return <tr key={keyFn(rowData)}>{renderedCells}</tr>;
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
