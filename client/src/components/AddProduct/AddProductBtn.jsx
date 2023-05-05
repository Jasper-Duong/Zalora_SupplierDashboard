import { Button } from "antd";
import React from "react";

export default function AddProductBtn({ showModal }) {
  const handleClick = () => {
    showModal();
  };
  return <Button onClick={handleClick}>Add Product</Button>;
}
