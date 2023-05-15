import { Col, Form, Input, Row } from "antd";
import { useForm } from "antd/es/form/Form";
import React, { useEffect } from "react";

export default function BaseForm({
  isEdit,
  isAdd,
  onFinish,
  initialValues,
  ActionBtns,
}) {
  const [form] = useForm();
  useEffect(() => {
    form.setFieldsValue(initialValues);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [initialValues]);
  return (
    <Form
      onFinish={onFinish}
      layout="vertical"
      disabled={!isEdit}
      form={form}
    >
      <Form.Item label="Product ID" name={"name0"} hidden>
        <Input />
      </Form.Item>
      <Row gutter={16}>
        <Col span={7}>
          <Form.Item label="Label 1" name="name1">
            <Input disabled={!isAdd} />
          </Form.Item>
        </Col>
        <Col span={14}>
          <Form.Item label="Label 2" name="name2">
            <Input />
          </Form.Item>
        </Col>
        <Col span={3}>{ActionBtns}</Col>
      </Row>
    </Form>
  );
}
