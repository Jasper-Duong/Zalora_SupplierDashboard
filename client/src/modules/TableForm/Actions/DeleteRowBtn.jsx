import { Popconfirm } from "antd";
import React from "react";
import styled from "styled-components";

export default function DeleteRowBtn({ handleDelete, key }) {
  return (
    <Styled>
      <Popconfirm
        className="delete-btn"
        title="Sure to delete?"
        onConfirm={() => handleDelete(key)}
      >
        Delete
      </Popconfirm>
    </Styled>
  );
}

const Styled = styled.span`
  .delete-btn {
    color: #dd3545;
    cursor: pointer;
    &:hover {
      opacity: 0.5;
    }
  }
`;
