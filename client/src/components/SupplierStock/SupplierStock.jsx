import React, { useState } from "react";
import MultiForms from "../../modules/MultiForms/MultiForms";
import EditForm from "../../modules/MultiForms/EditForm";
import AddForm from "../../modules/MultiForms/AddForm";
import { getStockBySupplierIdApi } from "../../services/stock";

export default function SupplierStock() {
  const data = getStockBySupplierIdApi(1);
  const [, setIsForceRender] = useState(0);
  const forceRender = () => {
    setIsForceRender((prev) => prev + 1);
  };
  return <MultiForms forceRender={forceRender} data={data} EditForm={EditForm} AddForm={AddForm} />;
}
