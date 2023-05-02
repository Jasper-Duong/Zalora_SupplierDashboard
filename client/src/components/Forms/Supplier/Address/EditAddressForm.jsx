import React, { useState } from "react";
import AddressForm from "./AddressForm";
import OnEditAddressBtns from "./OnEditAddressBtns";
import EditAddressBtns from "./EditAddressBtns";

export default function EditAddressForm({ address }) {
  const [isEdit, setIsEdit] = useState(false);

  return (
    <AddressForm
      address={address}
      isEdit={isEdit}
      AddressFormBtns={
        isEdit ? (
          <OnEditAddressBtns setIsEdit={setIsEdit} />
        ) : (
          <EditAddressBtns setIsEdit={setIsEdit} />
        )
      }
    />
  );
}
