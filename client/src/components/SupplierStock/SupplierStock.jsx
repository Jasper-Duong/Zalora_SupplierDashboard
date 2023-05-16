import {
  addStockByApi,
  deleteStockApi,
  getStockBySupplierIdApi,
  updateStockApi,
} from "../../services/stock";
import TableForm from "../../modules/TableForm/TableForm";
import { useParams } from "react-router-dom";
import TableFormColumns from "../../modules/TableForm/Columns/TableFormColumns";
import SupplierStockCell from "../../modules/TableForm/Cell/SupplierStockCell";
import AddRowSupplierStock from "../AddRowForm/AddRowSupplierStock";

export default function SupplierStock() {
  const { id } = useParams();
  const services = {
    getData: () => getStockBySupplierIdApi(id),
    updateItem: (productId, submitData) =>
      updateStockApi(id, productId, submitData),
    addItem: (productId, submitData) =>
      addStockByApi(id, productId, submitData),
    deleteItem: (productId) => deleteStockApi(id, productId),
  };

  return (
    <TableForm
      services={services}
      columns={TableFormColumns}
      Cell={SupplierStockCell}
      AddRowComponent={AddRowSupplierStock}
    />
  );
}
