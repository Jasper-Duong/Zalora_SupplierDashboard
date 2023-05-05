import { SET_SELECTED_PRODUCT } from "../types/products.type";

const setSelectedProduct = (data) => {
  return{
  type: SET_SELECTED_PRODUCT,
  payload: data,
}};

export { setSelectedProduct };