import { getStockBySupplierIdApi } from "../../services/stock";
import TableForm from "../../modules/TableForm/TableForm";
import { useParams } from "react-router-dom";

export default function SupplierStock() {
  const { id } = useParams();
  const service = () => getStockBySupplierIdApi(id);
  return <TableForm service={service} />;
}
