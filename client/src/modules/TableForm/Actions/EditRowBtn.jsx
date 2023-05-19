import Typography from "antd/es/typography/Typography";
import React from "react";

export default function EditRowBtn({ editingRow, handleEdit }) {
  return (
    <Typography.Link disabled={editingRow !== ""} onClick={handleEdit}>
      Edit
    </Typography.Link>
  );
}
