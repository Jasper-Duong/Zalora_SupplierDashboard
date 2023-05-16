let productsStock = [...Array(100).keys()].map((idx) => ({
  id: idx,
  name: `Product ${idx + 1}`,
  stock: (idx + 1) * 10,
}));
function getStockBySupplierIdApi(id) {
  console.log({ productsStock });
  return productsStock;
}

function addStockByApi(supplierId, productId, data) {
  const newId = Date.now();
  productsStock.push({ id: newId, ...data });
}

function updateStockApi(supplierId, productId, data) {
  const idx = productsStock.findIndex((item) => item.id === productId);
  let item = productsStock[idx];
  productsStock[idx] = { ...item, ...data };
  return data;
}

function deleteStockApi(supplierId, productId) {
  productsStock = productsStock.filter((item) => item.id !== productId);
}

export {
  getStockBySupplierIdApi,
  addStockByApi,
  updateStockApi,
  deleteStockApi,
};
