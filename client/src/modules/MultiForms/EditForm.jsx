import React, { useState } from "react";
import BaseForm from "./BaseForm";
import OnEditBtns from "./Buttons/OnEditBtns";
import EditBtns from "./Buttons/EditBtns";
import { deleteStockApi, updateStockApi } from "../../services/stock";

export default function EditForm({ item, forceRender }) {
  // getItemById(item.id)

  const [isEdit, setIsEdit] = useState(false);
  const initialValues = { name0: item.id, name1: item.name, name2: item.stock };
  const handleUpdate = (values) => {
    // PUT API
    const submitData = {
      id: values.name0,
      name: values.name1,
      stock: +values.name2,
    };
    console.log("Updating ", submitData);
    updateStockApi(1, item.id, submitData);
    setIsEdit(false);
    forceRender();
  };
  const handleDelete = () => {
    // DELETE API
    console.log("Deleting ", item);
    deleteStockApi(1, item.id);
    forceRender();
  };
  const ActionBtns = () =>
    isEdit ? (
      <OnEditBtns setIsEdit={setIsEdit} />
    ) : (
      <EditBtns handleDelete={handleDelete} setIsEdit={setIsEdit} />
    );
  return (
    <BaseForm
      initialValues={initialValues}
      isEdit={isEdit}
      ActionBtns={<ActionBtns />}
      onFinish={handleUpdate}
    />
  );
}
