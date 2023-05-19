import React from "react";
import ProductStockForm from "../Forms/Stock/ProductStockForm";
import AddProductStockBtns from "./AddProductStockBtns";

export default function AddProductStock({ setIsEdit }) {
  return (
    <ProductStockForm
      isEdit={true}
      ProductStockBtns={<AddProductStockBtns setIsEdit={setIsEdit} />}
    />
  );
}
