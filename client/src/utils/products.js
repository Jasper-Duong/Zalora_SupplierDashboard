const productToFormValue = (product) => {
  const { suppliers, ...rest } = product;
  return {
    ...rest,
    suppliers: suppliers.map(({ name, id }) => ({ value: id, label: name })),
  };
};

export { productToFormValue };
