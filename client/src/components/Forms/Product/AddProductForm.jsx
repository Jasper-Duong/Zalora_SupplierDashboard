import React from "react";
import ProductForm from "./ProductForm";
import { message } from "antd";
import { useNavigate } from "react-router-dom";

export default function AddProductForm() {
  const navigate = useNavigate();
  const onFinish = (values) => {
    message.success("Added Product");
    setTimeout(() => navigate("/"), 300);
  };
  return (
    <ProductForm
      onFinish={onFinish}
      submitBtnText={"Add Product"}
    />
  );
}
