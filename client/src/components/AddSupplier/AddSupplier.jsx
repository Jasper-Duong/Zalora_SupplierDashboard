import React from "react";
import AddSupplierForm from "../Forms/Supplier/AddSupplierForm";
import HomeHeader from "../../layout/HomeLayout/HomeHeader";

export default function AddSupplier() {
  return (
    <>
      <HomeHeader title={"Add Supplier"} />
      <AddSupplierForm />
    </>
  );
}
