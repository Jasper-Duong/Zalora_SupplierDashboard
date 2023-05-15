import { Col, Form, InputNumber, Row } from "antd";
import React from "react";
import ProductStockSelect from "../../ProductStock/ProductStockSelect";

export default function ProductStockForm({
  supplierStock,
  isEdit,
  ProductStockBtns,
}) {
  const initialValues = supplierStock
    ? { name: supplierStock.name, stock: supplierStock.stock }
    : {};
  return (
    <Form layout="vertical" disabled={!isEdit} initialValues={initialValues}>
      <Row gutter={16}>
        <Col span={16}>
          <Form.Item
            name={"name"}
            label={"Supplier Name"}
            rules={[{ required: true, message: "Please choose Supplier" }]}
          >
            <ProductStockSelect id={supplierStock?.id} />
          </Form.Item>
        </Col>
        <Col span={5}>
          <Form.Item
            name={"stock"}
            label={"Stock"}
            rules={[{ required: true, message: "Stock is required" }]}
          >
            <InputNumber />
          </Form.Item>
        </Col>
        <Col span={3}>{ProductStockBtns}</Col>
      </Row>
    </Form>
  );
}
