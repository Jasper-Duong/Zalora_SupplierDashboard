import React from "react";
import styled from "styled-components";

export default function FormContainer({ children }) {
  return (
    <Styled>
      <div className="form-container">{children}</div>
    </Styled>
  );
}

const Styled = styled.div`
  display: flex;
  justify-content: center;
  .form-container {
    width: 600px;
  }
`;
