import { message } from 'antd';
import React from 'react'
import SupplierForm from './SupplierForm';

export default function AddSupplierForm({ closeModal }) {
  const onFinish = (supplier) => {
    console.log({ supplier });
    message.success("Added new Supplier!");
    closeModal();
  };
  return (
    <SupplierForm
      onFinish={onFinish}
      closeModal={closeModal}
      submitBtnText={"Add new"}
    />
  );
}
