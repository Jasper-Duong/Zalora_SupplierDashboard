import "./App.css";
import EditProduct from "./components/EditProduct/EditProduct";
import EditSupplier from "./components/EditSupplier/EditSupplier";
import AntdModal from "./components/Modals/AntdModal";

function App() {
  return (
    <div className="App">
      <h1>Hello World</h1>
      <AntdModal
        openModalBtnText={"Edit Product"}
        title={"Edit Product"}
        BodyComponent={EditProduct}
      />
      <AntdModal
        openModalBtnText={"Edit Supplier"}
        title={"Edit Supplier"}
        BodyComponent={EditSupplier}
      />
    </div>
  );
}

export default App;
