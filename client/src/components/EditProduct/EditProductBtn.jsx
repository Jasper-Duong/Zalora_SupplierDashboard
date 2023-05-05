import { Button } from "antd";
import { products } from "../../mockdata/products";
import React from "react";
import { useDispatch } from "react-redux";
import { setSelectedProduct } from "../../store/actions/products.action";

export default function EditProductBtn({ showModal, data }) {
  //get selected product info (redux store)
  const selectedProduct = data;
  const dispatch = useDispatch();
  const handleClick = () => {
    // dispatch selectedProduct -> redux store
    dispatch(setSelectedProduct(selectedProduct));
    showModal();
  };
  return (
    <Button type="primary" onClick={handleClick}>
      Edit Product
    </Button>
  );
}
