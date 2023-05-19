import React, { useState } from "react";
import ProductStockForm from "../Forms/Stock/ProductStockForm";
import EditProductStockBtns from "./EditProductStockBtns";
import OnEditProductStockBtns from "./OnEditProductStockBtns";

export default function EditProductStockForm({ supplierStock }) {
  const [isEdit, setIsEdit] = useState(false);
  return (
    <ProductStockForm
      supplierStock={supplierStock}
      isEdit={isEdit}
      ProductStockBtns={
        isEdit ? (
          <OnEditProductStockBtns setIsEdit={setIsEdit} />
        ) : (
          <EditProductStockBtns setIsEdit={setIsEdit} />
        )
      }
    />
  );
}
