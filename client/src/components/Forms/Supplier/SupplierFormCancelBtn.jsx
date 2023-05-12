import { Button } from "antd";
import React from "react";
import { routes } from "../../../constants/routes";
import { Link } from "react-router-dom";

export default function SupplierFormCancelBtn() {
  return (
    <Link to={routes.SUPPLIER_TABLE.path}>
      <Button>Cancel</Button>
    </Link>
  );
}
