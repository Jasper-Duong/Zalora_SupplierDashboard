import React from "react";
import ProductStock from "../../ProductStock/ProductStock";
import { Tabs } from "antd";
import EditProductForm from "../../Forms/Product/EditProductForm";

export default function ProductTabs() {
  const items = [
    {
      key: "general",
      label: "General",
      children: <EditProductForm />,
    },
    {
      key: "stock",
      label: "Stock",
      children: <ProductStock />,
    },
  ];
  return <Tabs items={items} />;
}
