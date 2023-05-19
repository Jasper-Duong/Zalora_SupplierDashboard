import React from "react";
import ProductForm from "./ProductForm";
import { message } from "antd";
import { useNavigate } from "react-router-dom";

import { addProductByIdApi } from '../../../services/product'

export default function AddProductForm() {
  const navigate = useNavigate();
  const onFinish = async (values) => {
    console.log(values)
    try {
      await addProductByIdApi(values)
      setTimeout(() => navigate("/"), 300);
      message.success("Added Product");
    } catch (err) {
      message.error("Failed")
    }
  };
  return (
    <ProductForm
      onFinish={onFinish}
      submitBtnText={"Add Product"}
    />
  );
}
