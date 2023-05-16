import React from "react";
import { useParams } from "react-router-dom";
import {
  addAddress,
  deleteAddress,
  updateAddress,
} from "../../../mockdata/addresses";
import TableForm from "../../../modules/TableForm/TableForm";
import AddressesColumns from "../../Columns/AddressesColumns";
import AddressesCell from "../../../modules/TableForm/Cell/AddressesCell";
import AddRowAddress from "../../AddRowForm/AddRowAddress";
import { getAddressesApi } from "../../../services/addresses";

export default function SupplierAddresses() {
  const { id } = useParams();
  const services = {
    getData: () => getAddressesApi(+id),
    addItem: (_, submitData) => addAddress(id, submitData),
    updateItem: (addressId, submitData) => updateAddress(addressId, submitData),
    deleteItem: (addressId) => deleteAddress(addressId),
  };

  return (
    <TableForm
      services={services}
      columns={AddressesColumns}
      Cell={AddressesCell}
      AddRowComponent={AddRowAddress}
    />
  );
}
