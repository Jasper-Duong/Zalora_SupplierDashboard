import { message } from 'antd';
import React from 'react'
import SupplierForm from './SupplierForm';

export default function AddSupplierForm() {
  const onFinish = (supplier) => {
    console.log({ supplier });
    message.success("Added new Supplier!");
  };
  return (
    <SupplierForm
      onFinish={onFinish}
      submitBtnText={"Add new"}
    />
  );
}
