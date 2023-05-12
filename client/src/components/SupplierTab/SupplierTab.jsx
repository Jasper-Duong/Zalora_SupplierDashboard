import SuppliersTable from "../Table/SuppliersTable";
import AntdModal from "../Modals/AntdModal";
import AddSupplier from "../AddSupplier/AddSupplier";
import AddSupplierBtn from "../AddSupplier/AddSupplierBtn";

export const SupplierTab = () => {

  return (
    <div className="supplier-container">
      {/* <Button type="primary" style={{float: "left"}} onClick={addNewSupplier}>New supplier</Button> */}
      <AntdModal
        BodyComponent={AddSupplier}
        ShowModalBtn={AddSupplierBtn}
        title="Add Supplier"
      />
      <SuppliersTable />
    </div>
  );
};
