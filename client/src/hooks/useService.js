import { useContext, useEffect, useState } from "react";
import { LoadingContext } from "../contexts/LoadingContext";

const useService = ({ service, condition = true, deps = [] }) => {
  const [, setLoadingState] = useContext(LoadingContext);
  const [data, setData] = useState();
  const fetchData = async () => {
    try {
      setLoadingState(true);
      const result = await service();
      console.log("Fetch data", result);
      setLoadingState(false);
      setData(result);
    } catch (error) {
      console.log(error);
    }
  };
  useEffect(() => {
    if (condition) {
      fetchData();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, deps);
  return { data };
};
export default useService;
