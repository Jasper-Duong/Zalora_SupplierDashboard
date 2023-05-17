import { request } from "../config/axios";

async function getSupplierByIdApi(id) {
  return request({ method: "GET", url: `/suppliers/${id}` });
}

async function getMissingProductsBySupplierIdApi(id) {
  return request({ method: "GET", url: `/suppliers/${id}/products/missing` });
}

async function addSupplierByIdApi(supplier) {
  return request({
    method: "POST",
    url: `suppliers/`,
    data: JSON.stringify(supplier),
  });
}

async function updateSupplierByIdApi(id, supplier) {
  return request({
    method: "PUT",
    url: `suppliers/${id}`,
    data: JSON.stringify(supplier),
  });
}

export {
  getSupplierByIdApi,
  getMissingProductsBySupplierIdApi,
  addSupplierByIdApi,
  updateSupplierByIdApi,
};
