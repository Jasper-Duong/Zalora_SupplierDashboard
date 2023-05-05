import "./App.css";
import AddProduct from "./components/AddProduct/AddProduct";
import AddProductBtn from "./components/AddProduct/AddProductBtn";
import AddSupplier from "./components/AddSupplier/AddSupplier";
import AddSupplierBtn from "./components/AddSupplier/AddSupplierBtn";
import EditProduct from "./components/EditProduct/EditProduct";
import EditProductBtn from "./components/EditProduct/EditProductBtn";
import EditSupplier from "./components/EditSupplier/EditSupplier";
import EditSupplierBtn from "./components/EditSupplier/EditSupplierBtn";
import AntdModal from "./components/Modals/AntdModal";

function App() {
  return (
    <div className="App">
      <h1>Hello World</h1>
      <AntdModal
        title={"Add Product"}
        BodyComponent={AddProduct}
        ShowModalBtn={AddProductBtn}
      />
      <AntdModal
        // openModalBtnText={"Edit Product"}
        title={"Edit Product"}
        BodyComponent={EditProduct}
        ShowModalBtn={EditProductBtn}
      />
      <AntdModal
        title={"Add Supplier"}
        BodyComponent={AddSupplier}
        ShowModalBtn={AddSupplierBtn}
      />
      <AntdModal
        title={"Edit Supplier"}
        BodyComponent={EditSupplier}
        ShowModalBtn={EditSupplierBtn}
      />
    </div>
  );
}

export default App;
