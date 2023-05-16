import React from "react";
import ProductForm from "./ProductForm";
import { message } from "antd";
import { useParams } from "react-router-dom";

import { getProductByIdApi, updateProductByIdApi } from "../../../services/product";
import useService from "../../../hooks/useService";

export default function EditProductForm() {
  const params = useParams();
  const product = useService({service: () => getProductByIdApi(params.id)})
  const onFinish = async (values) => {
    console.log({ values });
    try {
      await updateProductByIdApi(params.id, values)
      message.success("Successfully Updated Product");
    } catch (err) {
      message.error("Failed" + err)
    }
  };

  return (
    <ProductForm
      product={product?.data}
      onFinish={onFinish}
      submitBtnText={"Update"}
    />
  );
}
