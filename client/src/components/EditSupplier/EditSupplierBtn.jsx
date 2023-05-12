import { Button } from "antd";
import React from "react";
import { EditOutlined } from "@ant-design/icons";
import { useNavigate } from "react-router-dom";
import { routes } from "../../constants/routes";

export default function EditSupplierBtn({ data }) {
  const navigate = useNavigate();
  const handleClick = () => {
    navigate(routes.EDIT_SUPPLIER(data.id).path);
  };
  return (
    <Button onClick={handleClick} type="primary" icon={<EditOutlined />} />
  );
}
