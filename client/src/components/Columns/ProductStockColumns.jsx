import React from 'react'
import TableFormActions from '../../modules/TableForm/Cell/TableFormActions';

export default function ProductStockColumns(editingRow) {
  const columns = [
    {
      title: "Supplier Name",
      dataIndex: "name",
      width: "45%",
      editable: false,
      key: "name",
    },
    {
      title: "Stock",
      dataIndex: "stock",
      width: "25%",
      editable: true,
      key: "stock",
    },
    {
      title: "Actions",
      dataIndex: "actions",
      key: "actions",
      render: (_, record) => <TableFormActions record={record} />,
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
            editing: editingRow === record.id,
          }),
        }
      : col
  );
}

const inputTypeMap = { stock: "number", name: "select" };
