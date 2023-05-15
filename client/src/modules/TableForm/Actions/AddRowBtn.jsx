import { Button, Col, Form, InputNumber, Popconfirm, Row } from "antd";
import React, { useState } from "react";
import styled from "styled-components";
import TableFormSelect from "../Input/TableFormSelect";
import { addStockByApi } from "../../../services/stock";
import { useForm } from "antd/es/form/Form";

export default function AddRowBtn({ forceRender }) {
  const [isAdd, setIsAdd] = useState(false);
  const [form] = useForm();
  const handleAdd = (values) => {
    addStockByApi(1, values.id, { name: values.name, stock: values.stock });
    setIsAdd(false);
    forceRender();
  };
  return (
    <Styled isAdd={isAdd}>
      {isAdd ? (
        <Form onFinish={handleAdd} form={form}>
          <Form.Item hidden name={"id"} />
          <Row gutter={16}>
            <Col span={12}>
              <Form.Item name={"name"} label="Product Name">
                <TableFormSelect form={form} />
              </Form.Item>
            </Col>
            <Col span={7}>
              <Form.Item name={"stock"} label="Stock">
                <InputNumber />
              </Form.Item>
            </Col>
            <Col span={5}>
              <div className="add-btn-confirm">
                <Button
                  type="primary"
                  htmlType="submit"
                  style={{ marginRight: 8 }}
                >
                  Add
                </Button>
                <Popconfirm
                  title="Sure to cancel?"
                  onConfirm={() => setIsAdd(false)}
                >
                  <a>Cancel</a>
                </Popconfirm>
              </div>
            </Col>
          </Row>
        </Form>
      ) : (
        <Button
          className="add-new-btn"
          type="primary"
          onClick={() => setIsAdd(true)}
        >
          Add new +
        </Button>
      )}
    </Styled>
  );
}

const Styled = styled.div`
  display: flex;
  justify-content: ${({ isAdd }) => (isAdd ? "space-between" : "flex-end")};
  .add-new-btn {
    margin-bottom: 1.25rem;
  }
  .add-btn-confirm {
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: flex-end;
  }
`;
