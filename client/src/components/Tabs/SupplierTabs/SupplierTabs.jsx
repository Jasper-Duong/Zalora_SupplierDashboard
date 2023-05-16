import { Tabs } from "antd";
import React from "react";
import EditSupplierForm from "../../Forms/Supplier/EditSupplierForm";
import EditAddress from "../../Forms/Supplier/Address/EditAddress";
import HomeHeader from "../../../layout/HomeLayout/HomeHeader";
import SupplierStock from "../../SupplierStock/SupplierStock";
import SupplierAddresses from "./SupplierAddresses";

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
      // children: <EditAddress />,
      children: <SupplierAddresses />,
    },
    {
      key: "stock",
      label: "Stock",
      children: <SupplierStock />,
    },
  ];
  return (
    <>
      <HomeHeader title={"Supplier Info"} />
      <Tabs destroyInactiveTabPane defaultActiveKey="general" items={items} />
    </>
  );
}
