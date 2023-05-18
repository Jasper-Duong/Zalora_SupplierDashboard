import { request } from "../config/axios";

async function getProductByIdApi(id) {
  return request.get(`products/${id}`);
}

async function getMissingSuppliersByProductIdApi(id) {
  return request.get(`products/${id}/suppliers/missing`);
}

async function addProductByIdApi(product) {
  return request({
    method: "POST",
    url: `products/`,
    data: JSON.stringify(product),
  });
}

async function updateProductByIdApi(id, product) {
  return request.put(`products/${id}`, JSON.stringify(product));
}

async function deleteProductApi(id) {
  return request.delete(`products/${id}`);
}

export {
  getProductByIdApi,
  getMissingSuppliersByProductIdApi,
  addProductByIdApi,
  updateProductByIdApi,
  deleteProductApi,
};
