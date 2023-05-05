import React from "react";
import ProductForm from "./ProductForm";
import { message } from "antd";

export default function EditProductForm({ closeModal, product }) {
  const onFinish = (values) => {
    console.log({ values });
    message.success("Successfully Updated Product");
    closeModal();
  };
  return (
    <ProductForm
      closeModal={closeModal}
      product={product}
      onFinish={onFinish}
      submitBtnText={"Update"}
    />
  );
}
