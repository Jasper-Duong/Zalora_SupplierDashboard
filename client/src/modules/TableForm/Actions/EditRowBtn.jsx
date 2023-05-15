import Typography from "antd/es/typography/Typography";
import React from "react";

export default function EditRowBtn({ editingKey, handleEdit, record }) {
  return (
    <Typography.Link
      disabled={editingKey !== ""}
      onClick={() => handleEdit(record)}
    >
      Edit
    </Typography.Link>
  );
}
