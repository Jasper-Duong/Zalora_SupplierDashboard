import { SET_SELECTED_SUPPLIER } from "../types/suppliers.type";

const setSelectedSupplier = (data) => ({
  type: SET_SELECTED_SUPPLIER,
  payload: data,
});

export { setSelectedSupplier };
