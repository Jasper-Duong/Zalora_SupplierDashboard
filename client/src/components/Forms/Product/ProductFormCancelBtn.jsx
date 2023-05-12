import { Button } from "antd";
import React from "react";
import { Link } from "react-router-dom";
import { routes } from "../../../constants/routes";

export default function ProductFormCancelBtn() {
  return (
    <Link to={routes.PRODUCT_TABLE.path}>
      <Button>Cancel</Button>
    </Link>
  );
}
