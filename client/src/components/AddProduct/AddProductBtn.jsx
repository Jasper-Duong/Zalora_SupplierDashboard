import { Button } from "antd";
import React from "react";

export default function AddProductBtn({ showModal }) {
  const handleClick = () => {
    showModal();
  };
  return (
    <Button
      style={{ float: "left", marginBottom: "1rem" }}
      className="add-product-btn"
      onClick={handleClick}
    >
      Add Product
    </Button>
  );
}
