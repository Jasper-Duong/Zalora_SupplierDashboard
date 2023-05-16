const address = [
  {
    id: 1,
    supplierId: 1,
    type: "office",
    address_info: "12 Ton Dan, District 4, HCM City, Vietnam",
  },
  {
    id: 2,
    supplierId: 1,
    type: "warehouse",
    address_info: "123 Street A, District B, Singapore",
  },
  {
    id: 3,
    supplierId: 2,
    type: "warehouse",
    address_info: "123 Street B, Tan Binh District, HCM City, Vietnam",
  },
];

let addresses = [...Array(5).keys()].map((idx) => ({
  id: idx,
  address_info: `Street ${(idx + 1) * 10}, District ${idx + 1}`,
  type: "warehouse",
}));

const getAddressesBySupplierId = (querySupplierId) => {
  return addresses;
};

const addAddress = (supplierId, submitData) => {
  const newId = Date.now();
  addresses.push({ id: newId, ...submitData });
};

const deleteAddress = (id) => {
  addresses = addresses.filter((address) => address.id !== id);
};

const updateAddress = (id, submitData) => {
  const idx = addresses.findIndex((ele) => ele.id === id);
  if (idx >= 0) {
    addresses.splice(idx, 1, { id, ...submitData });
  }
};

export {
  addresses,
  getAddressesBySupplierId,
  deleteAddress,
  addAddress,
  updateAddress,
};
