let productsStock = [
  {
    id: 1,
    name: "Product 1",
    stock: 10,
  },
  {
    id: 2,
    name: "Product 2",
    stock: 20,
  },
  {
    id: 3,
    name: "Product 3",
    stock: 30,
  },
];
function getStockBySupplierIdApi(id) {
  console.log({ productsStock });
  return productsStock;
}

function addStockByApi(supplierId, productId, data) {
  const newId = productsStock.length;
  productsStock.push({ ...data, id: newId });
}

function updateStockApi(supplierId, productId, data) {
  const idx = productsStock.findIndex((item) => item.id === productId);
  productsStock[idx] = data;
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
