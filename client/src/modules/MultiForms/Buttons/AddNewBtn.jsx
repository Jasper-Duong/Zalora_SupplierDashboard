import { Button } from "antd";
import React from "react";
import styled from "styled-components";
import { PlusCircleOutlined, CloseOutlined } from "@ant-design/icons";

export default function AddNewBtns({ setIsEdit }) {
  return (
    <Styled>
      <Button type="text" htmlType="submit">
        <PlusCircleOutlined className="plus-icon" />
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
    width: 1.25rem;
    height: 1.25rem;
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
  .cancel-icon svg {
    fill: red;
  }
`;
