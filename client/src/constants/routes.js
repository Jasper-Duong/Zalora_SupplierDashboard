const routes = {
  PRODUCT_TABLE: {
    path: "/products/table",
    homeHeader: "Product Table",
  },
  ADD_PRODUCT: { path: "/products/add", homeHeader: "Add Product" },
  EDIT_PRODUCT: (id) => ({
    path: `/products/edit/${id}`,
    homeHeader: "Edit Product",
  }),
  SUPPLIER_TABLE: { path: "/suppliers/table", homeHeader: "Supplier Table" },
  ADD_SUPPLIER: { path: "/suppliers/add", homeHeader: "Add Supplier" },
  EDIT_SUPPLIER: (id) => ({
    path: `/suppliers/edit/${id}`,
    homeHeader: "Edit Product",
  }),
};

export { routes };
