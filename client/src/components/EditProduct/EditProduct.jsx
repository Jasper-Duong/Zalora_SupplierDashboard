import React from "react";
// import { products } from "../../mockdata/products";
import { productToFormValue } from "../../utils/products";
import { useSelector } from "react-redux";
import EditProductForm from "../Forms/Product/EditProductForm";

export default function EditProduct({ closeModal }) {
  const { selectedProduct } = useSelector((state) => state.productReducer);
  return (
    <div>
      <EditProductForm
        closeModal={closeModal}
        product={productToFormValue(selectedProduct)}
      />
    </div>
  );
}
