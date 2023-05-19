import React from "react";
import BaseForm from "./BaseForm";
import AddBtns from "./Buttons/AddBtns";
import { addStockByApi } from "../../services/stock";

export default function AddForm({ setIsAdd, forceRender }) {
  const handleAdd = (values) => {
    const submitData = { stock: values.name2 };
    setIsAdd(false);
    addStockByApi(1, 0, submitData);
    console.log("Adding ", submitData);
    forceRender();
  };
  return (
    <BaseForm
      isEdit
      isAdd
      onFinish={handleAdd}
      ActionBtns={<AddBtns setIsAdd={setIsAdd} />}
    />
  );
}
