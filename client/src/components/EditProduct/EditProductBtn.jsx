import { Button } from "antd";
import React from "react";
import { useNavigate } from "react-router-dom";
import { routes } from "../../constants/routes";
import { EditOutlined } from "@ant-design/icons";

export default function EditProductBtn({ data }) {
  const navigate = useNavigate();
  const handleClick = () => {
    navigate(routes.EDIT_PRODUCT(data.id).path);
  };
  return (
    <Button type="primary" onClick={handleClick} icon={<EditOutlined />} />
  );
}
