import { useState } from "react";
import AddProductStock from "./AddProductStock";
import { Button } from "antd";
import EditProductStockForm from "./EditProductStockForm";

export default function ProductStock() {
  // getSuppliersStockByProductId
  const suppliersStock = [
    {
      id: 1,
      name: "Supplier 1",
      stock: 10,
    },
    {
      id: 2,
      name: "Supplier 2",
      stock: 20,
    },
  ];
  const [isAdd, setIsAdd] = useState(false);
  const renderProductStocks = () =>
    suppliersStock.map((supplierStock) => (
      <EditProductStockForm
        key={supplierStock.id}
        supplierStock={supplierStock}
      />
    ));
  return (
    <>
      {renderProductStocks()}
      {isAdd ? (
        <AddProductStock setIsEdit={setIsAdd} />
      ) : (
        <Button onClick={() => setIsAdd(true)}>Add new +</Button>
      )}
    </>
  );
}
