import React from "react";
import ProductForm from "./ProductForm";
import { message } from "antd";

export default function EditProductForm({ product }) {
  const onFinish = (values) => {
    console.log({ values });
    message.success("Successfully Updated Product");
  };
  return (
    <ProductForm
      product={product}
      onFinish={onFinish}
      submitBtnText={"Update"}
    />
  );
}
