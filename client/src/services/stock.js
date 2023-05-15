let productsStock = [...Array(100).keys()].map((idx) => ({
  key: idx,
  name: `Product ${idx + 1}`,
  stock: (idx + 1) * 10,
}));
function getStockBySupplierIdApi(id) {
  console.log({ productsStock });
  return productsStock;
}

function addStockByApi(supplierId, productId, data) {
  const newId = Date.now();
  productsStock.push({ key: newId, ...data });
}

function updateStockApi(supplierId, productId, data) {
  const idx = productsStock.findIndex((item) => item.key === productId);
  let item = productsStock[idx];
  productsStock[idx] = { ...item, ...data };
  return data;
}

function deleteStockApi(supplierId, productId) {
  productsStock = productsStock.filter((item) => item.key !== productId);
}

export {
  getStockBySupplierIdApi,
  addStockByApi,
  updateStockApi,
  deleteStockApi,
};
