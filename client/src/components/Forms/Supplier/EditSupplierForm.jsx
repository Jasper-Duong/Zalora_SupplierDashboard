import SupplierForm from "./SupplierForm";
import { useParams, useNavigate } from "react-router-dom";
import { message } from "antd";

import {
  getSupplierByIdApi,
  updateSupplierByIdApi,
} from "../../../services/supplier";
import useService from "../../../hooks/useService";

export default function EditSupplierForm() {
  const params = useParams();
  const navigate = useNavigate()
  const supplier = useService({
    service: () => getSupplierByIdApi(params.id),
  });
  const maxCount = 2;
  const onFinish = async (values) => {
    console.log(values);
    try {
      await updateSupplierByIdApi(params.id, values);
      message.success({ content: "Successfully Updated Supplier", maxCount });
      navigate("/suppliers/table")
    } catch (err) {
      message.error({ content: "Failed" + err, maxCount });
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
