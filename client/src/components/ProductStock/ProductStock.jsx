import { useParams } from "react-router-dom";
import {
  addProductStock,
  deleteProductStock,
  getStockByProductId,
  updateProductStock,
} from "../../mockdata/stock";
import TableForm from "../../modules/TableForm/TableForm";
import ProductStockColumns from "../Columns/ProductStockColumns";
import ProductStockCell from "../../modules/TableForm/Cell/ProductStockCell";
import AddRowProductStock from "../AddRowForm/AddRowProductStock";

export default function ProductStock() {
  const { id } = useParams();
  const services = {
    getData: () => getStockByProductId(id),
    updateItem: (supplierId, submitData) =>
      updateProductStock(supplierId, id, submitData),
    addItem: (supplierId, submitData) =>
      addProductStock(supplierId, id, submitData),
    deleteItem: (supplierId) => deleteProductStock(supplierId, id),
  };
  return (
    <TableForm
      services={services}
      columns={ProductStockColumns}
      Cell={ProductStockCell}
      AddRowComponent={AddRowProductStock}
    />
  );
}