import React from "react";
import { Spin } from "antd";
import styled from "styled-components";
export default function LoadingSpinner() {
  return (
    <Styled>
      <Spin />
    </Styled>
  );
}

const Styled = styled.div`
  height: 100%;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: rgba(0, 0, 0, 0.5);
`;
