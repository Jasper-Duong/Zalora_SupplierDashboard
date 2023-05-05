const addresses = [
  {
    id: 1,
    supplierId: 1,
    type: "office",
    addressInfo: "12 Ton Dan, District 4, HCM City, Vietnam",
  },
  {
    id: 2,
    supplierId: 1,
    type: "warehouse",
    addressInfo: "123 Street A, District B, Singapore",
  },
  {
    id: 3,
    supplierId: 2,
    type: "warehouse",
    addressInfo: "123 Street B, Tan Binh District, HCM City, Vietnam",
  },
];

const getAddressesBySupplierId = (querySupplierId) =>
  addresses.filter(({ supplierId }, _) => supplierId === querySupplierId);

export { addresses, getAddressesBySupplierId };
