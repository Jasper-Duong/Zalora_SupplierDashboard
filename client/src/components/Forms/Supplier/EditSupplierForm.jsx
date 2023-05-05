import React from "react";
import SupplierForm from "./SupplierForm";

export default function EditSupplierForm({ closeModal, supplier }) {
  const onFinish = (values) => {
    console.log({ values });
    closeModal();
  };
  return (
    <SupplierForm
      onFinish={onFinish}
      closeModal={closeModal}
      supplier={supplier}
      submitBtnText="Update"
    />
  );
}
