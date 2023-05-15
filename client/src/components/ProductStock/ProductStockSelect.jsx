import { Select } from "antd";
import React from "react";

export default function ProductStockSelect({ id }) {
  // getSuppliersByProductId
  const suppliers = [
    { id: 1, name: "Zalora" },
    { id: 2, name: "Tiki" },
    { id: 3, name: "Shopee" },
  ];
  return (
    <Select value={id || ""}>
      {suppliers.map(({ name, id }) => (
        <Select.Option key={id} value={id}>
          {name}
        </Select.Option>
      ))}
    </Select>
  );
}
