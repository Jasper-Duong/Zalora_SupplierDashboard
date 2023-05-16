const supplierStock = [...Array(15).keys()].map((idx) => ({
  id: idx,
  name: `Supplier ${idx + 1}`,
  stock: idx * 100,
}));

const getStockByProductId = (id) => supplierStock;

const addProductStock = (supplierId, productId, data) => {
  supplierStock.push({ id: supplierId, ...data });
};

const updateProductStock = (supplierId, productId, data) => {
  supplierStock.forEach((ele, idx) => {
    if(ele.id === supplierId) {
      supplierStock[idx] = {...ele, ...data}
    }
  })
};

const deleteProductStock = (supplierId, productId) => {
  const idx = supplierStock.findIndex((item) => item.id === supplierId);
  if (idx >= 0) {
    supplierStock.splice(idx, 1);
  }
};
export {
  getStockByProductId,
  addProductStock,
  updateProductStock,
  deleteProductStock,
};
