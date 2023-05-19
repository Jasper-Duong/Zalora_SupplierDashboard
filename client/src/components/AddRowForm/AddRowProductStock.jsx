import { Col, InputNumber, Row } from "antd";
import React from "react";
import TableFormSelect from "../../modules/TableForm/Input/TableFormSelect";
import { getMissingSuppliersByProductIdApi } from "../../services/product";
import { useForm } from "antd/es/form/Form";
import { Form } from "antd";
import { useParams } from "react-router-dom";

export default function AddRowProductStock({ onFinish, AddConfirmBtn }) {
  const { id } = useParams();
  const service = () => getMissingSuppliersByProductIdApi(id);
  const [form] = useForm();
  return (
    <Form onFinish={onFinish} form={form} style={{ width: "100%" }}>
      <Form.Item hidden name={"id"} />
      <Row gutter={16}>
        <Col span={12}>
          <Form.Item name={"name"} label="Supplier Name">
            <TableFormSelect
              service={service}
              placeholder="Select a Supplier"
              form={form}
            />
          </Form.Item>
        </Col>
        <Col span={7}>
          <Form.Item
            rules={[{ required: true, message: "Missing Stock" }]}
            name={"stock"}
            label="Stock"
          >
            <InputNumber />
          </Form.Item>
        </Col>
        <Col span={5}>
          <AddConfirmBtn />
        </Col>
      </Row>
    </Form>
  );
}
