import { Input, Select } from "antd";
import React from "react";
import EditableCell from "../Rows/EditableRow";

const SelectAddressType = ({defaultValue}) => (
  <Select defaultValue={defaultValue}>
    <Select.Option value="office">Office</Select.Option>
    <Select.Option value="warehouse">Warehouse</Select.Option>
  </Select>
);

export default function AddressesCell({ inputType, record,...restProps }) {
  const inputNode = inputType === "input" ? <Input /> : <SelectAddressType defaultValue={record}/>;

  return <EditableCell {...restProps} inputNode={inputNode} />;
}
