import { Button } from "antd";
import React from "react";

export default function EditSupplierBtn({ showModal }) {
  const handleClick = () => {
    showModal();
  };
  return (
    <Button onClick={handleClick} type="primary">
      Edit Supplier
    </Button>
  );
}
