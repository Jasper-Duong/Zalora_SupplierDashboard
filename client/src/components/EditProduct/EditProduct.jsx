import React from "react";
import EditProductForm from "../Forms/EditProductForm";
import { products } from "../../mockdata/products";
import { productToFormValue } from "../../utils/products";

export default function EditProduct({ closeModal }) {
  return (
    <div>
      <EditProductForm closeModal={closeModal} product={productToFormValue(products[2])} />
    </div>
  );
}
