import React from "react";
import SupplierTabs from "../Tabs/SupplierTabs/SupplierTabs";
import { suppliers } from "../../mockdata/suppliers";
import FormContainer from "../Container/FormContainer";

export default function EditSupplier() {
  return (
    <FormContainer>
      <SupplierTabs supplier={suppliers[0]} />
    </FormContainer>
  );
}
