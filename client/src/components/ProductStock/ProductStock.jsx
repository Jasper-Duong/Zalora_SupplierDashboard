import { useParams } from "react-router-dom";
import TableForm from "../../modules/TableForm/TableForm";
import ProductStockColumns from "../Columns/ProductStockColumns";
import ProductStockCell from "../../modules/TableForm/Cell/ProductStockCell";
import AddRowProductStock from "../AddRowForm/AddRowProductStock";
import { addStockByApi, deleteStockApi, getStockByProductIdApi, updateStockApi } from "../../services/stock";

export default function ProductStock() {
  const { id } = useParams();
  const services = {
    getData: () => getStockByProductIdApi(id),
    updateItem: (supplierId, submitData) =>
      updateStockApi(supplierId, id, submitData),
    addItem: (supplierId, submitData) =>
      addStockByApi(supplierId, id, submitData),
    deleteItem: (supplierId) => deleteStockApi(supplierId, id),
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
