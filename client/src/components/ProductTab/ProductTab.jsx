import { Button } from "antd";
import ProductsTable from "../Table/ProductsTable";
import AntdModal from "../Modals/AntdModal";
import AddProduct from "../AddProduct/AddProduct";
import AddProductBtn from "../AddProduct/AddProductBtn";

export const ProductTab = () => {
  return (
    <div className="product-container">
      <AntdModal
        ShowModalBtn={AddProductBtn}
        title="Add Product"
        BodyComponent={AddProduct}
      />

      <ProductsTable />
    </div>
  );
};
