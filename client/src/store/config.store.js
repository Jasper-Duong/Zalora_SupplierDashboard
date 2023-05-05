import { combineReducers, createStore } from "redux";
import productReducer from "./reducers/productReducer";
import supplierReducer from "./reducers/supplierReducer";

const rootReducer = combineReducers({
  productReducer,
  supplierReducer,
});

export const store = createStore(
  rootReducer,
  window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__()
);
