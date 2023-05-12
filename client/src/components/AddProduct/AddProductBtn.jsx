import { Button } from "antd";
import React from "react";
import { useNavigate } from "react-router-dom";

export default function AddProductBtn() {
  const navigate = useNavigate();
  const handleClick = () => {
    navigate("/products/add-product");
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
