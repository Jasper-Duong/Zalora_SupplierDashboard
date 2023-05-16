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
// // getSuppliersStockByProductId
// const suppliersStock = [
//   {
//     id: 1,
//     name: "Supplier 1",
//     stock: 10,
//   },
//   {
//     id: 2,
//     name: "Supplier 2",
//     stock: 20,
//   },
// ];
// const [isAdd, setIsAdd] = useState(false);
// const renderProductStocks = () =>
//   suppliersStock.map((supplierStock) => (
//     <EditProductStockForm
//       key={supplierStock.id}
//       supplierStock={supplierStock}
//     />
//   ));
// return (
//   <>
//     {renderProductStocks()}
//     {isAdd ? (
//       <AddProductStock setIsEdit={setIsAdd} />
//     ) : (
//       <Button onClick={() => setIsAdd(true)}>Add new +</Button>
//     )}
//   </>
// );
