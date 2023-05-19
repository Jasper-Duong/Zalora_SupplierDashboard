import { InputNumber } from "antd";
import EditableCell from "../../../modules/TableForm/Rows/EditableRow";

export default function ProductStockCell({ inputType, ...rest }) {
  return <EditableCell {...rest} inputNode={<InputNumber />} />;
}
