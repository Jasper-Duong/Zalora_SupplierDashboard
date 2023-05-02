import { Tabs } from "antd";
import React from "react";
import EditSupplierForm from "../../Forms/Supplier/EditSupplierForm";
import EditAddress from "../../Forms/Supplier/Address/EditAddress";

export default function SupplierTabs({ closeModal, supplier }) {
  const items = [
    {
      key: "general",
      label: "General",
      children: (
        <EditSupplierForm closeModal={closeModal} supplier={supplier} />
      ),
    },
    {
      key: "addresses",
      label: "Addresses",
      children: <EditAddress />,
    },
  ];
  return <Tabs defaultActiveKey="general" items={items} />;
}
