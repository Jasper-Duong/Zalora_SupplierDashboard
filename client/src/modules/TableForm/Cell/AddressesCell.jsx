import { Input, Select } from "antd";
import React, { useContext } from "react";
import EditableCell from "../Rows/EditableRow";
import { TableFormContext } from "../../../contexts/TableFormContext";

const SelectAddressType = ({ defaultValue, form }) => {
  const onSelectChange = (value) => {
    form.setFieldsValue({ type: value });
  };
  console.log({ defaultValue });
  return (
    <Select defaultValue={defaultValue} onChange={onSelectChange}>
      <Select.Option value="office">office</Select.Option>
      <Select.Option value="warehouse">warehouse</Select.Option>
    </Select>
  );
};

export default function AddressesCell({ inputType, record, ...restProps }) {
  const { form } = useContext(TableFormContext);
  const inputNode =
    inputType === "input" ? (
      <Input />
    ) : (
      <SelectAddressType defaultValue={record?.type} form={form} />
    );

  return <EditableCell {...restProps} inputNode={inputNode} />;
}
