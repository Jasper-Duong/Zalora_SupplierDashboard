import { message } from 'antd';
import React from 'react'
import SupplierForm from './SupplierForm';
import { useNavigate } from "react-router-dom";

import { addSupplierByIdApi } from '../../../services/supplier'

export default function AddSupplierForm() {
  const navigate = useNavigate();
  const onFinish = async (values) => {
    try {
      await addSupplierByIdApi(values)
      message.success("Added Supplier");
      setTimeout(() => navigate("/suppliers/table"), 300);
    } catch (err) {
      message.error(err)
    }
  };
  return (
    <SupplierForm
      onFinish={onFinish}
      submitBtnText={"Add new"}
    />
  );
}
