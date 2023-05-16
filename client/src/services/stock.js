import { request } from "../config/axios";

function getStockBySupplierIdApi(id) {
  return request({method: "GET", url: `/suppliers/${+id}/stocks`})  
}

async function getStockByProductIdApi(id) {
  return request({ method: "GET", url: `/products/${+id}/stocks` });
}

async function addStockByApi(supplierId, productId, data) {
  const submitData = {
    supplier_id: +supplierId,
    product_id: +productId,
    stock: data.stock,
  };
  console.log("Adding", submitData);
  return request({
    method: "POST",
    url: `/products-suppliers/`,
    data: submitData,
  });
}

async function updateStockApi(supplierId, productId, data) {
  const submitData = {
    product_id: +productId,
    supplier_id: +supplierId,
    stock: data.stock,
  };
  return request({
    method: "PUT",
    url: `/products-suppliers/`,
    data: submitData,
  });
}

async function deleteStockApi(supplierId, productId) {
  return request({
    method: "DELETE",
    url: `products-suppliers/${productId}/${supplierId}`,
  });
}

export {
  getStockBySupplierIdApi,
  getStockByProductIdApi,
  addStockByApi,
  updateStockApi,
  deleteStockApi,
};
