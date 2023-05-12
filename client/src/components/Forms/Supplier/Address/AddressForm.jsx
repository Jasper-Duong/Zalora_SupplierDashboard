import { Col, Form, Input, Row, Select } from "antd";
import React from "react";

export default function AddressForm({ isEdit, AddressFormBtns, address }) {
  const onFinish = (address) => {
    console.log({ address });
  };
  return (
    <Form
      onFinish={onFinish}
      layout="vertical"
      disabled={!isEdit}
      initialValues={
        address ? { type: address.type, addressInfo: address.addressInfo } : {type:"office"}
      }
    >
      <Row gutter={16}>
        <Col span={7}>
          <Form.Item label="Type" name="type">
            <Select>
              <Select.Option value="office">Office</Select.Option>
              <Select.Option value="warehouse">Warehouse</Select.Option>
            </Select>
          </Form.Item>
        </Col>
        <Col span={14}>
          <Form.Item
            label="Address Info"
            name="addressInfo"
            rules={[{ required: true, message: "Address Info is required" }]}
          >
            <Input />
          </Form.Item>
        </Col>
        <Col span={3}>{AddressFormBtns}</Col>
      </Row>
    </Form>
  );
}
