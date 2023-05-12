import React from "react";
import SupplierTabs from "../Tabs/SupplierTabs/SupplierTabs";
import { useSelector } from "react-redux";

export default function EditSupplier({ closeModal }) {
  const { selectedSupplier } = useSelector((state) => state.supplierReducer);
  return (
    <div>
      <SupplierTabs closeModal={closeModal} supplier={selectedSupplier} />
    </div>
  );
}
