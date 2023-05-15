import React from "react";
import AddProductForm from "../Forms/Product/AddProductForm";
import HomeHeader from "../../layout/HomeLayout/HomeHeader";

export default function AddProduct() {
  return (
    <>
      <HomeHeader title={"Add Product"} />
      <AddProductForm />
    </>
  );
}
