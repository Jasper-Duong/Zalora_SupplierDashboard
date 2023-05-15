import TableFormActions from "../Cell/TableFormActions";

export default function TableFormColumns(editingKey) {
  const columns = [
    {
      title: "Product Name",
      dataIndex: "name",
      width: "45%",
      editable: false,
    },
    {
      title: "Stock",
      dataIndex: "stock",
      width: "25%",
      editable: true,
    },
    {
      title: "Actions",
      dataIndex: "actions",
      render: (_, record) => (
        <TableFormActions record={record} />
      ),
    },
  ];
  return columns.map((col) =>
    col.editable
      ? {
          ...col,
          onCell: (record) => ({
            record,
            inputType: inputTypeMap[col.dataIndex],
            dataIndex: col.dataIndex,
            title: col.title,
            editing: editingKey === record.key,
          }),
        }
      : col
  );
}

const inputTypeMap = { stock: "number", name: "select" };
