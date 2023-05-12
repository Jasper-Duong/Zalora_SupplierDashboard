import { Button } from "antd";
import React from "react";
import styled from "styled-components";

export default function FormBtns({ CancelBtn, submitBtnText }) {
  return (
    <Styled>
      {CancelBtn}
      {/* <Button onClick={onCancel}>Cancel</Button> */}
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
