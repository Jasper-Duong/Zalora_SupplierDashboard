const products = [
  {
    id: 1,
    name: "Ela River Elevated Oxford Shirt",
    brandName: "Timberland",
    stock: 10,
    color: "light blue",
    size: "M",
    status: true,
    suppliers: [
      { name: "Zalora", id: 1 },
      { name: "Shopee", id: 3 },
    ],
  },
  {
    name: "Adidas Ultraboost 4.0",
    brandName: "Adidas",
    stock: 20,
    color: "oreo",
    size: "42",
    status: true,
    suppliers: [
      { name: "Zalora", id: 1 },
      { name: "Lazada", id: 2 },
    ],
  },
  {
    name: "Easy Iron Slimfit Pants",
    brandName: "Levi's",
    stock: 5,
    color: "navi",
    size: "40",
    status: true,
    suppliers: [
      { name: "Zalora", id: 1 },
      { name: "Alibaba", id: 6 },
    ],
  },
];

export { products };
