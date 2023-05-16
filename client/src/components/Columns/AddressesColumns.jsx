import React from "react";
import TableFormActions from "../../modules/TableForm/Cell/TableFormActions";

const inputTypeMap = { type: "select", address_info: "input" };

export default function AddressesColumns(editingRow) {
  const columns = [
    {
      title: "Type",
      dataIndex: "type",
      key: "type",
      width: "25%",
      editable: false,
    },
    {
      title: "Address Info",
      dataIndex: "address_info",
      key: "address_info",
      width: "45%",
      editable: true,
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
