import React, { useContext } from "react";
import { TableFormContext } from "../../../contexts/TableFormContext";
import { InputNumber } from "antd";
import TableFormSelect from "../Input/TableFormSelect";
import EditableCell from "../../../modules/TableForm/Rows/EditableRow";

export default function SupplierStockCell({ inputType, ...rest }) {
  const { form } = useContext(TableFormContext);
  const inputNode =
    inputType === "number" ? <InputNumber /> : <TableFormSelect form={form} />;
  return <EditableCell {...rest} inputNode={inputNode} />;
}
