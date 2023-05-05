import React from "react";
import ProductForm from "./ProductForm";
import { message } from "antd";

export default function AddProductForm({ closeModal }) {
  const onFinish = (values) => {
    message.success("Added Product");
    closeModal();
  };
  return (
    <ProductForm
      onFinish={onFinish}
      closeModal={closeModal}
      submitBtnText={"Add Product"}
    />
  );
}
