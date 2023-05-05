import { SET_SELECTED_SUPPLIER } from "../types/suppliers.type";

const DEFAULT_STATE = { selectedSupplier: {} };

const supplierReducer = (state = DEFAULT_STATE, { type, payload }) => {
  switch (type) {
    case SET_SELECTED_SUPPLIER:
      state.selectedSupplier = payload;
      return { ...state };

    default:
      return state;
  }
};

export default supplierReducer;
