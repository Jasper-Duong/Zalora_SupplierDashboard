import React from "react";
import styled from "styled-components";
import { EditOutlined, DeleteOutlined } from "@ant-design/icons";

export default function EditBtns({ setIsEdit, handleDelete }) {
  return (
    <Styled>
      <EditOutlined className="edit-icon" onClick={() => setIsEdit(true)} />
      <DeleteOutlined className="delete-icon" onClick={handleDelete} />
    </Styled>
  );
}

const Styled = styled.div`
  height: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
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
`;
