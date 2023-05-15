import React from "react";
import HomeLayout from "../layout/HomeLayout/HomeLayout";
import { LoadingProvider } from "../contexts/LoadingContext";

export default function HomePage() {
  return (
    <LoadingProvider>
      <HomeLayout />
    </LoadingProvider>
  );
}
