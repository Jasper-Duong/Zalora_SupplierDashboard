import { request } from "../config/axios";

async function getAddressesApi(id) {
  return request({ method: "GET", url: `/suppliers/${id}/addresses` });
}

async function addAddressApi(id, data) {
  const submitData = { supplier_id: +id, ...data };
  console.log("Adding", submitData);
  return request({ method: "POST", url: `/addresses/`, data: submitData });
}

async function updateAddressApi(id, data) {
  return request({ method: "PUT", url: `/addresses/${id}`, data });
}

async function deleteAddressApi(id) {
  return request({ method: "DELETE", url: `/addresses/${id}` });
}

export { getAddressesApi, addAddressApi, deleteAddressApi, updateAddressApi };
