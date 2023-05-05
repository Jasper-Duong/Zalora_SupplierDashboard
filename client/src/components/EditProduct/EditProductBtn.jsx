import { Button } from "antd";
import { products } from "../../mockdata/products";
import React from "react";
import { useDispatch } from "react-redux";
import { setSelectedProduct } from "../../store/actions/products.action";

const mockProduct = products[1];

export default function EditProductBtn({ showModal }) {
  //get selected product info (redux store)
  const selectedProduct = mockProduct;
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
