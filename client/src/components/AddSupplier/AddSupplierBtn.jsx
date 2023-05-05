import { Button } from "antd";
import React from "react";

export default function AddSupplierBtn({ showModal }) {
  const handleClick = () => {
    showModal();
  };
  return (
    <Button
      style={{ float: "left", marginBottom: "1rem" }}
      onClick={handleClick}
    >
      Add Supplier
    </Button>
  );
}
