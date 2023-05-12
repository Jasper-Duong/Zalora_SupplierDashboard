const supplierNames = [
  "Zalora",
  "Lazada",
  "Shopee",
  "Sephora",
  "Tiki",
  "Alibaba",
  "Amazon",
];

const suppliers = supplierNames.map((name, idx) => ({
  name,
  id: idx + 1,
  email: `${name}@jopmail.com`,
  contactNumber: `+1234567890${idx}`,
  status: true,
}));
const supplierOptions = suppliers.map(({ name, id }) => ({
  value: id,
  label: name,
}));

export { supplierOptions, suppliers };
