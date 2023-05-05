import { combineReducers, createStore } from "redux";
import productReducer from "./reducers/productReducer";

const rootReducer = combineReducers({
  productReducer,
});

export const store = createStore(
  rootReducer,
  window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__()
);
