import React from "react";
import AddressForm from "./AddressForm";
import AddAddressFormBtns from "./AddAddressFormBtns";

export default function AddAddress({ setIsEdit }) {
  return (
    <AddressForm
      isEdit={true}
      AddressFormBtns={<AddAddressFormBtns setIsEdit={setIsEdit} />}
    />
  );
}
