import React from "react";
import styled from "styled-components";

export default function HomeHeader({title}) {
  return (
    <Styled>
      <h2>{title}</h2>
    </Styled>
  );
}

const Styled = styled.div`
  background-color: rgba(245, 245, 245, 1);
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 2rem;
  height: 3.75rem;
`;
