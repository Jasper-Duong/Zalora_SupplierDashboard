import React from "react";
import SupplierForm from "./SupplierForm";
import { useParams } from "react-router-dom";
import { message } from "antd";

import { getSupplierByIdApi, updateSupplierByIdApi } from "../../../services/supplier";
import useService from "../../../hooks/useService";

export default function EditSupplierForm() {
  const params = useParams()
  const supplier = useService({service: () => getSupplierByIdApi(params.id)})
  const onFinish = async (values) => {
    console.log(values);
    try {
      await updateSupplierByIdApi(params.id, values)
      message.success("Successfully Updated Supplier");
    } catch (err) {
      message.error("Failed" + err)
    }
  };
  return (
    <SupplierForm
      onFinish={onFinish}
      supplier={supplier?.data}
      submitBtnText="Update"
    />
  );
}
