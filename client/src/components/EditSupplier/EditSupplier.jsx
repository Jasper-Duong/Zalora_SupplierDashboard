import React from "react";
import SupplierTabs from "../Tabs/SupplierTabs/SupplierTabs";
import { suppliers } from "../../mockdata/suppliers";

export default function EditSupplier({ closeModal }) {
  return (
    <div>
      <SupplierTabs closeModal={closeModal} supplier={suppliers[0]} />
    </div>
  );
}
