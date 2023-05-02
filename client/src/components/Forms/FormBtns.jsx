import { Button } from "antd";
import React from "react";
import styled from "styled-components";

export default function FormBtns({ closeModal, submitBtnText }) {
  return (
    <Styled>
      <Button onClick={closeModal}>Cancel</Button>
      <Button type="primary" htmlType="submit">
        {submitBtnText}
      </Button>
    </Styled>
  );
}
const Styled = styled.div`
  display: flex;
  justify-content: end;
  gap: 0.5rem;
`;
