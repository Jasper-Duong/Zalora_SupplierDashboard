import { Button } from "antd";
import React from "react";

export default function AddSupplierBtn({ showModal }) {
  const handleClick = () => {
    showModal();
  };
  return <Button onClick={handleClick}>Add Supplier</Button>;
}
