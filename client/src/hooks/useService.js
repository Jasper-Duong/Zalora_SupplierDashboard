import { useContext, useEffect, useState } from "react";
import { LoadingContext } from "../contexts/LoadingContext";

const useService = ({ service, condition = true, deps = [] }) => {
  const [, setLoadingState] = useContext(LoadingContext);
  const [data, setData] = useState();
  const fetchData = async () => {
    try {
      setLoadingState(true);
      const result = await service();
      setLoadingState(false);
      setData(result);
    } catch (error) {
      console.log(error);
    }
  };
  useEffect(() => {
    if (condition) {
      console.log("Fetch data");
      fetchData();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, deps);
  return { data };
};
export default useService;
