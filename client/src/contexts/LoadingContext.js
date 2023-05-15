import { createContext, useState } from "react";
import LoadingSpinner from "../components/LoadingSpinner/LoadingSpinner";

const DEFAULT_VALUE = { isLoading: false };

const LoadingContext = createContext(DEFAULT_VALUE);

const LoadingProvider = ({ children }) => {
  const [state, setState] = useState(DEFAULT_VALUE);
  return (
    <LoadingContext.Provider value={[state, setState]}>
      {state.isLoading && <LoadingSpinner />}
      {children}
    </LoadingContext.Provider>
  );
};

export { LoadingContext, LoadingProvider };
