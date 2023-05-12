import React from "react";
import BaseForm from "./BaseForm";
import AddBtns from "./Buttons/AddBtns";
import { addStockByApi } from "../../services/stock";

export default function AddForm({ setIsAdd, forceRender }) {
  const handleAdd = (values) => {
    const {name, stock} = values;
    addStockByApi(1, name, stock);
    console.log("Adding ", values);
    forceRender();
  };
  return (
    <BaseForm
      onFinish={handleAdd}
      isEdit={true}
      ActionBtns={<AddBtns setIsAdd={setIsAdd} />}
    />
  );
}
