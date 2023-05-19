import React from "react";
import styled from "styled-components";
import { CheckOutlined, CloseOutlined } from "@ant-design/icons";
import { Button } from "antd";

export default function OnEditBtns({ setIsEdit }) {
  return (
    <Styled>
      <Button type="text" htmlType="submit">
        <CheckOutlined className="check-icon" />
      </Button>
      <CloseOutlined className="cancel-icon" onClick={() => setIsEdit(false)} />
    </Styled>
  );
}
const Styled = styled.div`
  height: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  .ant-btn {
    padding: 0;
    &:hover {
      background-color: transparent;
    }
  }
  .anticon {
    cursor: pointer;
    &:hover {
      transform: scale(120%);
    }
    svg {
      width: 1.25rem;
      height: 1.25rem;
    }
  }
  .check-icon svg {
    fill: green;
  }
  .cancel-icon svg {
    fill: red;
  }
`;
