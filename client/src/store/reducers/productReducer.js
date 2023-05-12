import { SET_SELECTED_PRODUCT } from "../types/products.type";

const DEFAULT_STATE = { selectedProduct: {} };

const productReducer = (state = DEFAULT_STATE, { type, payload }) => {
  switch (type) {
    case SET_SELECTED_PRODUCT:
      state.selectedProduct = payload;
      return { ...state };

    default:
      return state;
  }
};
export default productReducer;
