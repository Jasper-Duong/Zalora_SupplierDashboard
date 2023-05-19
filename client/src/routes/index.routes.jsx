import React from "react";
import { useRoutes } from "react-router-dom";
import HomePage from "../pages/HomePage";
import ProductsTable from "../components/Table/ProductsTable";
import AddProduct from "../components/AddProduct/AddProduct";
import SuppliersTable from "../components/Table/SuppliersTable";
import AddSupplier from "../components/AddSupplier/AddSupplier";
import EditProduct from "../components/EditProduct/EditProduct";
import EditSupplier from "../components/EditSupplier/EditSupplier";

export default function Router() {
  return useRoutes([
    {
      path: "/",
      element: <HomePage />,
      children: [
        { path: "/", element: <ProductsTable /> },
        {
          path: "products",
          children: [
            { path: "table", element: <ProductsTable /> },
            { path: "add", element: <AddProduct /> },
            { path: "edit/:id", element: <EditProduct /> },
          ],
        },
        {
          path: "suppliers",
          children: [
            { path: "table", element: <SuppliersTable /> },
            { path: "add", element: <AddSupplier /> },
            { path: "edit/:id", element: <EditSupplier /> },
          ],
        },
      ],
    },
  ]);
}
