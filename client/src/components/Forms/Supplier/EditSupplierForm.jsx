import React from "react";
import SupplierForm from "./SupplierForm";

export default function EditSupplierForm({ supplier }) {
  const onFinish = (values) => {
    console.log({ values });
  };
  return (
    <SupplierForm
      onFinish={onFinish}
      supplier={supplier}
      submitBtnText="Update"
    />
  );
}
