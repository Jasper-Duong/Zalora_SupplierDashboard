import React from "react";
import HomeHeader from "../../layout/HomeLayout/HomeHeader";
import ProductTabs from "../Tabs/ProductTabs/ProductTabs";
import FormContainer from "../Container/FormContainer";

export default function EditProduct() {
  return (
    <>
      <HomeHeader title={"Edit Product"} />
      <FormContainer>
        <ProductTabs />
      </FormContainer>
    </>
  );
}
