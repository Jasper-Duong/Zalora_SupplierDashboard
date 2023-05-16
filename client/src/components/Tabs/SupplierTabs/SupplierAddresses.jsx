import React from "react";
import { useParams } from "react-router-dom";
import TableForm from "../../../modules/TableForm/TableForm";
import AddressesColumns from "../../Columns/AddressesColumns";
import AddressesCell from "../../../modules/TableForm/Cell/AddressesCell";
import AddRowAddress from "../../AddRowForm/AddRowAddress";
import { addAddressApi, deleteAddressApi, getAddressesApi, updateAddressApi } from "../../../services/addresses";

export default function SupplierAddresses() {
  const { id } = useParams();
  const services = {
    getData: () => getAddressesApi(+id),
    addItem: (_, submitData) => addAddressApi(id, submitData),
    updateItem: (addressId, submitData) => updateAddressApi(addressId, submitData),
    deleteItem: (addressId) => deleteAddressApi(addressId),
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
