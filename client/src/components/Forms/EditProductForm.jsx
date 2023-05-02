import { Form, Input, InputNumber, Select, Switch } from "antd";
import React from "react";
import FormBtns from "./FormBtns";
import { supplierOptions } from "../../mockdata/suppliers";
import { formLayout } from "../../constants/forms";

export default function EditProductForm({ closeModal, product }) {
  const onFinish = (values) => {
    console.log({ values });
    closeModal();
  };
  return (
    <Form
      onFinish={onFinish}
      labelCol={formLayout.labelCol}
      wrapperCol={formLayout.wrapperCol}
      initialValues={product ? { ...product } : {}}
    >
      <Form.Item
        label="Name"
        name="name"
        rules={[{ required: true, message: "Product Name is required" }]}
      >
        <Input />
      </Form.Item>
      <Form.Item
        label="Brand Name"
        name="brandName"
        rules={[{ required: true, message: "Brand Name is required" }]}
      >
        <Input />
      </Form.Item>
      <Form.Item
        label="Stock"
        name="stock"
        rules={[{ required: true, message: "Product Stock is required" }]}
      >
        <InputNumber min={0} />
      </Form.Item>
      <Form.Item
        label="Color"
        name="color"
        rules={[{ required: true, message: "Product Color is required" }]}
      >
        <Input />
      </Form.Item>
      <Form.Item
        label="Size"
        name="size"
        rules={[{ required: true, message: "Product Size is required" }]}
      >
        <Input />
      </Form.Item>
      <Form.Item label="Status" name="status" valuePropName="checked">
        <Switch checkedChildren="Active" unCheckedChildren="Inactive" />
      </Form.Item>
      <Form.Item
        label="Suppliers"
        name="suppliers"
        rules={[
          {
            required: true,
            message: "Product must has at least 1 supplier",
            type: "array",
          },
        ]}
      >
        <Select mode="multiple" options={supplierOptions} />
      </Form.Item>
      <FormBtns submitBtnText="Update" closeModal={closeModal} />
    </Form>
  );
}
