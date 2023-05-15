import { Button } from "antd";
import { getAddressesBySupplierId } from "../../../../mockdata/addresses";
import EditAddressForm from "./EditAddressForm";
import { useState } from "react";
import AddAddress from "./AddAddress";
import FormContainer from "../../../Container/FormContainer";

export default function EditAddress() {
  const zaloraAddresses = getAddressesBySupplierId(1);
  const renderAddresses = () =>
    zaloraAddresses.map((address) => (
      <EditAddressForm key={address.id} address={address} />
    ));
  const [isAdd, setIsAdd] = useState(false);
  return (
    <FormContainer>
      {renderAddresses()}
      {isAdd ? (
        <AddAddress setIsEdit={setIsAdd} />
      ) : (
        <Button onClick={() => setIsAdd(true)}>Add new +</Button>
      )}
    </FormContainer>
  );
}
