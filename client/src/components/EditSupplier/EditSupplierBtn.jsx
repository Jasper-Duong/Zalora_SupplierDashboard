import { Button } from "antd";
import React from "react";
import { useDispatch } from "react-redux";
import { setSelectedSupplier } from "../../store/actions/suppliers.action";

export default function EditSupplierBtn({ showModal, data }) {
  const dispatch = useDispatch();
  const handleClick = () => {
    dispatch(setSelectedSupplier(data));
    showModal();
  };
  return (
    <Button onClick={handleClick} type="primary">
      Edit Supplier
    </Button>
  );
}
